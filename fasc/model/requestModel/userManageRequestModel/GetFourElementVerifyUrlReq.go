package userManageRequestModel

import "github.com/darwinOrg/go-sign-sdk/fasc/model/commonModel"

type GetFourElementVerifyUrlReq struct {
	ClientUserId       string              `json:"clientUserId,omitempty"`
	Initiator          *commonModel.OpenId `json:"initiator,omitempty"`
	UserName           string              `json:"userName,omitempty"`
	UserIdentNo        string              `json:"userIdentNo,omitempty"`
	IdCardImage        bool                `json:"idCardImage"`
	FaceSide           string              `json:"faceSide,omitempty"`
	NationalEmblemSide string              `json:"nationalEmblemSide,omitempty"`
	BankAccountNo      string              `json:"bankAccountNo,omitempty"`
	Mobile             string              `json:"mobile,omitempty"`
	RedirectUrl        string              `json:"redirectUrl,omitempty"`
}
