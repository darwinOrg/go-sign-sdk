package orgRequestModel

// GetCorpMemberListReq 获取企业成员列表
type GetCorpMemberListReq struct {
	OwnerId      string          `json:"ownerId,omitempty"`
	OpenCorpId   string          `json:"openCorpId,omitempty"`
	ListFilter   *CorpListFilter `json:"listFilter,omitempty"`
	ListPageNo   int             `json:"listPageNo,omitempty"`
	ListPageSize int             `json:"listPageSize,omitempty"`
}

type CorpListFilter struct {
	RoleType   string `json:"roleType,omitempty"`
	DeptId     *int64 `json:"deptId,string"`
	FetchChild bool   `json:"fetchChild"`
}
