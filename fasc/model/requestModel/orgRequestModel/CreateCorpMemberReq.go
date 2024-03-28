package orgRequestModel

type CreateCorpMemberReq struct {
	OpenCorpId          string         `json:"openCorpId,omitempty"`
	EmployeeInfos       []EmployeeInfo `json:"employeeInfos,omitempty"`
	NotifyActiveByEmail bool           `json:"notifyActiveByEmail"`
}

type EmployeeInfo struct {
	MemberId           int64    `json:"memberId"`
	MemberName         string   `json:"memberName,omitempty"`
	InternalIdentifier string   `json:"internalIdentifier,omitempty"`
	MemberEmail        string   `json:"memberEmail,omitempty"`
	MemberMobile       string   `json:"memberMobile,omitempty"`
	MemberStatus       string   `json:"memberStatus,omitempty"`
	MemberDeptIds      []int64  `json:"memberDeptIds,omitempty"`
	RoleType           []string `json:"roleType,omitempty"`
}
