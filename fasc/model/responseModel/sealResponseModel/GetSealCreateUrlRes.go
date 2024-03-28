package sealResponseModel

type GetSealCreateUrlRes struct {
	RequestId string           `json:"requestId"`
	Code      string           `json:"code"`
	Msg       string           `json:"msg"`
	Data      GetSealCreateUrl `json:"data"`
}

type GetSealCreateUrl struct {
	SealCreateUrl string `json:"sealCreateUrl,omitempty"`
}
