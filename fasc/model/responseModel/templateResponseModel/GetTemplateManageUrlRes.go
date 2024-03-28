package templateResponseModel

type GetTemplateManageUrlRes struct {
	RequestId string               `json:"requestId,omitempty"`
	Code      string               `json:"code,omitempty"`
	Msg       string               `json:"msg,omitempty"`
	Data      GetTemplateManageUrl `json:"data,omitempty"`
}

type GetTemplateManageUrl struct {
	TemplateManageUrl string `json:"templateManageUrl,omitempty"`
}
