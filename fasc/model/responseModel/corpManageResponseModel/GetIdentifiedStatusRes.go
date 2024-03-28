package corpManageResponseModel

type GetIdentifiedStatusRes struct {
	RequestId string                     `json:"requestId"`
	Code      string                     `json:"code"`
	Msg       string                     `json:"msg"`
	Data      GetIdentifiedStatusResData `json:"data"`
}

type GetIdentifiedStatusResData struct {
	IdentStatus bool `json:"identStatus"`
}
