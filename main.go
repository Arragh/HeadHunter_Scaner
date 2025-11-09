package main

import (
	"HeadHunter_Scaner/client"
	"HeadHunter_Scaner/notification"
	"HeadHunter_Scaner/storage"
	"HeadHunter_Scaner/vacancy"
	"fmt"
	"time"
)

func main() {
	for {
		oldVacancies, err := storage.ReadDataFromFile("output.json")
		if err != nil {
			fmt.Printf("Ошибка получения старых вакансий: %v\n", err)
			panic(err)
		}

		baseUrl := "https://api.hh.ru/vacancies"

		newVacancies, err := client.FetchVacancies(baseUrl)
		if err != nil {
			fmt.Printf("Ошибка получения новых вакансий: %v\n", err)
			panic(err)
		}

		var dif = vacancy.Difference(*newVacancies, *oldVacancies)

		var meshedVacancies = vacancy.MergeVacancies(*oldVacancies, *newVacancies)

		err = storage.SaveDataToFile(meshedVacancies, "output.json")
		if err != nil {
			fmt.Printf("Ошибка сохранения данных: %v\n", err)
			panic(err)
		}

		if len(dif) > 0 {
			notification.TriggerAlert(&dif)
		}

		time.Sleep(60 * time.Second)
	}
}
