package repositories

import (
	"github.com/jinzhu/gorm"
)

//BaseRepository :
type BaseRepository struct {
	db *gorm.DB
}
