package interfaces

import (
	"go-shop-api/internal/models"

	"github.com/google/uuid"
)

type UserRepository interface {
	CreateUser(user *models.User) error
	
	GetUserByID(id uuid.UUID) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
	GetUserByPhone(phone string) (*models.User, error)

	EmailExists(email string) (bool, error)
	PhoneExists(phone string) (bool, error)

	UpdateUser(user *models.User) error

	DeleteUser(id uuid.UUID) error

	ListUsers(offset int, limit int) ([]models.User, error)
	CountUsers() (int64, error)
	WithTransaction(fn func(repo UserRepository) error) error
}