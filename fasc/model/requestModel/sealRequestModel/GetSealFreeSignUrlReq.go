package sealRequestModel

type GetSealFreeSignUrlReq struct {
	OpenCorpId   string  `json:"openCorpId,omitempty"`
	SealIds      []int64 `json:"sealIds,omitempty"`
	BusinessId   string  `json:"businessId,omitempty"`
	Email        string  `json:"email,omitempty"`
	ClientUserId string  `json:"clientUserId,omitempty"`
	RedirectUrl  string  `json:"redirectUrl,omitempty"`
}
