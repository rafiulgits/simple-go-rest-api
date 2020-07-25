package models

//Article domain model
type Article struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Body     string `json:"body"`
	AuthorID int    `json:"authorId"`
}
