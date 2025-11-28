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
func (c *DefaultHttpClient) Get(baseUrl string, params *[]configuration.UrlParameter) ([]byte, error) {
	buildedUrl := baseUrl
	if params != nil {
		tempUrl, err := buildUrl(baseUrl, params)
		if err != nil {
			return nil, fmt.Errorf("ошибка построения URL: %v", err)
		}

		buildedUrl = tempUrl
	}

	resp, err := http.Get(buildedUrl)
	if err != nil {
		fmt.Println("Не удалось получить ответ от удаленного сервера:", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Удаленный сервер ответил с ошибкой:", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("ошибка чтения тела ответа: %v", err)
	}

	return body, nil
}

// buildUrl собирает URL с учетом переданных параметров
func buildUrl(urlString string, params *[]configuration.UrlParameter) (string, error) {
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
