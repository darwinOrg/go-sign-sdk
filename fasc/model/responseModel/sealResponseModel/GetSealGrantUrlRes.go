package sealResponseModel

type GetSealGrantUrlRes struct {
	RequestId string          `json:"requestId"`
	Code      string          `json:"code"`
	Msg       string          `json:"msg"`
	Data      GetSealGrantUrl `json:"data"`
}

type GetSealGrantUrl struct {
	SealGrantUrl string `json:"sealGrantUrl,omitempty"`
}
