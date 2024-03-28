package orgResponseModel

type CreateCorpDeptRes struct {
	RequestId string         `json:"requestId"`
	Code      string         `json:"code"`
	Msg       string         `json:"msg"`
	Data      CreateCorpDept `json:"data"`
}

type CreateCorpDept struct {
	DeptId       int64 `json:"deptId,string"`
	DeptOrderNum int   `json:"deptOrderNum"`
}
