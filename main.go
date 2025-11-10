package main

import (
	"HeadHunter_Scaner/client"
	"HeadHunter_Scaner/notification"
	"HeadHunter_Scaner/storage"
	"HeadHunter_Scaner/vacancy"
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
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

		newVacancies, err := client.FetchVacancies()
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
			notification.TriggerAlert(&dif)
			pressToContinue()
		} else {
			time.Sleep(5 * time.Second)
		}
	}
}

func pressToContinue() {
	fmt.Print("🔥🔥🔥🔥🔥 => Прочитал? Нажми ENTER!!! <= 🔥🔥🔥🔥🔥")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
	fmt.Print("\n\n\n\n\n")
	fmt.Println("Продролжаем сканирование...")
}
