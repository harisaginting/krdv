package model

type RequestPage struct {
	Sort  string `json:"sort"`
	Order string `json:"order"`
	Page  int    `json:"page"`
	Limit int    `json:"limit"`
}
