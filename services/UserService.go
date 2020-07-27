package services

import (
	"restapi/models"
	"restapi/repositories"
)

//IUserService :
type IUserService interface {
	GetUserByID(id uint) (*models.User, error)
	GetAll() ([]*models.User, error)
	CreateUser(u *models.User) (*models.User, error)
	DeleteUser(id uint) error
}

//UserService :
type UserService struct {
	UserRepository repositories.IUserRepository
}

//NewUserService :
func NewUserService(userRepository repositories.IUserRepository) IUserService {
	return &UserService{
		UserRepository: userRepository,
	}
}

//GetUserByID :
func (service *UserService) GetUserByID(id uint) (*models.User, error) {
	return service.UserRepository.Get(id)
}

//GetAll :
func (service *UserService) GetAll() ([]*models.User, error) {
	return service.UserRepository.GetAll()
}

//CreateUser :
func (service *UserService) CreateUser(u *models.User) (*models.User, error) {
	return service.UserRepository.Create(u)
}

//DeleteUser :
func (service *UserService) DeleteUser(id uint) error {
	return service.UserRepository.Delete(id)
}
