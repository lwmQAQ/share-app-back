package ecode

const (
	ErrNotFound      = 404
	ErrUserNotFound  = 100
	ErrPassWordError = 101
	ErrSystemError   = 102
	ErrCodeExtError  = 103
	ErrCodeError     = 104
	ErrUserNotExist  = 105
)

// 错误信息映射
var errorMessages = map[int]string{
	ErrNotFound:      "页面不存在",
	ErrUserNotFound:  "用户不存在",
	ErrPassWordError: "密码错误",
	ErrSystemError:   "系统错误",
	ErrCodeExtError:  "验证码过期",
	ErrCodeError:     "验证码错误",
	ErrUserNotExist:  "用户未登录",
}

func GetErrorMsg(code int) string {

	return errorMessages[code]
}
