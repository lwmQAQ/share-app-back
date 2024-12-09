package ecode

const (
	ErrNotFound = 404

	ErrSystemError     = 102
	ErrCodeExtError    = 103
	ErrCodeError       = 104
	ErrUserNotExist    = 105
	ErrIllegalRequests = 106
)

// 错误信息映射
var errorMessages = map[int]string{
	ErrNotFound:        "页面不存在",
	ErrSystemError:     "系统错误",
	ErrCodeExtError:    "验证码过期",
	ErrCodeError:       "验证码错误",
	ErrUserNotExist:    "用户未登录",
	ErrIllegalRequests: "非法请求",
}

func GetErrorMsg(code int) string {

	return errorMessages[code]
}
