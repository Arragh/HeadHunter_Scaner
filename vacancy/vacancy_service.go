package vacancy

import "HeadHunter_Scaner/model"

func MergeVacancies(newVacancies, oldVacancies []model.Vacancy) *[]model.Vacancy {
	var result []model.Vacancy

	seen := make(map[string]bool)
	for _, v := range append(oldVacancies, newVacancies...) {
		if !seen[v.Id] {
			seen[v.Id] = true
			result = append(result, v)
		}
	}

	return &result
}

func Difference(newVacancies, oldVacancies []model.Vacancy) []model.Vacancy {
	inA := make(map[string]bool)
	for _, v := range newVacancies {
		inA[v.Id] = true
	}

	var result []model.Vacancy

	for _, v := range oldVacancies {
		if !inA[v.Id] {
			result = append(result, v)
		}
	}

	return result
}
