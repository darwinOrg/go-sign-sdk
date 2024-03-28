package templateResponseModel

type GetDocTemplateListRes struct {
	RequestId string                    `json:"requestId"`
	Code      string                    `json:"code"`
	Msg       string                    `json:"msg"`
	Data      GetDocTemplateListResData `json:"data"`
}

type GetDocTemplateListResData struct {
	DocTemplates  []DocTemplates `json:"docTemplates"`
	ListPageNo    int            `json:"listPageNo"`
	CountInPage   int            `json:"countInPage"`
	ListPageCount int            `json:"listPageCount"`
	TotalCount    int            `json:"totalCount"`
}

type DocTemplates struct {
	DocTemplateId     string `json:"docTemplateId"`
	DocTemplateName   string `json:"docTemplateName"`
	CatalogName       string `json:"catalogName,omitempty"`
	DocTemplateStatus string `json:"docTemplateStatus"`
	CreateTime        string `json:"createTime"`
	UpdateTime        string `json:"updateTime"`
}
