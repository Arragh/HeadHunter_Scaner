package headhunter

import (
	"encoding/json"
	"fmt"
)

// deserializeBody делает демаршалинг тела ответа
func deserializeBody(body []byte) (*VacancyResponse, error) {
	var unpacked VacancyResponse

	err := json.Unmarshal(body, &unpacked)
	if err != nil {
		return nil, fmt.Errorf("ошибка демаршалинга: %v", err)
	}

	return &unpacked, nil
}
