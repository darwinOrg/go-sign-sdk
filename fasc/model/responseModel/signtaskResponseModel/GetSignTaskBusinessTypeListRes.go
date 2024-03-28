package signtaskResponseModel

type GetSignTaskBusinessTypeListRes struct {
	RequestId string                        `json:"requestId"`
	Code      string                        `json:"code"`
	Msg       string                        `json:"msg"`
	Data      []GetSignTaskBusinessTypeData `json:"data"`
}

type GetSignTaskBusinessTypeData struct {
	BusinessTypeId   int64  `json:"businessTypeId,omitempty"`
	BusinessTypeName string `json:"businessTypeName,omitempty"`
}
