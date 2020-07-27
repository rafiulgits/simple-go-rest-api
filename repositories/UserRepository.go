package repositories

import (
	"restapi/conn"
	"restapi/models"
)

//IUserRepository :
type IUserRepository interface {
	Get(id uint) (*models.User, error)
	GetAll() ([]*models.User, error)
	Create(u *models.User) (*models.User, error)
	Delete(id uint) error
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
	var user models.User
	err := repo.db.Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
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

//Create :
func (repo *UserRepository) Create(u *models.User) (*models.User, error) {
	var err error
	if repo.db.NewRecord(u) {
		repo.db.Create(&u)
		if !repo.db.NewRecord(u) {
			return u, nil
		}
		return nil, err
	}
	return nil, err
}

//Delete :
func (repo *UserRepository) Delete(id uint) error {
	var user models.User
	err := repo.db.Where("id = ?", id).First(&user).Error
	if err != nil {
		return err
	}
	e := repo.db.Delete(&user).Error
	return e
}
