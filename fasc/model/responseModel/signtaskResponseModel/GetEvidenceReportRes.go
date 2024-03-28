package signtaskResponseModel

type GetEvidenceReportRes struct {
	RequestId string                `json:"requestId"`
	Code      string                `json:"code"`
	Msg       string                `json:"msg"`
	Data      GetEvidenceReportData `json:"data"`
}

type GetEvidenceReportData struct {
	DownloadUrl string `json:"downloadUrl,omitempty"`
}
