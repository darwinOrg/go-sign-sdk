package signtaskResponseModel

type SignTaskCatalogListRes struct {
	RequestId string                    `json:"requestId"`
	Code      string                    `json:"code"`
	Msg       string                    `json:"msg"`
	Data      []SignTaskCatalogListData `json:"data"`
}

type SignTaskCatalogListData struct {
	CatalogId   string `json:"catalogId,omitempty"`
	CatalogName string `json:"catalogName,omitempty"`
}
