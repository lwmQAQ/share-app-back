package handle

import (
	"encoding/json"
	"net/http"
	"time"
	"tools-back/internal/types"

	"github.com/gin-gonic/gin"
)

// SSE 处理函数
func SSEHandler(c *gin.Context, msgchan *chan *types.TranslationTaskResp) {
	// 设置响应头
	c.Header("Content-Type", "text/event-stream")
	c.Header("Cache-Control", "no-cache")
	c.Header("Connection", "keep-alive")

	// 每隔1秒发送一次消息
	for {
		select {
		// 如果通道有消息，则从通道中读取消息并发送给客户端
		case msg := <-*msgchan:
			// 将 msg 转换为 JSON 或其他格式（假设已实现）
			// 假设 TranslationTaskResp 是可以序列化为 JSON 的结构体
			taskJSON, err := json.Marshal(msg)
			if err != nil {
				c.String(http.StatusInternalServerError, "序列化结构体失败")
				return
			}
			// 发送消息到前端
			c.SSEvent("message", string(taskJSON))
			c.Writer.Flush()

		// 如果通道没有消息，则继续发送定时器事件（保持连接活跃）
		case <-time.After(2 * time.Second):
			c.SSEvent("heartbeat", time.Now().Format(time.RFC3339))
			c.Writer.Flush()
		}
	}
}
