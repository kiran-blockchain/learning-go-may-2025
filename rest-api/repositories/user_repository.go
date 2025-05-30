package repositories

import (
	"rest-api/models"
	"rest-api/interfaces"
)

type userRepo struct{}

func (r *userRepo) GetUserById() (models.User) {
	// In a real project, this would interact with the DB.
	return models.User{ID: 1, Name: "Alice"}
	
}

func (r *userRepo) GetAllUsers()([]models.User){
	// In a real project, this would interact with the DB.
	return []models.User{
		{ID: 1, Name: "Alice"},
		{ID: 2, Name: "Bob"},
	}
}

func NewUserRepository() interfaces.UserRepository {
	return &userRepo{}
}


