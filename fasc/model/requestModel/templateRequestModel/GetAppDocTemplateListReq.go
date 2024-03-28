package templateRequestModel

type GetAppDocTemplateListReq struct {
	ListFilter *DocListFilterInfo `json:"listFilter,omitempty"`
}

type DocListFilterInfo struct {
	AppDocTemplateName string `json:"appDocTemplateName,omitempty"`
}
