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
	Type    int      `json:"type"`
	Content string   `json:"content"`
	Link    string   `json:"link"`
}

type CreatePostResp struct {
}

// 创建评论
type CreateCommentReq struct {
	UserID   uint64  `json:"userId"`
	PostID   uint64  `json:"postId" validate:"required"`                // 所属帖子 ID，必填
	ParentID *uint64 `json:"parentId,omitempty"`                        // 父评论 ID，可选
	Content  string  `json:"content" validate:"required,min=1,max=500"` // 评论内容，必填，限制长度
}

type CreateCommentResp struct {
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
