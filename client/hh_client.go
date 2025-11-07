package client

import (
	"HeadHunter_Scaner/model"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func FetchVacancies(baseUrl string) (*[]model.Vacancy, error) {
	buildedUrl, err := buildUrl(baseUrl)
	if err != nil {
		return nil, fmt.Errorf("ошибка построения URL: %v", err)
	}

	body, err := getHttpResponseBody(buildedUrl)
	if err != nil {
		return nil, fmt.Errorf("ошибка получения тела ответа: %v", err)
	}

	newVacancies, err := deserializeHttpResponseBody(body)
	if err != nil {
		return nil, fmt.Errorf("ошибка демаршалинга: %v", err)
	}

	return &newVacancies.Items, nil
}

func buildUrl(baseUrl string) (string, error) {
	parsedUrl, err := url.Parse(baseUrl)
	if err != nil {
		return "", fmt.Errorf("ошибка парсинга URL: %v", err)
	}

	params := url.Values{}
	params.Add("area", "113")
	params.Add("period", "30")
	params.Add("work_format", "REMOTE")
	params.Add("search_field", "name")
	params.Add("text", "C#")
	params.Add("excluded_text", "QA,AQA")

	parsedUrl.RawQuery = params.Encode()

	return parsedUrl.String(), nil
}

func getHttpResponseBody(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("ошибка запроса: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("ошибка статуса ответа: %v", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("ошибка чтения тела ответа: %v", err)
	}

	return body, nil
}

func deserializeHttpResponseBody(body []byte) (*model.VacancyResponse, error) {
	var unpacked model.VacancyResponse

	err := json.Unmarshal(body, &unpacked)
	if err != nil {
		return nil, fmt.Errorf("ошибка демаршалинга: %v", err)
	}

	return &unpacked, nil
}
