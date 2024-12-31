package dao

import (
	"apps-server/config"
	"apps-server/internal/models"
	"apps-server/internal/mysqldb"

	"gorm.io/gorm"
)

type AppsDaoImpl struct {
	db *gorm.DB
}

func NewAppsDaoImpl(config *config.MysqlConfig) AppsDao {
	return &AppsDaoImpl{
		db: mysqldb.NewMysql(config),
	}
}

func (d *AppsDaoImpl) GetAppsList(Type int) (*[]models.Apps, error) {
	var apps []models.Apps
	err := d.db.Where("type = ?", Type).Find(&apps).Error
	if err != nil {
		return nil, err
	}
	return &apps, nil
}

func (d *AppsDaoImpl) GetAppDetials(id string) (*models.Apps, error) {
	var app models.Apps
	err := d.db.First(&app, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &app, nil
}
