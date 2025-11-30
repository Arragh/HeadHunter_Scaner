package headhunter

import (
	"reflect"
	"testing"
)

func TestDifference_Valid(t *testing.T) {
	newVacanciesIds := []int64{1, 2, 3, 4, 5}
	oldVacanciesIds := []int64{1, 2, 3}
	want := []int64{4, 5}

	got := Difference(newVacanciesIds, oldVacanciesIds)

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("got = \"%v\", want %v", got, want)
	}
}

func TestParseVacanciesIds_Valid(t *testing.T) {
	vacancies := []Vacancy{
		{Id: "1"},
		{Id: "2"},
		{Id: "3"},
	}

	got, err := ParseVacanciesIds(vacancies)
	if err != nil {
		t.Fatalf("ошибка при парсинге вакансий: %v", err)
	}

	want := []int64{1, 2, 3}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got: %v, want: %v", got, want)
	}
}

func TestParseVacanciesIds_InValidJsonValues(t *testing.T) {
	vacancies := []Vacancy{
		{Id: "one"},
		{Id: "two"},
		{Id: "three"},
	}

	got, err := ParseVacanciesIds(vacancies)
	if err == nil {
		t.Fatalf("ожидалась ошибка парсинга вакансий, got: %v", got)
	}
}
