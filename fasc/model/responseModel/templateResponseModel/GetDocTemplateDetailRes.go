package templateResponseModel

import "github.com/darwinOrg/go-sign-sdk/fasc/model/commonModel"

type GetDocTemplateDetailRes struct {
	RequestId string                      `json:"requestId"`
	Code      string                      `json:"code"`
	Msg       string                      `json:"msg"`
	Data      GetDocTemplateDetailResData `json:"data"`
}

type GetDocTemplateDetailResData struct {
	DocTemplateId     string              `json:"docTemplateId"`
	DocTemplateName   string              `json:"docTemplateName"`
	CatalogName       string              `json:"catalogName,omitempty"`
	DocTemplateStatus string              `json:"docTemplateStatus"`
	DocFields         []commonModel.Field `json:"docFields"`
}
