package sealResponseModel

type GetSealFreeSignUrlRes struct {
	RequestId string             `json:"requestId"`
	Code      string             `json:"code"`
	Msg       string             `json:"msg"`
	Data      GetSealFreeSignUrl `json:"data"`
}

type GetSealFreeSignUrl struct {
	FreeSignUrl string `json:"freeSignUrl,omitempty"`
}
