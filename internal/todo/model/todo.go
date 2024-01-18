package model

type Todo struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"desc"`
	Status      string `json:"status"`
}
