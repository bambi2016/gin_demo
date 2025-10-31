package service

import (
	"github.com/example/user-api/model"
	"github.com/example/user-api/repository"
)

// UserService defines the interface for user business logic
type UserService interface {
	CreateUser(user *model.User) error
	GetUserByID(id uint64) (*model.User, error)
	GetAllUsers() ([]*model.User, error)
	UpdateUser(user *model.User) error
	DeleteUser(id uint64) error
}

// userService implements UserService
type userService struct {
	userRepo repository.UserRepository
}

// NewUserService creates a new instance of UserService
func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

// CreateUser adds a new user
func (s *userService) CreateUser(user *model.User) error {
	// Business logic can be added here (e.g., validation, duplicate checks)
	return s.userRepo.Create(user)
}

// GetUserByID retrieves a user by ID
func (s *userService) GetUserByID(id uint64) (*model.User, error) {
	return s.userRepo.GetByID(id)
}

// GetAllUsers retrieves all users
func (s *userService) GetAllUsers() ([]*model.User, error) {
	return s.userRepo.GetAll()
}

// UpdateUser updates an existing user
func (s *userService) UpdateUser(user *model.User) error {
	// Business logic can be added here (e.g., validation, permission checks)
	return s.userRepo.Update(user)
}

// DeleteUser removes a user by ID
func (s *userService) DeleteUser(id uint64) error {
	return s.userRepo.Delete(id)
}
