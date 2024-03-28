package docProcessResponseModel

// UploadFileByUrlRes 通过网络文件地址上传
type UploadFileByUrlRes struct {
	RequestId string                 `json:"requestId"`
	Code      string                 `json:"code"`
	Msg       string                 `json:"msg"`
	Data      UploadFileByUrlResData `json:"data"`
}

type UploadFileByUrlResData struct {
	FddFileUrl string `json:"fddFileUrl,omitempty"`
}
