package signtaskRequestModel

import "github.com/darwinOrg/go-sign-sdk/fasc/model/commonModel"

type AddSignTaskField struct {
	DocId     string              `json:"docId,omitempty"`
	DocFields []commonModel.Field `json:"docFields,omitempty"`
}

// AddSignTaskFieldReq 添加签署任务控件
type AddSignTaskFieldReq struct {
	SignTaskId string             `json:"signTaskId,omitempty"`
	ActorId    string             `json:"actorId,omitempty"`
	Fields     []AddSignTaskField `json:"fields,omitempty"`
}
