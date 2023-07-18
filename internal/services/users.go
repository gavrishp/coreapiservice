package services

import (
	"github.com/gavrishp/coreapiservicetest/internal/common/models"
	"github.com/gavrishp/coreapiservicetest/internal/dto"
	"github.com/gavrishp/coreapiservicetest/internal/repository"
)

type UsersService struct {
	usersRepository *repository.UsersRepository
}

type IUserService interface {
	AddUser(requestBody *dto.AddUserRequestBody)
	GetUser(userId string)
}

func NewUsersService(userRepository *repository.UsersRepository) *UsersService {
	return &UsersService{usersRepository: userRepository}
}

func (s *UsersService) AddUser(requestBody *dto.AddUserRequestBody) (*models.User, error) {

	var user models.User

	user.Name = requestBody.Name
	user.Role = requestBody.Role
	user.Description = requestBody.Description

	response, err := s.usersRepository.AddUser(&user)
	return response, err
}

func (s *UsersService) GetUser(userId string) (*models.User, error) {
	response, err := s.usersRepository.GetUser(userId)
	return response, err
}

func (s *UsersService) GetUsers() (*[]models.User, error) {
	response, err := s.usersRepository.GetUsers()
	return response, err
}

func (s *UsersService) DeleteUser(userId string) error {
	return s.usersRepository.DeleteUser(userId)
}

func (s *UsersService) UpdateUser(userId string, requestBody *dto.UpdateUserRequestBody) (*models.User, error) {
	var user models.User

	user.Name = requestBody.Name
	user.Role = requestBody.Role
	user.Description = requestBody.Description

	response, err := s.usersRepository.AddUser(&user)
	return response, err
}
