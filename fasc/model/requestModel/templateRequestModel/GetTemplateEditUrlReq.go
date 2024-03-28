package templateRequestModel

type GetTemplateEditUrlReq struct {
	OpenCorpId  string `json:"openCorpId,omitempty"`
	TemplateId  string `json:"templateId,omitempty"`
	RedirectUrl string `json:"redirectUrl,omitempty"`
}
