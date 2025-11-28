package main

import (
	"fmt"
	"hhscaner/configuration"
	"hhscaner/service/headhunter"
	"hhscaner/service/httphandler"
	"hhscaner/service/notifier"
	"hhscaner/service/storage"
	"log"
	"time"
)

var fileName string = "viewed_vacancies.txt"
var triesCount int = 1

func main() {
	config, err := configuration.GetConfigurartion()
	if err != nil {
		fmt.Printf("Ошибка загрузки конфигурации: %v", err)
		log.Fatal(err)
	}

	httpClient := &httphandler.DefaultHttpClient{}

	for {
		fmt.Printf("Попытка %d\n", triesCount)
		triesCount++

		oldVacanciesIds, err := storage.ReadData(fileName)
		if err != nil {
			log.Fatal(err)
		}

		vacanciesIds, err := headhunter.GetVacanciesIds(config, httpClient)
		if err != nil {
			log.Fatal(err) // TODO: изменить на просто логирование
		}

		dif, err := headhunter.Difference(vacanciesIds, oldVacanciesIds)
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
				notifier.SendNotificationToTelegram(vacancyUrl, httpClient)
			}
		} else {
			time.Sleep(time.Duration(config.RequestIntervalInSeconds) * time.Second)
		}
	}
}
