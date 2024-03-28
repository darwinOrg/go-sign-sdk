package userManageResponseModel

type GetIdCardImageDownloadUrlRes struct {
	RequestId string                        `json:"requestId"`
	Code      string                        `json:"code"`
	Msg       string                        `json:"msg"`
	Data      GetIdCardImageDownloadUrlData `json:"data"`
}

type GetIdCardImageDownloadUrlData struct {
	DownloadUrl string `json:"downloadUrl,omitempty"`
}
