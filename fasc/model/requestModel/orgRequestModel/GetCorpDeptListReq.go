package orgRequestModel

type Integer int

type GetCorpDeptListReq struct {
	OpenCorpId   string `json:"openCorpId,omitempty"`
	ParentDeptId *int64 `json:"parentDeptId,omitempty"`
}
