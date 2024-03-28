package sealRequestModel

import "github.com/darwinOrg/go-sign-sdk/fasc/model/commonModel"

type GetSealGrantUrlReq struct {
	OpenCorpId   string                  `json:"openCorpId,omitempty"`
	SealId       int64                   `json:"sealId,string"`
	MemberInfo   *commonModel.MemberInfo `json:"memberInfo,omitempty"`
	ClientUserId string                  `json:"clientUserId,omitempty"`
	RedirectUrl  string                  `json:"redirectUrl,omitempty"`
}
