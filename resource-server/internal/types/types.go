package types

import "resource-server/internal/ecode"

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
	UserID  uint64
	Title   string
	Tags    []string
	Content string
}

type CreatePostResp struct {
}
