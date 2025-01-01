package dao

import (
	"errors"
	"fmt"
	"log"
	"tools-back/internal/models"

	"gorm.io/gorm"
)

type TaskDaoImpl struct {
	db *gorm.DB
}

func NewTaskDao(db *gorm.DB) TaskDao {
	return &TaskDaoImpl{
		db: db,
	}
}
func (d *TaskDaoImpl) InsertTask(insert *models.Task) (int, error) {
	// 执行插入操作
	result := d.db.Create(insert)

	// 如果插入失败，返回错误
	if result.Error != nil {
		log.Println("Insert task failed:", result.Error)
		return 0, result.Error
	}

	// 返回插入的 ID 和 nil 错误
	return int(insert.ID), nil
}

func (d *TaskDaoImpl) UpdateTask(updates *map[string]interface{}, id uint) error {
	// 检查是否有需要更新的字段
	if len(*updates) == 0 {
		return errors.New("no fields to update")
	}

	// 执行更新操作，只更新不为空的字段
	result := d.db.Model(&models.Task{}).Where("taskid = ?", id).Updates(updates)
	// 检查是否有错误
	if result.Error != nil {
		return fmt.Errorf("更新失败：%s", result.Error)
	}
	return nil
}
