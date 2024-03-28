package docProcessRequestModel

import "github.com/darwinOrg/go-sign-sdk/fasc/model/commonModel"

type GetExamineUrlReq struct {
	Initiator *commonModel.OpenId `json:"initiator,omitempty"`
	FileId    string              `json:"fileId,omitempty"`
}
