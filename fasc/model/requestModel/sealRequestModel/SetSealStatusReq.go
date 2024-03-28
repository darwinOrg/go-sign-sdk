package sealRequestModel

type SetSealStatusReq struct {
	OpenCorpId string `json:"openCorpId,omitempty"`
	SealId     int64  `json:"sealId,string"`
	SealStatus string `json:"sealStatus,omitempty"`
}
