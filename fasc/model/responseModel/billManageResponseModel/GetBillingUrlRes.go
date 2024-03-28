package billManageResponseModel

type GetBillingUrlRes struct {
	RequestId string               `json:"requestId"`
	Code      string               `json:"code"`
	Msg       string               `json:"msg"`
	Data      GetBillingUrlResData `json:"data"`
}

type GetBillingUrlResData struct {
	EUrl string `json:"eUrl"`
}
