package signtaskRequestModel

import "github.com/darwinOrg/go-sign-sdk/fasc/model/commonModel"

// AddSignTaskDocReq 添加签署任务文档信息
type AddSignTaskDocReq struct {
	SignTaskId string            `json:"signTaskId,omitempty"`
	Docs       []commonModel.Doc `json:"docs,omitempty"`
}
