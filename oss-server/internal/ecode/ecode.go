package ecode

// 定义错误码常量
const ()

// 错误信息映射
var errorMessages = map[int]string{}

func GetErrorMsg(code int) string {
	return errorMessages[code]
}
