package vacancy

import (
	"HeadHunter_Scaner/model"
	"fmt"
	"strconv"
)

func MergeVacancies(oldVacanciesIds []int, newVacancies []model.Vacancy) (*[]int, error) {
	var temp []int

	for _, v := range newVacancies {
		intId, err := strconv.Atoi(v.Id)
		if err != nil {
			return nil, fmt.Errorf("ошибка преобразования ID в int: %v", err)
		}
		temp = append(temp, intId)
	}

	var result []int

	seen := make(map[int]bool)
	for _, id := range append(oldVacanciesIds, temp...) {
		if !seen[id] {
			seen[id] = true
			result = append(result, id)
		}
	}

	return &result, nil
}

func Difference(newVacancies []model.Vacancy, oldVacanciesIds []int) ([]model.Vacancy, error) {
	temp := make(map[int]bool)
	for _, id := range oldVacanciesIds {
		temp[id] = true
	}

	var result []model.Vacancy

	for _, v := range newVacancies {
		intId, err := strconv.Atoi(v.Id)
		if err != nil {
			return nil, fmt.Errorf("ошибка преобразования ID в int: %v", err)
		}

		if !temp[intId] {
			result = append(result, v)
		}
	}

	return result, nil
}
