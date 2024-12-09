package model

type Category struct {
	ID       int       `json:"id"`
	Name     string    `json:"name"`
	Products []Product `json:"products"`
}
