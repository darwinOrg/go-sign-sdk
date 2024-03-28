package commonModel

// UserInfoExtend 个人用户信息补充
type UserInfoExtend struct {
	BankAccountNo string `json:"bankAccountNo,omitempty"`
	Mobile        string `json:"mobile,omitempty"`
}
