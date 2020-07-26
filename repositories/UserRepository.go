package repositories

import (
	"restapi/conn"
	"restapi/models"
)

//IUserRepository :
type IUserRepository interface {
	Get(id uint) (*models.User, error)
	GetAll() ([]*models.User, error)
}

//UserRepository :
type UserRepository struct {
	*BaseRepository
}

//NewUserRepository :
func NewUserRepository(db *conn.DB) IUserRepository {
	return &UserRepository{
		&BaseRepository{
			db: db.Table("Users"),
		},
	}
}

//Get :
func (repo *UserRepository) Get(id uint) (*models.User, error) {
	return nil, nil
}

//GetAll :
func (repo *UserRepository) GetAll() ([]*models.User, error) {
	var users []*models.User
	err := repo.db.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}
