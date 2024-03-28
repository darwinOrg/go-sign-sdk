package docProcessResponseModel

type GetExamineUrlRes struct {
	RequestId string        `json:"requestId"`
	Code      string        `json:"code"`
	Msg       string        `json:"msg"`
	Data      GetExamineUrl `json:"data"`
}

type GetExamineUrl struct {
	ExamineId  string `json:"examineId,omitempty"`
	ExamineUrl string `json:"examineUrl,omitempty"`
}
