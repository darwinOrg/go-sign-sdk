package sealRequestModel

type SealDeleteReq struct {
	OpenCorpId string `json:"openCorpId,omitempty"`
	SealId     int64  `json:"sealId,string"`
}
