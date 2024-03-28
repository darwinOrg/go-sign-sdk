package templateResponseModel

type GetTemplateCreateUrlRes struct {
	RequestId string               `json:"requestId,omitempty"`
	Code      string               `json:"code,omitempty"`
	Msg       string               `json:"msg,omitempty"`
	Data      GetTemplateCreateUrl `json:"data,omitempty"`
}

type GetTemplateCreateUrl struct {
	TemplateCreateUrl string `json:"templateCreateUrl,omitempty"`
}
