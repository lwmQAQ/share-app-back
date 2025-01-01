package adapter

import "tools-back/internal/models"

func BuildInsertTask(DownloadUrl string, userID uint64, FileName string) *models.Task {
	return &models.Task{
		SourceURL: DownloadUrl,
		Status:    0,
		UserID:    uint(userID),
		FileName:  FileName,
	}
}

func BuildUpdateTask(task *models.Task) *map[string]interface{} {
	updates := map[string]interface{}{}
	if task.DualURL != "" {
		updates["dual_url"] = task.DualURL
	}
	if task.MonoURL != "" {
		updates["mono_url"] = task.MonoURL
	}
	if task.Status != 0 {
		updates["status"] = task.Status
	}
	return &updates
}
