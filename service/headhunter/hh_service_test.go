package headhunter

import (
	"encoding/json"
	"hhscaner/configuration"
	"hhscaner/test/mocks"
	"reflect"
	"strconv"
	"testing"
)

func TestDifference(t *testing.T) {
	newVacanciesIds := []int64{1, 2, 3, 4, 5}
	oldVacanciesIds := []int64{1, 2, 3}
	want := []int64{4, 5}

	got := Difference(newVacanciesIds, oldVacanciesIds)

	if !reflect.DeepEqual(want, got) {
		t.Errorf("Difference() = \"%v\", want %v", got, want)
	}
}

func TestGetVacanciesIds(t *testing.T) {
	want := []int64{1, 2}

	vacancyResponse := VacancyResponse{}

	for _, v := range want {
		vacancyResponse.Items = append(vacancyResponse.Items, Vacancy{
			Id: strconv.Itoa(int(v)),
		})
	}

	body, _ := json.Marshal(vacancyResponse)

	mockHttpClient := mocks.MockHttpClient{
		Response: body,
		Error:    nil,
	}

	config := configuration.Config{
		RequestIntervalInSeconds: 1,
		HeadHunter: configuration.HeadHunter{
			ApiUrl: "https://mock.mock",
		},
		UrlParameters: []configuration.UrlParameter{
			{
				Key:   "text",
				Value: "golang",
			},
		},
	}

	got, err := GetVacanciesIds(&config, &mockHttpClient)
	if err != nil {
		t.Errorf("ошибка при вызове метода GetVacanciesIds(): %s", err)
	}

	if !reflect.DeepEqual(want, got) {
		t.Errorf("GetVacanciesIds() = \"%v\", want %v", got, want)
	}
}
