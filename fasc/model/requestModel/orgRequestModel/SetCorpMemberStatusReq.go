package orgRequestModel

type SetCorpMemberStatusReq struct {
	OpenCorpId   string  `json:"openCorpId,omitempty"`
	MemberIds    []int64 `json:"memberIds,omitempty"`
	MemberStatus string  `json:"memberStatus,omitempty"`
}
