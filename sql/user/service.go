package user

import "github.com/google/uuid"

// Service is capable of basic users operations.
type Service interface {
	InsertUser(u User) error
	GetAllUsers() ([]User, error)
	GetUser(id uuid.UUID) (User, error)
	DeleteAllUsers() error
}
