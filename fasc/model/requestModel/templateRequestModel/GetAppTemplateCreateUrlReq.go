package templateRequestModel

type GetAppTemplateCreateUrlReq struct {
	Type           string `json:"type,omitempty"`
	CreateSerialNo string `json:"createSerialNo,omitempty"`
	RedirectUrl    string `json:"redirectUrl,omitempty"`
}
