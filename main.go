package main

import (
	"HeadHunter_Scaner/model"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

func main() {
	fmt.Println(
		"👋 Hello, World!",
	)

	var baseUrl = "https://api.hh.ru/vacancies"

	buildedUrl, err := BuildUrl(baseUrl)
	if err != nil {
		fmt.Printf("Ошибка построения URL: %v\n", err)
		panic(err)
	}

	body, err := GetHttpResponseBody(buildedUrl)
	if err != nil {
		fmt.Printf("Ошибка запроса: %v\n", err)
		panic(err)
	}

	deserializedBody, err := DeserializeHttpResponseBody(body)
	if err != nil {
		fmt.Printf("Ошибка демаршалинга: %v\n", err)
		panic(err)
	}

	oldVacancies, err := ReadDataFromJsonFile("output.json")
	if err != nil {
		fmt.Printf("Ошибка чтения данных: %v\n", err)
		panic(err)
	}

	fmt.Println(oldVacancies)

	err = SaveDataToJsonFile(deserializedBody, "output.json")
	if err != nil {
		fmt.Printf("Ошибка сохранения данных: %v\n", err)
		panic(err)
	}
}

func BuildUrl(baseUrl string) (string, error) {
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

func ReadDataFromJsonFile(filename string) (*model.VacancyResponse, error) {
	_, err := os.Stat(filename)
	if err != nil && os.IsNotExist(err) {
		err = os.WriteFile(filename, []byte(`{"items":[]}`), 0644)
		if err != nil {
			return nil, fmt.Errorf("ошибка записи в чистый файл: %v", err)
		}
	}

	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("ошибка открытия файла: %v", err)
	}

	defer file.Close()

	byteData, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("ошибка чтения файла: %v", err)
	}

	var unpacked model.VacancyResponse

	err = json.Unmarshal(byteData, &unpacked)
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
