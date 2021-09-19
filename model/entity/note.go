package entity

type Note struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Content string `json:"content"`
}
