package sealResponseModel

type GetAppointedSealUrlRes struct {
	RequestId string             `json:"requestId"`
	Code      string             `json:"code"`
	Msg       string             `json:"msg"`
	Data      GetAppointedSealUr `json:"data"`
}

type GetAppointedSealUr struct {
	AppointedSealUrl string `json:"appointedSealUrl,omitempty"`
}
