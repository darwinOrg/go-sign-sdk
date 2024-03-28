package orgRequestModel

type ModifyCorpMemberReq struct {
	OpenCorpId         string  `json:"openCorpId,omitempty"`
	MemberId           int64   `json:"memberId,omitempty"`
	MemberName         string  `json:"memberName,omitempty"`
	InternalIdentifier string  `json:"internalIdentifier,omitempty"`
	MemberEmail        string  `json:"memberEmail,omitempty"`
	MemberMobile       string  `json:"memberMobile,omitempty"`
	MemberDeptIds      []int64 `json:"memberDeptIds,omitempty"`
}
