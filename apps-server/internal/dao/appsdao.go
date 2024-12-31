package dao

import "apps-server/internal/models"

type AppsDao interface {
	GetAppsList(int) (*[]models.Apps, error)
	GetAppDetials(string) (*models.Apps, error)
}
