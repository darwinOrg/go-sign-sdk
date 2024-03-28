package templateResponseModel

type SignTemplateListInfo struct {
	SignTemplateId     string `json:"signTemplateId,omitempty"`
	SignTemplateName   string `json:"signTemplateName,omitempty"`
	BusinessTypeName   string `json:"businessTypeName,omitempty"`
	CatalogName        string `json:"catalogName,omitempty"`
	SignTemplateStatus string `json:"signTemplateStatus,omitempty"`
	CreateTime         string `json:"createTime,omitempty"`
	UpdateTime         string `json:"updateTime,omitempty"`
}

type SignTemListResData struct {
	SignTemplates []SignTemplateListInfo `json:"signTemplates,omitempty"`
	ListPageNo    int                    `json:"listPageNo,omitempty"`
	CountInPage   int                    `json:"countInPage,omitempty"`
	ListPageCount int                    `json:"listPageCount,omitempty"`
	TotalCount    int                    `json:"totalCount,omitempty"`
}

type GetSignTemplateListRes struct {
	RequestId string             `json:"requestId,omitempty"`
	Code      string             `json:"code,omitempty"`
	Msg       string             `json:"msg,omitempty"`
	Data      SignTemListResData `json:"data,omitempty"`
}
