package types

import "login-server/internal/ecode"

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

type GetUserInfoResp struct {
	ID         uint64 `json:"id" validate:"required"`
	Name       string `json:"name"`
	Avatar     string `json:"avatar"`
	Sex        int    `json:"sex,options=1|2"`
	Bio        string `json:"bio"`
	Experience int    `json:"experience"`
	Level      int    `json:"level"`
}

type LoginCodeReq struct {
	Email string `json:"email" validate:"required,email"`
	Code  string `json:"code" validate:"required"`
}

type LoginReq struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type LoginResp struct {
	ID    uint64 `json:"id"`
	Token string `json:"token"`
}

type RegisterReq struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
	Code     string `json:"code" validate:"required,len=6,regexp=^[0-9]{6}$"`
}

type RegisterResp struct {
	ID    uint64 `json:"id"`
	Token string `json:"token"`
}

type SendCodeReq struct {
	Email string `form:"email" validate:"required,email"` // 使用query参数
}

type UpdateUserPasswordReq struct {
	ID       uint64 `json:"id" validate:"required"`
	Email    string `json:"email,optional" validate:"email"`
	Password string `json:"password,optional"`
	Code     string `json:"code" validate:"required,len=6,regexp=^[0-9]{6}$"`
}

type UpdateUserReq struct {
	ID     uint64 `json:"id" validate:"required"`
	Name   string `json:"name,optional"`
	Avatar string `json:"avatar,optional"`
	Sex    int    `json:"sex,options=1|2"`
	Bio    string `json:"bio,optional"`
}

type UpdateUserResp struct {
}

type OSSUploadReq struct {
	FileName string `json:"filename"`
	Scene    int32  `json:"scene"`
}
type OSSUploadResp struct {
	UploadUrl   string `json:"uploadurl"`
	DownloadUrl string `json:"downloadurl"`
}

type IPResp struct {
	IP string `json:"ip"`
}
