package userManageRequestModel

// GetUserReq 查询个人用户基本信息
type GetUserReq struct {
	OpenUserId   string `json:"openUserId,omitempty"`
	ClientUserId string `json:"clientUserId,omitempty"`
}
