package templateRequestModel

type GetTemplatePreviewUrlReq struct {
	OpenCorpId  string `json:"openCorpId,omitempty"`
	TemplateId  string `json:"templateId,omitempty"`
	RedirectUrl string `json:"redirectUrl,omitempty"`
}
