package adapter

import (
	"resource-server/internal/models"
	"resource-server/internal/types"
	"time"
)

func BuildInsertComment(req types.CreateCommentReq, user *models.User, level int, path string) *models.Comment {

	return &models.Comment{
		PostID:    req.PostID,
		Content:   req.Content,
		Likes:     0,
		ParentID:  req.ParentID,
		CreatedAt: time.Now(),
		Level:     level,
		Path:      path,
	}
}
