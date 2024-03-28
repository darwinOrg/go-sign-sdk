package orgRequestModel

type GetCorpMemberActiveUrlReq struct {
	OpenCorpId string  `json:"openCorpId,omitempty"`
	MemberIds  []int64 `json:"memberIds,omitempty"`
}
