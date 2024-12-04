package types

import (
	"resource-server/internal/ecode"
	"resource-server/internal/models"
	"time"
)

type Response struct {
	Code         int32       `json:"code"`
	ErrorMessage string      `json:"errorMessage"`
	Data         interface{} `json:"data"`
}

func Success(data interface{}) Response {
	return Response{
		Code:         200,
		ErrorMessage: "",
		Data:         data,
	}
}
func Danger(data interface{}) Response {
	return Response{
		Code:         444,
		ErrorMessage: "",
		Data:         data,
	}
}
func Error(code int) Response {
	return Response{
		Code:         500,
		ErrorMessage: ecode.GetErrorMsg(code),
		Data:         nil,
	}
}
func ErrorMsg(msg string) Response {
	return Response{
		Code:         500,
		ErrorMessage: msg,
		Data:         nil,
	}
}

type CreatePostReq struct {
	UserID  uint64   `json:"userId"`
	Title   string   `json:"title"`
	Tags    []string `json:"tags"`
	Content string   `json:"content"`
}

type CreatePostResp struct {
}

type GetPostByIdReq struct {
	PostID string `json:"postId"`
}

type GetPostByIdResp struct {
	PostID     string      `json:"postId"`                       //资源贴id
	Title      string      `json:"title"`                        // 帖子标题
	Content    string      `json:"content"`                      // 富文本内容，HTML 格式
	Author     models.User `json:"author"`                       // 帖子的作者信息
	LikesCount int64       `json:"likes_count"`                  // 点赞数
	CommentNum int64       `json:"CommentNum"`                   //评论数
	Tags       []string    `json:"tags,omitempty"`               // 帖子标签列表
	UpdatedAt  time.Time   `bson:"updated_at" json:"updated_at"` // 更新时间
}
