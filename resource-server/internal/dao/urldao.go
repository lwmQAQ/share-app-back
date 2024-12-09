package dao

import "resource-server/internal/models"

type UrlDao interface {
	CreateUrl(url *models.Url) error
	GetUrlByCode(email string) (*models.Url, error)
	DeleteUrl(code string) error
}
