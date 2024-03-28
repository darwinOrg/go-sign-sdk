package templateRequestModel

type GetTemplateCreateUrlReq struct {
	OpenCorpId     string `json:"openCorpId,omitempty"`
	Type           string `json:"type,omitempty"`
	CreateSerialNo string `json:"createSerialNo,omitempty"`
	RedirectUrl    string `json:"redirectUrl,omitempty"`
}
