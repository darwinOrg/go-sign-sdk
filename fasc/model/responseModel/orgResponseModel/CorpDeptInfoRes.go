package orgResponseModel

type GetCorpDeptListRes struct {
	RequestId string         `json:"requestId"`
	Code      string         `json:"code"`
	Msg       string         `json:"msg"`
	Data      []CorpDeptInfo `json:"data"`
}
type CorpDeptInfo struct {
	DeptId       int64  `json:"deptId,string"`
	DeptName     string `json:"deptName,omitempty"`
	DeptOrderNum int    `json:"deptOrderNum"`
	ParentDeptId *int64 `json:"parentDeptId,string"`
}
