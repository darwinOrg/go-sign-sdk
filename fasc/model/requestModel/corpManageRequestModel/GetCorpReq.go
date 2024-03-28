package corpManageRequestModel

// GetCorpReq 查询企业用户基本信息
type GetCorpReq struct {
	OpenCorpId   string `json:"openCorpId,omitempty"`
	ClientCorpId string `json:"clientCorpId,omitempty"`
}
