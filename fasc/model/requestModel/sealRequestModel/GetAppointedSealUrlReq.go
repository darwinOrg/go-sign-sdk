package sealRequestModel

type GetAppointedSealUrlReq struct {
	OpenCorpId   string `json:"openCorpId,omitempty"`
	SealId       int64  `json:"sealId,string"`
	ClientUserId string `json:"ClientUserId,omitempty"`
	RedirectUrl  string `json:"redirectUrl,omitempty"`
}
