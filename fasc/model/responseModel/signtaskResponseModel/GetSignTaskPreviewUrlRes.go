package signtaskResponseModel

type GetSignTaskPreviewUrlRes struct {
	RequestId string                    `json:"requestId"`
	Code      string                    `json:"code"`
	Msg       string                    `json:"msg"`
	Data      GetSignTaskPreviewUrlData `json:"data"`
}

type GetSignTaskPreviewUrlData struct {
	SignTaskPreviewUrl string `json:"signTaskPreviewUrl,omitempty"`
}
