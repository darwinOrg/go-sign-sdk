package docProcessResponseModel

// GetLocalUploadFileUrlRes 获取本地上传文件地址
type GetLocalUploadFileUrlRes struct {
	RequestId string                       `json:"requestId"`
	Code      string                       `json:"code"`
	Msg       string                       `json:"msg"`
	Data      GetLocalUploadFileUrlResData `json:"data"`
}

type GetLocalUploadFileUrlResData struct {
	UploadUrl  string `json:"uploadUrl"`
	FddFileUrl string `json:"fddFileUrl"`
}
