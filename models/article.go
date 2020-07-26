package models

//Article domain model
type Article struct {
	ID       uint   `gorm:"primary_key" json:"id"`
	Title    string `gorm:"type:varchar(250);not null" json:"title"`
	Body     string `gorm:"type:text" json:"body"`
	Author   User   `gorm:"foreignkey:AuthorID"`
	AuthorID uint   `json:"authorId"`
}

//TableName for Article model
func (Article) TableName() string {
	return "Articles"
}
