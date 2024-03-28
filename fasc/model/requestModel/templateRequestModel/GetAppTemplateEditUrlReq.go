package templateRequestModel

type GetAppTemplateEditUrlReq struct {
	AppTemplateId string `json:"appTemplateId,omitempty"`
	RedirectUrl   string `json:"redirectUrl,omitempty"`
}
