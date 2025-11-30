package httphandler

import (
	"hhscaner/configuration"
	"testing"
)

func TestBuildUrl_Valid(t *testing.T) {
	urlString := "https://mock.mock"
	params := []configuration.UrlParameter{
		{
			Key:   "key1",
			Value: "value1",
		},
		{
			Key:   "key2",
			Value: "value2",
		},
		{
			Key:   "key3",
			Value: "value3",
		},
	}

	want := "https://mock.mock?key1=value1&key2=value2&key3=value3"

	got, err := BuildUrl(urlString, &params)
	if err != nil {
		t.Errorf("ошибка построения Url: %v", err)
	}

	if got != want {
		t.Errorf("got = %v, want %v", got, want)
	}
}

func TestBuildUrl_InvalidUrl(t *testing.T) {
	urlString := "https:// mock .mock"
	params := []configuration.UrlParameter{
		{
			Key:   "key1",
			Value: "value1",
		},
		{
			Key:   "key2",
			Value: "value2",
		},
		{
			Key:   "key3",
			Value: "value3",
		},
	}

	got, err := BuildUrl(urlString, &params)
	if err == nil {
		t.Errorf("ожидалась ошибка построения Url, got: %v", got)
	}
}
