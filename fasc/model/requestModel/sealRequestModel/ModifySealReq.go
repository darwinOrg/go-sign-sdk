package sealRequestModel

type ModifySealReq struct {
	OpenCorpId   string `json:"openCorpId,omitempty"`
	SealId       int64  `json:"sealId,string"`
	SealName     string `json:"sealName,omitempty"`
	SealTag      string `json:"sealTag,omitempty"`
	CategoryType string `json:"categoryType,omitempty"`
}
