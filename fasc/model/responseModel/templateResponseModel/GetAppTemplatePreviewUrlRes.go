package templateResponseModel

type GetAppTemplatePreviewUrlRes struct {
	RequestId string                       `json:"requestId,omitempty"`
	Code      string                       `json:"code,omitempty"`
	Msg       string                       `json:"msg,omitempty"`
	Data      GetAppTemplatePreviewUrlData `json:"data,omitempty"`
}

type GetAppTemplatePreviewUrlData struct {
	AppTemplatePreviewUrl string `json:"appTemplatePreviewUrl,omitempty"`
}
