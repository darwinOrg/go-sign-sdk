package corpManageRequestModel

type GetIdentifiedStatusReq struct {
	CorpName    string `json:"corpName"`
	CorpIdentNo string `json:"corpIdentNo"`
}
