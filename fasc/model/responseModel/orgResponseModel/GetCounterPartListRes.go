package orgResponseModel

type GetCounterPartListRes struct {
	RequestId string                 `json:"requestId"`
	Code      string                 `json:"code"`
	Msg       string                 `json:"msg"`
	Data      GetCounterPartListData `json:"data"`
}

type GetCounterPartListData struct {
	CounterPartList []GetCounterPartList `json:"counterpartInfos,omitempty"`
}

type GetCounterPartList struct {
	CorpName    string `json:"corpName,omitempty"`
	CorpIdentNo string `json:"corpIdentNo,omitempty"`
	Status      string `json:"status,omitempty"`
}
