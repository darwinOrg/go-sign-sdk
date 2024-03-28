package signtaskResponseModel

type SignTaskFieldUrlData struct {
	DownloadUrl string `json:"downloadUrl"`
}

type GetSignTaskFieldUrlRes struct {
	RequestId string               `json:"requestId"`
	Code      string               `json:"code"`
	Msg       string               `json:"msg"`
	Data      SignTaskFieldUrlData `json:"data"`
}
