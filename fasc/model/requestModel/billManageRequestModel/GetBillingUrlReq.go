package billManageRequestModel

import "github.com/darwinOrg/go-sign-sdk/fasc/model/commonModel"

// GetBillingUrlReq 获取计费链接
type GetBillingUrlReq struct {
	OpenId      *commonModel.OpenId `json:"openId,omitempty"`
	UrlType     string              `json:"urlType,omitempty"`
	RedirectUrl string              `json:"redirectUrl,omitempty"`
}
