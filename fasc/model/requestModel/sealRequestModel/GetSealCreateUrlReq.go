package sealRequestModel

type GetSealCreateUrlReq struct {
	OpenCorpId     string `json:"openCorpId,omitempty"`
	SealName       string `json:"sealName,omitempty"`
	CategoryType   string `json:"categoryType,omitempty"`
	SealTag        string `json:"sealTag,omitempty"`
	CreateSerialNo string `json:"createSerialNo,omitempty"`
	ClientUserId   string `json:"clientUserId,omitempty"`
	RedirectUrl    string `json:"redirectUrl,omitempty"`
}
