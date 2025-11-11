package main

import (
	"HeadHunter_Scaner/client"
	"HeadHunter_Scaner/config"
	"HeadHunter_Scaner/notifier"
	"HeadHunter_Scaner/storage"
	"HeadHunter_Scaner/vacancy"
	"fmt"
	"time"
)

func main() {
	config, err := config.GetConfigurartion()
	if err != nil {
		fmt.Printf("Ошибка загрузки конфигурации: %v", err)
		panic(err)
	}

	viewedVacancies := "viewed_vacancies.json"
	triesCount := 1

	for {
		fmt.Printf("Попытка %d\n", triesCount)
		triesCount++

		oldVacanciesIds, err := storage.ReadDataFromFile(viewedVacancies)
		if err != nil {
			fmt.Printf("Ошибка получения старых вакансий: %v\n", err)
			panic(err)
		}

		newVacancies, err := client.FetchVacancies(config)
		if err != nil {
			fmt.Printf("Ошибка получения новых вакансий: %v\n", err)
			panic(err)
		}

		dif, err := vacancy.Difference(*newVacancies, *oldVacanciesIds)
		if err != nil {
			fmt.Printf("Ошибка вычисления новых вакансий: %v\n", err)
			panic(err)
		}

		mergedVacancies, err := vacancy.MergeVacancies(*oldVacanciesIds, *newVacancies)
		if err != nil {
			fmt.Printf("Ошибка объединения вакансий: %v\n", err)
			panic(err)
		}

		err = storage.SaveDataToFile(mergedVacancies, viewedVacancies)
		if err != nil {
			fmt.Printf("Ошибка сохранения данных: %v\n", err)
			panic(err)
		}

		if len(dif) > 0 {
			notifier.TriggerAlert(&dif)
		} else {
			time.Sleep(time.Duration(config.RequestIntervalInSeconds) * time.Second)
		}
	}
}
