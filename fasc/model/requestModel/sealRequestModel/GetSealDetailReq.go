package sealRequestModel

type GetSealDetailReq struct {
	OpenCorpId string `json:"openCorpId,omitempty"`
	SealId     int64  `json:"sealId,string"`
}
