package models

import (
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRole string

const (
	RoleAdmin     UserRole = "admin"
	RoleUser      UserRole = "user"
	RoleManager   UserRole = "manager"
	RoleSeller    UserRole = "seller"
	RoleSupporter UserRole = "supporter"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	FirstName string    `gorm:"type:varchar(100)"`
	LastName  string    `gorm:"type:varchar(100)"`
	Email     string    `gorm:"uniqueIndex;not null"`
	Password  string
	Phone     *string   `gorm:"type:varchar(15);uniqueIndex"`
	Role      UserRole `gorm:"type:varchar(20);not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type UserInput struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email" binding:"required,email"`
	Phone     string `json:"phone" binding:"min=10,max=15"`
	Password  string `json:"password" binding:"required,min=8"`
}

type UserOutput struct {
	ID        string    `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	Role      UserRole  `json:"role"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (u *User) SetPassword(raw string) error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(raw), 12)
	if err != nil {
		return  err
	}

	u.Password = string(hashed)
	return nil
}

func (u *User) CheckPassword(raw string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(raw))
	return err == nil
}

func (u *User) ToOutput() UserOutput {
	var phone string
	if u.Phone == nil {
		phone = ""
	} else {
		phone = *u.Phone
	}
	return UserOutput{
		ID: u.ID.String(),
		FirstName: u.FirstName,
		LastName: u.LastName,
		Email: u.Email,
		Phone: phone,
		Role: u.Role,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}
