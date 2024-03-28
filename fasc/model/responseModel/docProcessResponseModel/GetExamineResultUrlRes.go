package docProcessResponseModel

type GetExamineResultUrlRes struct {
	RequestId string              `json:"requestId"`
	Code      string              `json:"code"`
	Msg       string              `json:"msg"`
	Data      GetExamineResultUrl `json:"data"`
}

type GetExamineResultUrl struct {
	ExamineUrl string `json:"examineUrl,omitempty"`
}
