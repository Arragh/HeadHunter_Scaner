package httphandler

import (
	"fmt"
	"hhscaner/configuration"
	"io"
	"net/http"
	"net/url"
)

type DefaultHttpClient struct{}

// Get отправляет GET-запрос на указанный URL и возвращает тело ответа
func (c *DefaultHttpClient) Get(urlString string) ([]byte, error) {
	resp, err := http.Get(urlString)
	if err != nil {
		return nil, fmt.Errorf("не удалось получить ответ от удаленного сервера: %s", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("удаленный сервер ответил с ошибкой: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("ошибка чтения тела ответа: %v", err)
	}

	return body, nil
}

// BuildUrl собирает URL с учетом переданных параметров
func BuildUrl(urlString string, params *[]configuration.UrlParameter) (string, error) {
	parsedUrl, err := url.Parse(urlString)
	if err != nil {
		return "", fmt.Errorf("ошибка парсинга URL: %v", err)
	}

	if *params != nil {
		parsedParams := url.Values{}
		for _, param := range *params {
			if param.Value != "" {
				parsedParams.Add(param.Key, param.Value)
			}
		}

		parsedUrl.RawQuery = parsedParams.Encode()
	}

	return parsedUrl.String(), nil
}
