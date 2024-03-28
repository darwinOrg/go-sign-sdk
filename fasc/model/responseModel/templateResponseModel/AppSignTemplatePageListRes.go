package templateResponseModel

type AppSignTemplatePageListRes struct {
	RequestId string                      `json:"requestId,omitempty"`
	Code      string                      `json:"code,omitempty"`
	Msg       string                      `json:"msg,omitempty"`
	Data      AppSignTemplatePageListData `json:"data,omitempty"`
}

type AppSignTemplatePageListData struct {
	AppSignTemplates []AppSignTemplates `json:"appSignTemplates,omitempty"`
	ListPageNo       int                `json:"listPageNo,omitempty"`
	CountInPage      int                `json:"countInPage,omitempty"`
	ListPageCount    int                `json:"listPageCount,omitempty"`
	TotalCount       int                `json:"totalCount,omitempty"`
}

type AppSignTemplates struct {
	AppSignTemplateId     string `json:"appSignTemplateId,omitempty"`
	AppSignTemplateName   string `json:"appSignTemplateName,omitempty"`
	AppSignTemplateStatus string `json:"appSignTemplateStatus,omitempty"`
	CreateTime            string `json:"createTime,omitempty"`
	UpdateTime            string `json:"updateTime,omitempty"`
}
