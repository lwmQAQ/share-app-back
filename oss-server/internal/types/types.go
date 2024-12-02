package types

import "oss-server/internal/ecode"

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

type OSSUploadReq struct {
	FileName string `json:"filename"`
	Scene    int32  `json:"scene"`
}

type OSSUploadResp struct {
	UploadUrl   string `json:"uploadurl"`
	DownloadUrl string `json:"downloadurl"`
}
