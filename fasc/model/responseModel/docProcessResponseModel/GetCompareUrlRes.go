package docProcessResponseModel

type GetCompareUrlRes struct {
	RequestId string        `json:"requestId"`
	Code      string        `json:"code"`
	Msg       string        `json:"msg"`
	Data      GetCompareUrl `json:"data"`
}

type GetCompareUrl struct {
	CompareId  string `json:"compareId,omitempty"`
	CompareUrl string `json:"compareUrl,omitempty"`
}
