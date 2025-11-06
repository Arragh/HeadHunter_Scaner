package main

import (
	"HeadHunter_Scaner/model"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	fmt.Println(
		"👋 Hello, World!",
	)

	var url = "https://api.hh.ru/vacancies"

	SetUrlRarams(&url)

	body, err := GetHttpResponseBody(url)
	if err != nil {
		fmt.Printf("Ошибка запроса: %v\n", err)
		return
	}

	deserializedBody, err := DeserializeHttpResponseBody(body)
	if err != nil {
		fmt.Printf("Ошибка демаршалинга: %v\n", err)
	}

	err = SaveDataToJsonFile(deserializedBody, "output.json")
	if err != nil {
		fmt.Printf("Ошибка сохранения данных: %v\n", err)
	}
}

func SetUrlRarams(url *string) {
	var area = fmt.Sprintf("area=%d", 113)
	var period = fmt.Sprintf("period=%d", 30)
	var workFormat = fmt.Sprintf("work_format=%s", "REMOTE")
	var searchField = fmt.Sprintf("search_field=%s", "name")
	var includeWords = fmt.Sprintf("text=%s", "C%23")
	var excludeWords = fmt.Sprintf("excluded_text=%s", "QA,AQA")

	var params = "?" + area + "&" + period + "&" + workFormat + "&" + searchField + "&" + includeWords + "&" + excludeWords

	*url += params
}

func GetHttpResponseBody(url string) ([]byte, error) {
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

func DeserializeHttpResponseBody(body []byte) (*model.VacancyResponse, error) {
	var unpacked model.VacancyResponse

	err := json.Unmarshal(body, &unpacked)
	if err != nil {
		return nil, fmt.Errorf("ошибка демаршалинга: %v", err)
	}

	return &unpacked, nil
}

func SaveDataToJsonFile(data *model.VacancyResponse, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("ошибка создания файла: %v", err)
	}
	defer file.Close()

	indented, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return fmt.Errorf("ошибка форматирования данных: %v", err)
	}

	_, err = file.Write(indented)
	if err != nil {
		return fmt.Errorf("ошибка записи данных в файл: %v", err)
	}

	fmt.Println("Данные сохранены! 🎉")

	return nil
}
