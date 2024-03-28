package templateRequestModel

import "github.com/darwinOrg/go-sign-sdk/fasc/model/commonModel"

// GetSignTemplateDetailReq 获取签署模板详情
type GetSignTemplateDetailReq struct {
	OwnerId        *commonModel.OpenId `json:"ownerId,omitempty"`
	SignTemplateId string              `json:"signTemplateId,omitempty"`
}
