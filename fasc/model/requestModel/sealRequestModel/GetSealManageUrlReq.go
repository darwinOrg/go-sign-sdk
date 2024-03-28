package sealRequestModel

type GetSealManageUrlReq struct {
	OpenCorpId   string `json:"openCorpId,omitempty"`
	ClientUserId string `json:"clientUserId,omitempty"`
	RedirectUrl  string `json:"redirectUrl,omitempty"`
}
