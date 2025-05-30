package interfaces

import(
	"rest-api/models"
)
type UserRepository interface {
	GetAllUsers() []models.User
	GetUserById() models.User
}