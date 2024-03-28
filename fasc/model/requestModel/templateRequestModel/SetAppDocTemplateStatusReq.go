package templateRequestModel

type SetAppDocTemplateStatusReq struct {
	AppDocTemplateId     string `json:"appDocTemplateId,omitempty"`
	AppDocTemplateStatus string `json:"appDocTemplateStatus,omitempty"`
}
