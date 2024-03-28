package orgRequestModel

type CreateCorpDeptReq struct {
	OpenCorpId   string `json:"openCorpId,omitempty"`
	ParentDeptId *int64 `json:"parentDeptId,omitempty"`
	DeptName     string `json:"deptName,omitempty"`
	DeptOrderNum int    `json:"deptOrderNum,omitempty"`
}
