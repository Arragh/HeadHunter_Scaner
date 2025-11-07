package main

import (
	"HeadHunter_Scaner/client"
	"HeadHunter_Scaner/notification"
	"HeadHunter_Scaner/storage"
	"HeadHunter_Scaner/vacancy"
	"fmt"
)

func main() {
	fmt.Println(
		"👋 Hello, World!",
	)

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

	var dif = vacancy.Difference(*oldVacancies, *newVacancies)
	if len(dif) > 0 {
		go notification.TriggerAlert(&dif)
	}

	var meshedVacancies = vacancy.MergeVacancies(*oldVacancies, *newVacancies)

	err = storage.SaveDataToFile(meshedVacancies, "output.json")
	if err != nil {
		fmt.Printf("Ошибка сохранения данных: %v\n", err)
		panic(err)
	}
}
