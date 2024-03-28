package docProcessRequestModel

import "github.com/darwinOrg/go-sign-sdk/fasc/model/commonModel"

type GetCompareResultUrlReq struct {
	Initiator *commonModel.OpenId `json:"initiator,omitempty"`
	CompareId string              `json:"compareId,omitempty"`
}
