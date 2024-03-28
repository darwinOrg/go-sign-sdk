package templateRequestModel

type SetAppSignTemplateStatusReq struct {
	AppSignTemplateId     string `json:"appSignTemplateId,omitempty"`
	AppSignTemplateStatus string `json:"appSignTemplateStatus,omitempty"`
}
