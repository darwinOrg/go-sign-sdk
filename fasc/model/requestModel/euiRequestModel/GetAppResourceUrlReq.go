package euiRequestModel

import "github.com/darwinOrg/go-sign-sdk/fasc/model/commonModel"

// GetAppResourceUrlReq 获取应用级资源访问链接
type GetAppResourceUrlReq struct {
	OpenId   *commonModel.OpenId `json:"openId,omitempty"`
	Resource *AppResource        `json:"resource,omitempty"`
}

type AppResource struct {
	ResourceId string `json:"resourceId,omitempty"`
	Action     string `json:"action,omitempty"`
	Params     string `json:"params,omitempty"`
}
