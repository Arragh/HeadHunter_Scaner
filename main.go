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

	oldVacancies, err := getOldVacancies()
	if err != nil {
		fmt.Printf("Ошибка получения ID старых вакансий: %v\n", err)
		panic(err)
	}

	newVacancies, err := getNewVacancies()
	if err != nil {
		fmt.Printf("Ошибка получения новых вакансий: %v\n", err)
		panic(err)
	}

	var dif = difference(*oldVacancies, *newVacancies)
	if len(dif) > 0 {
		fmt.Println("============================")
		fmt.Println("Имеются новые вакансии:")

		for i := range dif {
			fmt.Println(dif[i].Url)
		}

		fmt.Println("============================")
	}

	var meshedVacancies = meshOldAndNewVacancies(*oldVacancies, *newVacancies)

	err = saveDataToJsonFile(meshedVacancies, "output.json")
	if err != nil {
		fmt.Printf("Ошибка сохранения данных: %v\n", err)
		panic(err)
	}
}

func meshOldAndNewVacancies(newVacancies, oldVacancies []model.Vacancy) *[]model.Vacancy {
	var result []model.Vacancy

	seen := make(map[string]bool)
	for _, v := range append(oldVacancies, newVacancies...) {
		if !seen[v.Id] {
			seen[v.Id] = true
			result = append(result, v)
		}
	}

	return &result
}

func difference(newVacancies, oldVacancies []model.Vacancy) []model.Vacancy {
	// Создаём map для быстрой проверки наличия в a
	inA := make(map[string]bool)
	for _, v := range newVacancies {
		inA[v.Id] = true
	}

	var result []model.Vacancy

	for _, v := range oldVacancies {
		if !inA[v.Id] { // если элемента нет в a
			result = append(result, v)
		}
	}

	return result
}

func getOldVacancies() (*[]model.Vacancy, error) {
	oldVacancies, err := readDataFromJsonFile("output.json")
	if err != nil {
		return nil, fmt.Errorf("ошибка чтения данных: %v", err)
	}

	return oldVacancies, nil
}

func getNewVacancies() (*[]model.Vacancy, error) {
	var baseUrl = "https://api.hh.ru/vacancies"

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

func getVacanciesIds(vacancies *[]model.Vacancy) *[]string {
	var vacanciesIds []string
	for _, vacancy := range *vacancies {
		vacanciesIds = append(vacanciesIds, vacancy.Id)
	}

	return &vacanciesIds
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

func readDataFromJsonFile(filename string) (*[]model.Vacancy, error) {
	_, err := os.Stat(filename)
	if err != nil && os.IsNotExist(err) {
		err = os.WriteFile(filename, []byte(`[]`), 0644)
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

	var unpacked []model.Vacancy

	err = json.Unmarshal(byteData, &unpacked)
	if err != nil {
		return nil, fmt.Errorf("ошибка демаршалинга: %v", err)
	}

	return &unpacked, nil
}

func saveDataToJsonFile(data *[]model.Vacancy, filename string) error {
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
