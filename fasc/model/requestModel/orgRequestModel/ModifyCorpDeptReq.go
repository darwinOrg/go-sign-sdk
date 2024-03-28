package orgRequestModel

type ModifyCorpDeptReq struct {
	OpenCorpId   string `json:"openCorpId,omitempty"`
	DeptId       int64  `json:"deptId"`
	DeptName     string `json:"deptName,omitempty"`
	DeptOrderNum int    `json:"deptOrderNum"`
}
