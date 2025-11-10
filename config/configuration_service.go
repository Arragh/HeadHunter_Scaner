package config

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func LoadConfigurartion() (*Config, error) {
	file, err := os.Open("config.json")
	if err != nil {
		return nil, fmt.Errorf("ошибка открытия файла: %v", err)
	}
	defer file.Close()

	byteData, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("ошибка чтения файла: %v", err)
	}

	var unpacked Config
	err = json.Unmarshal(byteData, &unpacked)
	if err != nil {
		return nil, fmt.Errorf("ошибка демаршалинга: %v", err)
	}

	return &unpacked, nil
}

type Config struct {
	BaseUrl                  string         `json:"baseUrl"`
	RequestIntervalInSeconds int            `json:"requestIntervalInSeconds"`
	UrlParameters            []UrlParameter `json:"urlParameters"`
}

type UrlParameter struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
