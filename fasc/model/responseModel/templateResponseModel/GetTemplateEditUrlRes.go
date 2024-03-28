package templateResponseModel

type GetTemplateEditUrlRes struct {
	RequestId string             `json:"requestId,omitempty"`
	Code      string             `json:"code,omitempty"`
	Msg       string             `json:"msg,omitempty"`
	Data      GetTemplateEditUrl `json:"data,omitempty"`
}

type GetTemplateEditUrl struct {
	TemplateEditUrl string `json:"templateEditUrl,omitempty"`
}
