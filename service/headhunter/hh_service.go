// Package headhunter предоставляет функции для работы с api hh.ru и вакансиями
package headhunter

import (
	"fmt"
	"hhscaner/configuration"
	"hhscaner/service/httphandler"
	"hhscaner/service/serializer"
	"strconv"
)

// Difference сравнивает полученные вакансии с сохраненными и возвращает разницу
func Difference(newVacanciesIds []int64, oldVacanciesIds []int64) []int64 {
	temp := make(map[int64]bool)

	for _, id := range oldVacanciesIds {
		temp[id] = true
	}

	var result []int64
	for _, id := range newVacanciesIds {
		if !temp[id] {
			result = append(result, id)
		}
	}

	return result
}

// GetVacanciesIds обращается к api hh.ru и возвращает список ID вакансий
func GetVacanciesIds(config *configuration.Config, client httphandler.HttpClient) ([]int64, error) {
	body, err := client.Get(config.HeadHunter.ApiUrl+"/vacancies", &config.UrlParameters)
	if err != nil {
		return nil, fmt.Errorf("ошибка получения тела ответа: %v", err)
	}

	vacancies, err := serializer.Deserialize[VacancyResponse](body)
	if err != nil {
		return nil, fmt.Errorf("ошибка демаршалинга: %v", err)
	}

	var newVacanciesIds []int64

	for _, v := range vacancies.Items {
		intId, err := strconv.Atoi(v.Id)
		if err != nil {
			return nil, fmt.Errorf("ошибка преобразования ID в int: %v", err)
		}

		newVacanciesIds = append(newVacanciesIds, int64(intId))
	}

	return newVacanciesIds, nil
}
