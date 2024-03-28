package signtaskResponseModel

type BatchSignUrlRes struct {
	RequestId string           `json:"requestId"`
	Code      string           `json:"code"`
	Msg       string           `json:"msg"`
	Data      BatchSignUrlData `json:"data"`
}
type BatchSignUrlData struct {
	BatchSignUrl string `json:"batchSignUrl,omitempty"`
}
