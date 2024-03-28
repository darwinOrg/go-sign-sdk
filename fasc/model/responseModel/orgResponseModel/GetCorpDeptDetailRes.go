package orgResponseModel

type GetCorpDeptDetailRes struct {
	RequestId string            `json:"requestId"`
	Code      string            `json:"code"`
	Msg       string            `json:"msg"`
	Data      GetCorpDeptDetail `json:"data"`
}

type GetCorpDeptDetail struct {
	DeptId       int64  `json:"deptId,string"`
	DeptName     string `json:"deptName,omitempty"`
	DeptOrderNum int    `json:"deptOrderNum"`
	ParentDeptId int64  `json:"parentDeptId,string"`
}
