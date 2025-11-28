package mocks

import "hhscaner/configuration"

type MockHttpClient struct {
	Response []byte
	Error    error
}

func (m *MockHttpClient) Get(url string, params *[]configuration.UrlParameter) ([]byte, error) {
	return m.Response, m.Error
}
