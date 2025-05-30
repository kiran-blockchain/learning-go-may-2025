package services

import (
    "rest-api/models"
     "rest-api/interfaces"
)

type UserService interface {
    GetAllUsers() []models.User
}

type userService struct {
    repo interfaces.UserRepository
}

//pass the repository dependency 
func NewUserService(r interfaces.UserRepository) UserService {
    return &userService{repo: r}
}

func (s *userService) GetAllUsers() []models.User {
    return s.repo.GetAllUsers()
}

