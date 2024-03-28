package docProcessRequestModel

import "github.com/darwinOrg/go-sign-sdk/fasc/model/commonModel"

type GetOcrEditExamineResultDataReq struct {
	Initiator *commonModel.OpenId `json:"initiator,omitempty"`
	ExamineId string              `json:"examineId,omitempty"`
}
