package dao

import (
	"resource-server/internal/models"

	"gorm.io/gorm"
)

type UrlDaoImpl struct {
	db *gorm.DB
}

func NewUrlDaoImpl(db *gorm.DB) UrlDao {
	return &UrlDaoImpl{
		db: db,
	}
}

func (d *UrlDaoImpl) CreateUrl(url *models.Url) error {
	return d.db.Create(url).Error
}

func (d *UrlDaoImpl) GetUrlByCode(email string) (*models.Url, error) {
	var url *models.Url
	if err := d.db.Where("code = ?", email).First(url).Error; err != nil {
		return nil, err
	}
	return url, nil
}

func (d *UrlDaoImpl) DeleteUrl(code string) error {
	// 删除条件：主键等于 code
	if err := d.db.Delete(&models.Url{}, code).Error; err != nil {
		return err // 返回错误信息
	}
	return nil // 删除成功
}
