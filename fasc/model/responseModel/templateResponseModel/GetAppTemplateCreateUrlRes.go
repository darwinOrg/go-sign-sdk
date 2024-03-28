package templateResponseModel

type GetAppTemplateCreateUrlRes struct {
	RequestId string                      `json:"requestId,omitempty"`
	Code      string                      `json:"code,omitempty"`
	Msg       string                      `json:"msg,omitempty"`
	Data      GetAppTemplateCreateUrlData `json:"data,omitempty"`
}

type GetAppTemplateCreateUrlData struct {
	AppTemplateCreateUrl string `json:"appTemplateCreateUrl,omitempty"`
}
