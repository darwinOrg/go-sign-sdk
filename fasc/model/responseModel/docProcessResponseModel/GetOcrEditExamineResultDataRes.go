package docProcessResponseModel

type GetOcrEditExamineResultDataRes struct {
	RequestId string                      `json:"requestId"`
	Code      string                      `json:"code"`
	Msg       string                      `json:"msg"`
	Data      GetOcrEditExamineResultData `json:"data"`
}

type GetOcrEditExamineResultData struct {
	ExamineData string `json:"examineData,omitempty"`
}
