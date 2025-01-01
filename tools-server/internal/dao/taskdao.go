package dao

import "tools-back/internal/models"

type TaskDao interface {
	InsertTask(*models.Task) (int, error)
	UpdateTask(*map[string]interface{}, uint) error
}
