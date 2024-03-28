package sealRequestModel

type GetUserSealListReq struct {
	OpenCorpId string `json:"openCorpId,omitempty"`
	MemberId   int64  `json:"memberId,string"`
}
