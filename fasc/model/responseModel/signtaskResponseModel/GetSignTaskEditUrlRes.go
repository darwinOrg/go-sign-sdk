package signtaskResponseModel

type GetSignTaskEditUrlRes struct {
	RequestId string                 `json:"requestId"`
	Code      string                 `json:"code"`
	Msg       string                 `json:"msg"`
	Data      GetSignTaskEditUrlData `json:"data"`
}

type GetSignTaskEditUrlData struct {
	SignTaskEditUrl string `json:"signTaskEditUrl,omitempty"`
}
