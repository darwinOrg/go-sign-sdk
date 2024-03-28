package userManageResponseModel

// UnBindUserRes 解绑个人用户
type UnBindUserRes struct {
	RequestId string `json:"requestId"`
	Code      string `json:"code"`
	Msg       string `json:"msg"`
}
