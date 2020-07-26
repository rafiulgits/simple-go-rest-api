package services

import (
	"restapi/models"
	"restapi/repositories"
)

//IUserService :
type IUserService interface {
	GetUserByID(id uint) (*models.User, error)
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
