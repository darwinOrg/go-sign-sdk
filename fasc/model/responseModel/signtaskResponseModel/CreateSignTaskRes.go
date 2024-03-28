package signtaskResponseModel

type CreateSignTaskRes struct {
	RequestId string                `json:"requestId"`
	Code      string                `json:"code"`
	Msg       string                `json:"msg"`
	Data      CreateSignTaskResData `json:"data"`
}
type CreateSignTaskResData struct {
	SignTaskId string `json:"signTaskId,omitempty"`
}
