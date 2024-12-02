package ecode

const (
	ErrNotFound = 404
)

// 错误信息映射
var errorMessages = map[int]string{
	ErrNotFound: "页面不存在",
}

func GetErrorMsg(code int) string {
	return errorMessages[code]
}
