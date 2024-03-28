package docProcessResponseModel

type GetCompareResultUrlRes struct {
	RequestId string              `json:"requestId"`
	Code      string              `json:"code"`
	Msg       string              `json:"msg"`
	Data      GetCompareResultUrl `json:"data"`
}

type GetCompareResultUrl struct {
	CompareUrl string `json:"compareUrl,omitempty"`
}
