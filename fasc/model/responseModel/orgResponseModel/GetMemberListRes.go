package orgResponseModel

type GetMemberListRes struct {
	RequestId string        `json:"requestId"`
	Code      string        `json:"code"`
	Msg       string        `json:"msg"`
	Data      GetMemberList `json:"data"`
}
type GetMemberList struct {
	EmployeeInfos []EmployeeInfo `json:"employeeInfos,omitempty"`
	ListPageNo    int            `json:"listPageNo,omitempty"`
	CountInPage   int            `json:"countInPage,omitempty"`
	ListPageCount int            `json:"listPageCount,omitempty"`
	TotalCount    int            `json:"totalCount,omitempty"`
}
