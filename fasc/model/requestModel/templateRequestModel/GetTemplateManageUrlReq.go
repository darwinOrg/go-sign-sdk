package templateRequestModel

type GetTemplateManageUrlReq struct {
	OpenCorpId  string `json:"openCorpId,omitempty"`
	RedirectUrl string `json:"redirectUrl,omitempty"`
}
