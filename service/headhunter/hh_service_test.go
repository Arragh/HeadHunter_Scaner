package headhunter

import (
	"encoding/json"
	"hhscaner/configuration"
	"hhscaner/test/mock"
	"reflect"
	"strconv"
	"testing"

	"github.com/golang/mock/gomock"
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
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	want := []int64{1, 2}
	vacancyResponse := VacancyResponse{}
	for _, v := range want {
		vacancyResponse.Items = append(vacancyResponse.Items, Vacancy{
			Id: strconv.Itoa(int(v)),
		})
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

	body, _ := json.Marshal(vacancyResponse)

	mockHttpClient := mock.NewMockHttpClient(ctrl)
	mockHttpClient.
		EXPECT().
		Get(gomock.Any(), gomock.Any()).
		Return(body, nil)

	got, err := GetVacanciesIds(&config, mockHttpClient)
	if err != nil {
		t.Errorf("ошибка при вызове метода GetVacanciesIds(): %s", err)
	}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("GetVacanciesIds() = \"%v\", want %v", got, want)
	}
}
