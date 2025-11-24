// Package configuration содержит структуры и методы для хранения и получения конфигурации
package configuration

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

// GetConfigurartion загружает конфигурацию из файла
func GetConfigurartion() (*Config, error) {
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
	RequestIntervalInSeconds int            `json:"requestIntervalInSeconds"`
	HeadHunter               HeadHunter     `json:"headHunter"`
	Telegram                 Telegram       `json:"telegram"`
	UrlParameters            []UrlParameter `json:"urlParameters"`
}

type HeadHunter struct {
	BaseUrl string `json:"baseUrl"`
	ApiUrl  string `json:"apiUrl"`
}

type Telegram struct {
	ApiUrl   string `json:"apiUrl"`
	BotToken string `json:"botToken"`
	ChatId   string `json:"chatId"`
}

type UrlParameter struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
