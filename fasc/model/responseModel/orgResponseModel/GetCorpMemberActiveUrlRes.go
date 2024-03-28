package orgResponseModel

type GetCorpMemberActiveUrlRes struct {
	RequestId string               `json:"requestId"`
	Code      string               `json:"code"`
	Msg       string               `json:"msg"`
	Data      []SimpleEmployeeInfo `json:"data"`
}
