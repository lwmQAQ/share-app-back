package constant

import (
	"fmt"
	"time"
)

const BASEKEY = "resource:"
const USER_TOKEN_STRING = "userToken:uid_"
const USER_INFO_STRING = "userInfo:uid_"
const CODE_STRING = "code:email_"
const TASK_STRING = "task:uuid_"
const FUNC_STRING = "func:uid_"
const Theme_STRING = "theme:type_"
const Limit_STRING = "limit:ip_"
const USER_TOKEN_EX = 5 * 24 * time.Hour
const USER_INFO_EX = 15 * time.Minute
const FUNC_EX = 15 * time.Minute

func BuildCodeKey(email string) string {
	return fmt.Sprintf("%s%s%s", BASEKEY, CODE_STRING, email)
}
func BuildTaskKey(id string) string {
	return fmt.Sprintf("%s%s%s", BASEKEY, TASK_STRING, id)
}
func BuildTokenKey(id uint64) string {
	return fmt.Sprintf("%s%s%d", BASEKEY, USER_TOKEN_STRING, id)
}

func BuildInfoKey(id uint64) string {
	return fmt.Sprintf("%s%s%d", BASEKEY, USER_INFO_STRING, id)
}
func BuildThemeKey(Type int64) string {
	return fmt.Sprintf("%s%s%d", BASEKEY, Theme_STRING, Type)
}
func BuildLimitKey(ip string) string {
	return fmt.Sprintf("%s%s%s", BASEKEY, TASK_STRING, ip)
}
func BuildFuncKey(id uint64) string {
	return fmt.Sprintf("%s%s%d", BASEKEY, FUNC_STRING, id)
}
