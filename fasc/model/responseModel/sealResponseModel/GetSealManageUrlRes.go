package sealResponseModel

type GetSealManageUrlRes struct {
	RequestId string           `json:"requestId"`
	Code      string           `json:"code"`
	Msg       string           `json:"msg"`
	Data      GetSealManageUrl `json:"data"`
}

type GetSealManageUrl struct {
	ResourceUrl string `json:"resourceUrl,omitempty"`
}
