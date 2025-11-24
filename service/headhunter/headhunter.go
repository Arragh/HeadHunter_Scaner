// Package headhunter предоставляет функции для работы с api hh.ru и вакансиями
package headhunter

import (
	"encoding/json"
	"fmt"
	"hhscaner/configuration"
	"hhscaner/service/httphandler"
	"strconv"
)

// Difference сравнивает полученные вакансии с сохраненными и возвращает разницу
func Difference(newVacanciesIds []int64, oldVacanciesIds []int64) ([]int64, error) {
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

	return result, nil
}

// GetVacanciesIds обращается к api hh.ru и возвращает список ID вакансий
func GetVacanciesIds(config *configuration.Config) ([]int64, error) {
	body, err := httphandler.Get(config.HeadHunter.ApiUrl+"/vacancies", &config.UrlParameters)
	if err != nil {
		return nil, fmt.Errorf("ошибка получения тела ответа: %v", err)
	}

	vacancies, err := deserializeBody(body)
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

// deserializeBody делает демаршалинг тела ответа
func deserializeBody(body []byte) (*VacancyResponse, error) {
	var unpacked VacancyResponse

	err := json.Unmarshal(body, &unpacked)
	if err != nil {
		return nil, fmt.Errorf("ошибка демаршалинга: %v", err)
	}

	return &unpacked, nil
}

type VacancyResponse struct {
	Items []Vacancy `json:"items"`
}

type Vacancy struct {
	Id                     string     `json:"id"`
	Name                   string     `json:"name"`
	HasTest                bool       `json:"has_test"`
	ResponseLetterRequired bool       `json:"response_letter_required"`
	Url                    string     `json:"alternate_url"`
	Department             Department `json:"department"`
	Salary                 Salary     `json:"salary"`
}

type Department struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Salary struct {
	From     float64 `json:"from"`
	To       float64 `json:"to"`
	Currency string  `json:"currency"`
	Gross    bool    `json:"gross"`
}
