package model

type Category struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	UserId      int    `json:"userId"`
	Icon        string `json:"icon"`
	Description string `json:"description"`
	Color       string `json:"color"`
}

type Expense struct {
	Type        string  `json:"type"`
	Description string  `json:"description"`
	Amount      float64 `json:"amount"`
	Date        string  `json:"date"`
	Location    string  `json:"location"`
	UserId      int     `json:"userId"`
	CategoryID  int     `json:"categoryId"`
}
