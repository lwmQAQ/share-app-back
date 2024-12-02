package dao

import (
	"login-server/internal/models"

	"gorm.io/gorm"
)

type UserDaoImpl struct {
	db *gorm.DB
}

func NewUserDaoImpl(db *gorm.DB) UserDao {
	return &UserDaoImpl{
		db: db,
	}
}

func (d *UserDaoImpl) GetUserByEmail(email string) (*models.User, error) {
	var user *models.User
	if err := d.db.Where("email = ?", email).First(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (d *UserDaoImpl) CreateUser(user *models.User) error {
	return d.db.Create(user).Error
}
