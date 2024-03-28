package corpManageRequestModel

// GetCorpAuthUrlReq 获取企业用户授权链接
type GetCorpAuthUrlReq struct {
	ClientCorpId        string            `json:"clientCorpId,omitempty"`
	ClientUserId        string            `json:"clientUserId,omitempty"`
	AccountName         string            `json:"accountName,omitempty"`
	CorpIdentInfo       *CorpIdentInfoReq `json:"corpIdentInfo,omitempty"`
	CorpNonEditableInfo []string          `json:"corpNonEditableInfo,omitempty"`
	OprIdentInfo        *OprIdentInfoReq  `json:"oprIdentInfo"`
	OprNonEditableInfo  []string          `json:"oprNonEditableInfo,omitempty"`
	CorpName            string            `json:"corpName,omitempty"`
	CorpIdentType       string            `json:"corpIdentType,omitempty"`
	CorpIdentNo         string            `json:"corpIdentNo,omitempty"`
	CorpIdentInfoMatch  bool              `json:"corpIdentInfoMatch"`
	AuthScopes          []string          `json:"authScopes,omitempty"`
	RedirectUrl         string            `json:"redirectUrl,omitempty"`
}

type CorpIdentInfoReq struct {
	CorpName        string   `json:"corpName,omitempty"`
	CorpIdentType   string   `json:"corpIdentType,omitempty"`
	CorpIdentNo     string   `json:"corpIdentNo,omitempty"`
	LegalRepName    string   `json:"legalRepName,omitempty"`
	License         string   `json:"license,omitempty"`
	CorpIdentMethod []string `json:"corpIdentMethod,omitempty"`
}

type OprIdentInfoReq struct {
	UserName       string   `json:"userName,omitempty"`
	UserIdentType  string   `json:"userIdentType,omitempty"`
	UserIdentNo    string   `json:"userIdentNo,omitempty"`
	Mobile         string   `json:"mobile,omitempty"`
	BankAccountNo  string   `json:"bankAccountNo,omitempty"`
	OprIdentMethod []string `json:"oprIdentMethod,omitempty"`
}
