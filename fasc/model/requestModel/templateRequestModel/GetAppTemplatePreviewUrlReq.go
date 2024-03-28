package templateRequestModel

type GetAppTemplatePreviewUrlReq struct {
	AppTemplateId string `json:"appTemplateId,omitempty"`
	RedirectUrl   string `json:"redirectUrl,omitempty"`
}
