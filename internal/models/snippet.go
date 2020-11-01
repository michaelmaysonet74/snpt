package models

type Snippet struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	IsLoved bool   `json:"isLoved"`
}
