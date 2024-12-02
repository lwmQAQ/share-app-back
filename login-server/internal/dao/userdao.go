package dao

import "login-server/internal/models"

type UserDao interface {
	GetUserByEmail(email string) (*models.User, error)
}
