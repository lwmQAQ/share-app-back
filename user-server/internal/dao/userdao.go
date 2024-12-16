package dao

import "user-server/internal/models"

type UserDao interface {
	CreateUser(user *models.User) error
	GetUserByEmail(email string) (*models.User, error)
	SelectUserById(id uint64) (*models.User, error)
	UserUpdate(updates map[string]interface{}, id uint64) error
}
