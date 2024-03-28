package templateResponseModel

import "github.com/darwinOrg/go-sign-sdk/fasc/model/commonModel"

type AppDocTemplateDetailRes struct {
	RequestId string                   `json:"requestId,omitempty"`
	Code      string                   `json:"code,omitempty"`
	Msg       string                   `json:"msg,omitempty"`
	Data      AppDocTemplateDetailData `json:"data,omitempty"`
}

type AppDocTemplateDetailData struct {
	AppDocTemplateId     string              `json:"appDocTemplateId,omitempty"`
	AppDocTemplateName   string              `json:"appDocTemplateName,omitempty"`
	AppDocTemplateStatus string              `json:"appDocTemplateStatus,omitempty"`
	DocFields            []commonModel.Field `json:"docFields,omitempty"`
}
