package orgRequestModel

type GetCorpDeptDetailReq struct {
	OpenCorpId string `json:"openCorpId"`
	DeptId     int64  `json:"deptId"`
}
