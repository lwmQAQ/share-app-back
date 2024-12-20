package dao

import (
	"errors"
	"fmt"
	"user-server/internal/models"

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
	var user = new(models.User)
	if err := d.db.Where("email = ?", email).First(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (d *UserDaoImpl) CreateUser(user *models.User) error {
	return d.db.Create(user).Error
}

func (d *UserDaoImpl) SelectUserById(id uint64) (*models.User, error) {
	var user models.User

	if err := d.db.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (d *UserDaoImpl) UserUpdate(updates map[string]interface{}, id uint64) error {
	// 检查是否有需要更新的字段
	if len(updates) == 0 {
		return errors.New("no fields to update")
	}

	// 执行更新操作，只更新不为空的字段
	result := d.db.Model(&models.User{}).Where("id = ?", id).Updates(updates)
	// 检查是否有错误
	if result.Error != nil {
		return fmt.Errorf("更新失败：%s", result.Error)
	}
	return nil
}
