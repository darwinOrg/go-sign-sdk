package sealRequestModel

// GetSealListReq 获取印章列表
type GetSealListReq struct {
	OwnerId    string          `json:"ownerId,omitempty"`
	OpenCorpId string          `json:"openCorpId,omitempty"`
	ListFilter *SealListFilter `json:"listFilter,omitempty"`
}

type SealListFilter struct {
	CategoryType []string `json:"categoryType,omitempty"`
}
