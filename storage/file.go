package storage

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func ReadDataFromFile(filename string) (*[]int, error) {
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

	var unpacked []int

	err = json.Unmarshal(byteData, &unpacked)
	if err != nil {
		return nil, fmt.Errorf("ошибка демаршалинга: %v", err)
	}

	return &unpacked, nil
}

func SaveDataToFile(data *[]int, filename string) error {
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

	return nil
}
