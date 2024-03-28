package templateResponseModel

type GetTemplatePreviewUrlRes struct {
	RequestId string                `json:"requestId,omitempty"`
	Code      string                `json:"code,omitempty"`
	Msg       string                `json:"msg,omitempty"`
	Data      GetTemplatePreviewUrl `json:"data,omitempty"`
}

type GetTemplatePreviewUrl struct {
	TemplatePreviewUrl string `json:"templatePreviewUrl,omitempty"`
}
