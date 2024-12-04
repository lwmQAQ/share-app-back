package commentserver

import (
	"resource-server/utils"
)

type CommentServer struct {
	mongoutil utils.MongoUtil
}

func NewCommentServer(mongoutil utils.MongoUtil) *CommentServer {
	return &CommentServer{
		mongoutil: mongoutil,
	}
}
