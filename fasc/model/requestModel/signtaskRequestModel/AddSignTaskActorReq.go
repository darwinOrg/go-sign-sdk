package signtaskRequestModel

import "github.com/darwinOrg/go-sign-sdk/fasc/model/commonModel"

// AddSignTaskActorReq 添加签署任务参与方

type AddSignTaskActorReq struct {
	SignTaskId string                      `json:"signTaskId,omitempty"`
	Actors     []commonModel.SignTaskActor `json:"actors,omitempty"`
}
