package commonModel

type MemberInfo struct {
	MemberIds      []int64 `json:"memberIds,omitempty"`
	GrantStartTime string  `json:"grantStartTime,omitempty"`
	GrantEndTime   string  `json:"grantEndTime,omitempty"`
}
