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
