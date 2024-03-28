package templateResponseModel

type AppDocTemplatePageListRes struct {
	RequestId string                     `json:"requestId,omitempty"`
	Code      string                     `json:"code,omitempty"`
	Msg       string                     `json:"msg,omitempty"`
	Data      AppDocTemplatePageListData `json:"data,omitempty"`
}

type AppDocTemplatePageListData struct {
	AppDocTemplates []DocAppTemplate `json:"appDocTemplates,omitempty"`
	ListPageNo      int              `json:"listPageNo,omitempty"`
	CountInPage     int              `json:"countInPage,omitempty"`
	ListPageCount   int              `json:"listPageCount,omitempty"`
	TotalCount      int              `json:"totalCount,omitempty"`
}

type DocAppTemplate struct {
	AppDocTemplateId     string `json:"appDocTemplateId,omitempty"`
	AppDocTemplateName   string `json:"appDocTemplateName,omitempty"`
	AppDocTemplateStatus string `json:"appDocTemplateStatus,omitempty"`
	CreateTime           string `json:"createTime,omitempty"`
	UpdateTime           string `json:"updateTime,omitempty"`
}
