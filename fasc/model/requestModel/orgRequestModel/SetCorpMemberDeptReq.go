package orgRequestModel

type SetCorpMemberDeptReq struct {
	OpenCorpId    string  `json:"openCorpId,omitempty"`
	MemberIds     []int64 `json:"memberIds,omitempty"`
	MemberDeptIds []int64 `json:"memberDeptIds,omitempty"`
	Model         string  `json:"model,omitempty"`
}
