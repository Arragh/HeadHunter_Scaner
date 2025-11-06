package model

type Salary struct {
	From     float64 `json:"from"`
	To       float64 `json:"to"`
	Currency string  `json:"currency"`
	Gross    bool    `json:"gross"`
}
