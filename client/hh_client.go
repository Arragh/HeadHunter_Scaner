package client

import (
	"HeadHunter_Scaner/config"
	"HeadHunter_Scaner/handler"
	"encoding/json"
	"fmt"
)

func FetchVacancies(config *config.Config) (*[]Vacancy, error) {
	body, err := handler.Get(config.BaseUrl, &config.UrlParameters)
	if err != nil {
		return nil, fmt.Errorf("ошибка получения тела ответа: %v", err)
	}

	newVacancies, err := deserializevacanciesFromBody(body)
	if err != nil {
		return nil, fmt.Errorf("ошибка демаршалинга: %v", err)
	}

	return &newVacancies.Items, nil
}

func deserializevacanciesFromBody(body []byte) (*VacancyResponse, error) {
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
