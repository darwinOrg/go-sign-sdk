package signtaskResponseModel

type AbolishSignTaskRes struct {
	RequestId string              `json:"requestId"`
	Code      string              `json:"code"`
	Msg       string              `json:"msg"`
	Data      AbolishSignTaskData `json:"data"`
}

type AbolishSignTaskData struct {
	AbolishedSignTaskId string `json:"abolishedSignTaskId,omitempty"`
}
