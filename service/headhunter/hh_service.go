// Package headhunter предоставляет функции для работы с api hh.ru и вакансиями
package headhunter

import (
	"fmt"
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
func ParseVacanciesIds(vacancies []Vacancy) ([]int64, error) {
	var newVacanciesIds []int64

	for _, v := range vacancies {
		intId, err := strconv.Atoi(v.Id)
		if err != nil {
			return nil, fmt.Errorf("ошибка преобразования ID в int: %v", err)
		}

		newVacanciesIds = append(newVacanciesIds, int64(intId))
	}

	return newVacanciesIds, nil
}
