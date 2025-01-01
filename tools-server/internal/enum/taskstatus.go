package enum

type TaskStatus int

const (
	Prepping TaskStatus = iota
	Complete
	Error
)

var taskStatusMessages = map[TaskStatus]string{
	Prepping: "准备中",
	Complete: "完成",
	Error:    "错误",
}

func GetTaskStatusMessage(status TaskStatus) string {
	if msg, exists := taskStatusMessages[status]; exists {
		return msg
	}
	return "未知状态"
}
