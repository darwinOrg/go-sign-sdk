package templateRequestModel

type GetAppSignTemplateListReq struct {
	ListFilter *SignListFilterInfo `json:"listFilter,omitempty"`
}

type SignListFilterInfo struct {
	AppSignTemplateName string `json:"appSignTemplateName,omitempty"`
}
