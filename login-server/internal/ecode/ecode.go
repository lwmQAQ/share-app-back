package ecode

const (
	ErrNotFound      = 404
	ErrUserNotFound  = 100
	ErrPassWordError = 101
	ErrSystemError   = 102
)

// 错误信息映射
var errorMessages = map[int]string{
	ErrNotFound:      "页面不存在",
	ErrUserNotFound:  "用户不存在",
	ErrPassWordError: "密码错误",
	ErrSystemError:   "系统错误",
}

func GetErrorMsg(code int) string {

	return errorMessages[code]
}
