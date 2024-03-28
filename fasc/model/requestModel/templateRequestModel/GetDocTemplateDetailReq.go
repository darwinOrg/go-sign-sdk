package templateRequestModel

import "github.com/darwinOrg/go-sign-sdk/fasc/model/commonModel"

// GetDocTemplateDetailReq 获取文档模板详情
type GetDocTemplateDetailReq struct {
	OwnerId       *commonModel.OpenId `json:"ownerId,omitempty"`
	DocTemplateId string              `json:"docTemplateId,omitempty"`
}
