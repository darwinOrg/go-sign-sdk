package commonModel

// CorpIdentInfo 企业身份信息
type CorpIdentInfo struct {
	CorpName      string `json:"corpName,omitempty"`
	CorpIdentType string `json:"corpIdentType,omitempty"`
	CorpIdentNo   string `json:"corpIdentNo,omitempty"`
	LegalRepName  string `json:"legalRepName,omitempty"`
}
