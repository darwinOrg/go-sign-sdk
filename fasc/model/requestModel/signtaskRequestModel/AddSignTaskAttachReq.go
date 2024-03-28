package signtaskRequestModel

import "github.com/darwinOrg/go-sign-sdk/fasc/model/commonModel"

// AddSignTaskAttachReq 添加签署任务附件

type AddSignTaskAttachReq struct {
	SignTaskId string               `json:"signTaskId,omitempty"`
	Attachs    []commonModel.Attach `json:"attachs,omitempty"`
}
