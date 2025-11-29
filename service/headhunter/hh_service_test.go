package headhunter

import (
	"reflect"
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

func TestParseVacanciesIds(t *testing.T) {
	vacancies := []Vacancy{
		{Id: "123"},
		{Id: "456"},
		{Id: "789"},
	}

	got, err := ParseVacanciesIds(vacancies)
	if err != nil {
		t.Errorf("ошибка при парсинге вакансий: %v", err)
	}

	want := []int64{123, 456, 789}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v, want: %v", got, want)
	}
}
