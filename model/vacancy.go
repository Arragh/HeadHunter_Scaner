package model

type Vacancy struct {
	Id                       string     `json:"id"`
	name                     string     `json:"name"`
	has_test                 bool       `json:"has_test"`
	response_letter_required bool       `json:"response_letter_required"`
	alternate_url            string     `json:"alternate_url"`
	Department               Department `json:"department"`
	Salary                   Salary     `json:"salary"`
}
