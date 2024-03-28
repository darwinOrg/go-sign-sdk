package sealRequestModel

// GetSealUserListReq 获取企业用印员列表
type GetSealUserListReq struct {
	OwnerId    string          `json:"ownerId,omitempty"`
	OpenCorpId string          `json:"openCorpId,omitempty"`
	ListFilter *SealListFilter `json:"listFilter,omitempty"`
}
