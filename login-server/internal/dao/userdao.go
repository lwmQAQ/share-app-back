package dao

import "login-server/internal/models"

type UserDao interface {
	CreateUser(user *models.User) error
	GetUserByEmail(email string) (*models.User, error)
}
