package signtaskResponseModel

type GetSignTaskPicDownloadUrlRes struct {
	RequestId string                        `json:"requestId"`
	Code      string                        `json:"code"`
	Msg       string                        `json:"msg"`
	Data      GetSignTaskPicDownloadUrlData `json:"data"`
}

type GetSignTaskPicDownloadUrlData struct {
	DownloadUrl string `json:"downloadUrl,omitempty"`
}
