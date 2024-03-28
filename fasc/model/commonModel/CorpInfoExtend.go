package commonModel

// CorpInfoExtend 企业信息补充
type CorpInfoExtend struct {
	BankName         string `json:"bankName,omitempty"`
	BankBranchName   string `json:"bankBranchName,omitempty"`
	BankAccountNo    string `json:"bankAccountNo,omitempty"`
	BankProvinceName string `json:"bankProvinceName,omitempty"`
	BankCityName     string `json:"bankCityName,omitempty"`
}
