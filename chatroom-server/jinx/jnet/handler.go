package jnet

import (
	"chatroom-server/jinx/jiface"
	"chatroom-server/jinx/utils"
	"fmt"
	"strconv"
	"sync"
)

// handler只能同时处理 WokerPollSize 个请求,别的请求需要在消息队列中排队
type Handler struct {
	Apis          map[uint32]jiface.IRouter
	TaskQueue     []chan jiface.IRequest //任务队列
	WokerPollSize uint32                 //工作池woker大小
	NextQueue     int                    //轮询方式
}

func NewHandler() jiface.IHandler {
	return &Handler{
		Apis:          make(map[uint32]jiface.IRouter),
		TaskQueue:     make([]chan jiface.IRequest, utils.MyApplication.Server.WokerPollSize),
		WokerPollSize: utils.MyApplication.Server.WokerPollSize,
		NextQueue:     0,
	}
}
func (h *Handler) BindRouter(msgId uint32, router jiface.IRouter) {
	if _, ok := h.Apis[msgId]; ok {
		panic("路由已被注册" + strconv.Itoa(int(msgId)))
	}
	h.Apis[msgId] = router
}

func (h *Handler) UseRouter(req jiface.IRequest) {
	if h.Apis == nil {
		fmt.Println("未绑定路由")
		return
	}
	msgID := req.GetMsgID()
	Func := h.Apis[msgID]
	Func.PreHandle(req)
	Func.Handle(req)
	Func.PostHandle(req)
}

// 启动工作池
func (h *Handler) StartWorkerPool() {
	fmt.Println("开辟连接池")
	for i := 0; i < int(h.WokerPollSize); i++ {
		//开辟空间
		h.TaskQueue[i] = make(chan jiface.IRequest, utils.MyApplication.Server.MaxTaskQueueNum)
		//启动队列
		go h.startWorker(i, h.TaskQueue[i])
	}
}

// 释放工作池
func (h *Handler) StopWorkerPool() {
	fmt.Println("释放连接池")
	for i, taskChan := range h.TaskQueue {
		if taskChan != nil {
			close(taskChan) // 关闭通道，通知工作协程停止运行
			fmt.Printf("关闭消息队列 %d\n", i)
		}
	}
}

func (h *Handler) startWorker(workerID int, taskQueue chan jiface.IRequest) {
	fmt.Printf("消息队列 %d 启动\n", workerID)
	for {
		request, ok := <-taskQueue
		if !ok { // 检测通道是否关闭
			fmt.Printf("消息队列 %d 停止\n", workerID)
			return
		}
		if request == nil { // 防止意外读取到 nil 请求
			continue
		}
		h.UseRouter(request)
	}
}

var mu sync.Mutex

// 消息队列负载均衡
func (h *Handler) HandleRequest(req jiface.IRequest) {
	mu.Lock()
	defer mu.Unlock()
	h.TaskQueue[h.NextQueue] <- req
	//轮询
	h.NextQueue = (h.NextQueue + 1) % int(h.WokerPollSize)
}
