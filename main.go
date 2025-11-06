package main

import (
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

	var url = "https://api.hh.ru/dictionaries"

	body, err := GetHttpResponseBody(url)
	if err != nil {
		fmt.Printf("Ошибка запроса: %v\n", err)
		return
	}

	// resp, err := http.Get(url)
	// if err != nil {
	// 	fmt.Printf("Ошибка запроса: :%v\n", err)
	// 	return
	// }
	// defer resp.Body.Close()

	// if resp.StatusCode != http.StatusOK {
	// 	fmt.Printf("Ошибка статуса ответа: %v\n", resp.Status)
	// }

	// body, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	fmt.Printf("Ошибка чтения тела ответа: %v\n", err)
	// }

	var unpacked map[string][]DictValue
	err = json.Unmarshal(body, &unpacked)
	if err != nil {
		fmt.Printf("Ошибка демаршалинга: %v\n", err)
	}

	// fmt.Println(unpacked)

	file, err := os.Create("output.json")
	if err != nil {
		fmt.Printf("Ошибка создания файла: %v\n", err)
	}
	defer file.Close()

	indented, err := json.MarshalIndent(unpacked, "", "  ")
	if err != nil {
		fmt.Printf("Ошибка форматирования: %v\n", err)
	}

	_, err = file.Write(indented)
	if err != nil {
		fmt.Printf("Ошибка записи в файл: %v\n", err)
	}

	fmt.Println("Данные сохранены! 🎉")
}

type DictValue struct {
	Id   string `json:"id"`
	Name string `json:"name"`
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
