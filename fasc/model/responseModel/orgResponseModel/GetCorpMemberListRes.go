package orgResponseModel

// GetCorpMemberListRes 获取企业成员列表
type GetCorpMemberListRes struct {
	RequestId string                   `json:"requestId"`
	Code      string                   `json:"code"`
	Msg       string                   `json:"msg"`
	Data      GetCorpMemberListResData `json:"data"`
}

type GetCorpMemberListResData struct {
	EmployeeInfos []EmployeeInfo `json:"employeeInfos"`
	ListPageNo    int            `json:"listPageNo"`
	CountInPage   int            `json:"countInPage"`
	ListPageCount int            `json:"listPageCount"`
	TotalCount    int            `json:"totalCount"`
}

type EmployeeInfo struct {
	MemberId           int64    `json:"memberId,string"`
	MemberName         string   `json:"memberName"`
	InternalIdentifier string   `json:"internalIdentifier"`
	MemberEmail        string   `json:"memberEmail"`
	MemberStatus       string   `json:"memberStatus"`
	MemberDeptIds      []string `json:"memberDeptIds,string"`
	RoleType           []string `json:"roleType"`
}
