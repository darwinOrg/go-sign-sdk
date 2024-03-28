package templateResponseModel

type GetAppTemplateEditUrlRes struct {
	RequestId string                    `json:"requestId,omitempty"`
	Code      string                    `json:"code,omitempty"`
	Msg       string                    `json:"msg,omitempty"`
	Data      GetAppTemplateEditUrlData `json:"data,omitempty"`
}

type GetAppTemplateEditUrlData struct {
	AppTemplateEditUrl string `json:"appTemplateEditUrl,omitempty"`
}
