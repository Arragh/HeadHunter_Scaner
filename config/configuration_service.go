package config

import (
	"HeadHunter_Scaner/model"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func LoadConfigurartion() (*model.Config, error) {
	file, err := os.Open("config.json")
	if err != nil {
		return nil, fmt.Errorf("ошибка открытия файла: %v", err)
	}
	defer file.Close()

	byteData, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("ошибка чтения файла: %v", err)
	}

	var unpacked model.Config
	err = json.Unmarshal(byteData, &unpacked)
	if err != nil {
		return nil, fmt.Errorf("ошибка демаршалинга: %v", err)
	}

	return &unpacked, nil
}
