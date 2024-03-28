package signtaskRequestModel

import "github.com/darwinOrg/go-sign-sdk/fasc/model/commonModel"

// GetOwnerSignTaskUrlReq 获取指定归属方的签署任务文档下载地址
type GetOwnerSignTaskUrlReq struct {
	OwnerId    *commonModel.OpenId `json:"ownerId,omitempty"`
	SignTaskId string              `json:"signTaskId,omitempty"`
	FileType   string              `json:"fileType,omitempty"`
	Id         string              `json:"id,omitempty"`
}
