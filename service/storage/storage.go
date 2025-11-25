// Package storage реализует функции для чтения и записи данных в файл
package storage

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// CheckFileExists проверяет, существует ли файл, и если нет, то создает его
func CheckFileExists(fileName string) error {
	_, err := os.Stat(fileName)
	if err != nil && os.IsNotExist(err) {
		err = os.WriteFile(fileName, []byte(""), 0644)
		if err != nil {
			return fmt.Errorf("ошибка создания файла %s: %v", fileName, err)
		}
	}

	return nil
}

// ReadData читает данные из файла и возвращает их в виде среза int64
func ReadData(fileName string) ([]int64, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var numbers []int64

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		txtValue := strings.TrimSpace(scanner.Text())

		if txtValue == "" {
			continue
		}

		intValue, err := strconv.ParseInt(txtValue, 10, 64)
		if err != nil {
			return nil, err
		}

		numbers = append(numbers, intValue)
	}

	if scanner.Err() != nil {
		return nil, scanner.Err()
	}

	return numbers, nil
}

// SaveData сохраняет данные в файл
func SaveData(data []int64, filename string) error {
	options := os.O_APPEND | os.O_WRONLY
	file, err := os.OpenFile(filename, options, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, value := range data {
		_, err := fmt.Fprintln(file, value)
		if err != nil {
			return err
		}
	}

	return nil
}
