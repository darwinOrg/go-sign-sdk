package userManageRequestModel

// GetUserAuthUrlReq 获取个人用户授权链接
type GetUserAuthUrlReq struct {
	ClientUserId       string         `json:"clientUserId,omitempty"`
	UserName           string         `json:"userName,omitempty"`
	UserIdentType      string         `json:"userIdentType,omitempty"`
	UserIdentNo        string         `json:"userIdentNo,omitempty"`
	UserIdentInfoMatch bool           `json:"userIdentInfoMatch"`
	AuthScopes         []string       `json:"authScopes,omitempty"`
	RedirectUrl        string         `json:"redirectUrl,omitempty"`
	AccountName        string         `json:"accountName,omitempty"`
	UserIdentInfo      *UserIdentInfo `json:"userIdentInfo,omitempty"`
	NonEditableInfo    []string       `json:"nonEditableInfo,omitempty"`
}

type UserIdentInfo struct {
	UserName      string   `json:"userName,omitempty"`
	UserIdentType string   `json:"userIdentType,omitempty"`
	UserIdentNo   string   `json:"userIdentNo,omitempty"`
	Mobile        string   `json:"mobile,omitempty"`
	BankAccountNo string   `json:"bankAccountNo,omitempty"`
	IdentMethod   []string `json:"identMethod,omitempty"`
}
