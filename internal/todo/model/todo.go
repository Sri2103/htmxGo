package model

type Todo struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"desc"`
	Status      bool   `json:"status"`
	UserID      int    `json:"user_id"`
}
