package sealRequestModel

type GetPersonalSealFreeSignUrlReq struct {
	OpenUserId  string  `json:"openUserId,omitempty"`
	SealIds     []int64 `json:"sealIds,omitempty"`
	BusinessId  string  `json:"businessId,omitempty"`
	Email       string  `json:"email,omitempty"`
	RedirectUrl string  `json:"redirectUrl,omitempty"`
}
