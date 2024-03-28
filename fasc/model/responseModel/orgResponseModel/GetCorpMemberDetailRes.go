package orgResponseModel

type GetCorpMemberDetailRes struct {
	RequestId string              `json:"requestId"`
	Code      string              `json:"code"`
	Msg       string              `json:"msg"`
	Data      GetCorpMemberDetail `json:"data"`
}

type GetCorpMemberDetail struct {
	MemberId           int64    `json:"memberId,string"`
	MemberName         string   `json:"memberName,omitempty"`
	InternalIdentifier string   `json:"internalIdentifier,omitempty"`
	MemberEmail        string   `json:"memberEmail,omitempty"`
	FddId              string   `json:"fddId,omitempty"`
	MemberStatus       string   `json:"memberStatus,omitempty"`
	MemberDeptIds      []string `json:"memberDeptIds,omitempty"`
	RoleType           []string `json:"roleType,omitempty"`
}
