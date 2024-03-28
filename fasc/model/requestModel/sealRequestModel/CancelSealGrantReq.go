package sealRequestModel

type CancelSealGrantReq struct {
	OpenCorpId string `json:"openCorpId,omitempty"`
	SealId     int64  `json:"sealId,string"`
	MemberId   int64  `json:"memberId,string"`
}
