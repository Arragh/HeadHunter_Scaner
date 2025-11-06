package model

type Vacancy struct {
	Id                     string     `json:"id"`
	Name                   string     `json:"name"`
	HasTest                bool       `json:"has_test"`
	ResponseLetterRequired bool       `json:"response_letter_required"`
	Url                    string     `json:"alternate_url"`
	Department             Department `json:"department"`
	Salary                 Salary     `json:"salary"`
}
