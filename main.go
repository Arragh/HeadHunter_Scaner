package main

import (
	"fmt"
	"hhscaner/configuration"
	"hhscaner/service/headhunter"
	"hhscaner/service/httphandler"
	"hhscaner/service/notifier"
	"hhscaner/service/serializer"
	"hhscaner/service/storage"
	"log"
	"time"
)

const (
	fileName string = "viewed_vacancies.txt"
)

func main() {
	config, err := configuration.GetConfigurartion()
	if err != nil {
		fmt.Printf("Ошибка загрузки конфигурации: %v", err)
		log.Fatal(err)
	}

	httpClient := &httphandler.DefaultHttpClient{}
	var triesCount int = 1

	for {
		fmt.Printf("Попытка %d\n", triesCount)
		triesCount++

		oldVacanciesIds, err := storage.ReadData(fileName)
		if err != nil {
			log.Fatal(err)
		}

		vacanciesIds, err := getVacanciesIds(config, httpClient)
		if err != nil {
			log.Fatal(err) // TODO: изменить на просто логирование
		}

		dif := headhunter.Difference(vacanciesIds, oldVacanciesIds)
		if err != nil {
			log.Fatal(err)
		}

		err = storage.SaveData(dif, fileName)
		if err != nil {
			log.Fatal(err)
		}

		if len(dif) > 0 {
			notifier.TriggerAlert(dif)

			for _, id := range dif {
				vacancyUrl := fmt.Sprintf("%s/vacancy/%d", config.HeadHunter.BaseUrl, id)
				fmt.Println(vacancyUrl)
				sendNotificationToTelegram(config, httpClient, vacancyUrl)
			}
		} else {
			time.Sleep(time.Duration(config.RequestIntervalInSeconds) * time.Second)
		}
	}
}

func getVacanciesIds(config *configuration.Config, client httphandler.HttpClient) ([]int64, error) {
	baseUrl := config.HeadHunter.ApiUrl + "/vacancies"
	builderUrl, err := httphandler.BuildUrl(baseUrl, &config.UrlParameters)
	if err != nil {
		return nil, fmt.Errorf("ошибка построения URL: %v", err)
	}

	body, err := client.Get(builderUrl)
	if err != nil {
		return nil, fmt.Errorf("ошибка получения тела ответа: %v", err)
	}

	vacancies, err := serializer.Deserialize[headhunter.VacancyResponse](body)
	if err != nil {
		return nil, fmt.Errorf("ошибка демаршалинга: %v", err)
	}

	vacanciesIds, err := headhunter.ParseVacanciesIds(vacancies.Items)
	if err != nil {
		return nil, fmt.Errorf("ошибка парсинга вакансий: %v", err)
	}

	return vacanciesIds, nil
}

func sendNotificationToTelegram(config *configuration.Config, client httphandler.HttpClient, vacancyUrl string) error {
	params := []configuration.UrlParameter{
		{
			Key:   "chat_id",
			Value: config.Telegram.ChatId,
		},
		{
			Key:   "text",
			Value: vacancyUrl,
		},
	}

	baseUrl := config.Telegram.ApiUrl + "/bot" + config.Telegram.BotToken + "/sendMessage"

	buildedUrl, err := httphandler.BuildUrl(baseUrl, &params)
	if err != nil {
		return fmt.Errorf("ошибка построения URL: %v", err)
	}

	err = notifier.SendNotificationToTelegram(client, buildedUrl, vacancyUrl)
	if err != nil {
		return fmt.Errorf("ошибка отправки уведомления: %v", err)
	}

	return nil
}
