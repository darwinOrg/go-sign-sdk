package docProcessRequestModel

import "github.com/darwinOrg/go-sign-sdk/fasc/model/commonModel"

type GetCompareUrlReq struct {
	Initiator    *commonModel.OpenId `json:"initiator,omitempty"`
	OriginFileId string              `json:"originFileId,omitempty"`
	TargetFileId string              `json:"targetFileId,omitempty"`
}
