package headhunter

type VacancyResponse struct {
	Items []Vacancy `json:"items"`
}

type Vacancy struct {
	Id                     string     `json:"id"`
	Name                   string     `json:"name"`
	HasTest                bool       `json:"has_test"`
	ResponseLetterRequired bool       `json:"response_letter_required"`
	Url                    string     `json:"alternate_url"`
	Department             Department `json:"department"`
	Salary                 Salary     `json:"salary"`
}

type Department struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Salary struct {
	From     float64 `json:"from"`
	To       float64 `json:"to"`
	Currency string  `json:"currency"`
	Gross    bool    `json:"gross"`
}
