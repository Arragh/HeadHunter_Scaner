package serializer

import (
	"encoding/json"
	"fmt"
)

// Deserialize десерилизует массив байт в структуру
func Deserialize[T any](body []byte) (*T, error) {
	var unpacked T

	err := json.Unmarshal(body, &unpacked)
	if err != nil {
		return nil, fmt.Errorf("ошибка демаршалинга: %v", err)
	}

	return &unpacked, nil
}
