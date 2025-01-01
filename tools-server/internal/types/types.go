package types

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

func ErrorMsg(msg string) Response {
	return Response{
		Code:         500,
		ErrorMessage: msg,
		Data:         nil,
	}
}

type OssUploadReq struct {
	Postfix string `form:"postfix"` // 使用 form 标签
}

type OssUploadResp struct {
	UploadUrl   string `json:"upload_url"`
	DownLoadUrl string `json:"download_url"`
}

type TranslationReq struct {
	FileName    string `json:"file_name"`
	DownLoadUrl string `json:"download_url"`
}

type TranslationResp struct {
	TaskID int `json:"task_id"`
}
type RpcChanMessage struct {
	DownLoadUrl string
	TaskID      int
	FileName    string
}

type TranslationTaskResp struct {
	Status    int    `json:"status"`
	TaskID    int    `json:"task_id"`
	FileName  string `json:"file_name"`
	SourceURL string `json:"source_url"`
	MonoURL   string `json:"mono_url"`
	DualURL   string `json:"dual_url"`
}
