package commonModel

// UserIdentInfo 个人用户信息
type UserIdentInfo struct {
	UserName  string `json:"userName,omitempty"`
	IdentType string `json:"identType,omitempty"`
	IdentNo   string `json:"identNo,omitempty"`
}
