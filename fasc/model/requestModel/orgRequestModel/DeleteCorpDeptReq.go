package orgRequestModel

type DeleteCorpDeptReq struct {
	OpenCorpId string `json:"openCorpId,omitempty"`
	DeptId     int64  `json:"deptId,omitempty"`
}
