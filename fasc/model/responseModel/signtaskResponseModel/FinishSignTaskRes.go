package signtaskResponseModel

type FinishSignTaskRes struct {
	RequestId string `json:"requestId"`
	Code      string `json:"code"`
	Msg       string `json:"msg"`
}
