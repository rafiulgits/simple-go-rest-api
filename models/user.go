package models

//User domain model
type User struct {
	ID       uint      `gorm:"primary_key" json:"id"`
	Name     string    `gorm:"type:varchar(100);not null" json:"name"`
	Articles []Article `gorm:"foreignkey:UserID"`
}

//TableName for User model
func (User) TableName() string {
	return "User"
}
