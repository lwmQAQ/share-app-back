package adapter

import (
	"resource-server/internal/models"
	"resource-server/internal/types"
	"time"
)

func BuildInsertPost(author *models.User, req *types.CreatePostReq) *models.Post {
	t := time.Now()
	return &models.Post{
		Title:       req.Title,
		Content:     req.Title,
		Author:      *author,
		Tags:        req.Tags,
		LikesCount:  0,
		CommentsIDs: nil,
		CreatedAt:   t,
		UpdatedAt:   t,
	}
}
