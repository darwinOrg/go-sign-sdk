package signtaskRequestModel

import (
	"github.com/darwinOrg/go-sign-sdk/fasc/model/commonModel"
)

// GetOwnerSignTaskListReq 获取指定归属方的签署任务列表
type GetOwnerSignTaskListReq struct {
	OwnerId      *commonModel.OpenId `json:"ownerId,omitempty"`
	OwnerRole    string              `json:"ownerRole,omitempty"`
	ListFilter   *SignTaskListFilter `json:"listFilter,omitempty"`
	ListPageNo   int                 `json:"listPageNo,omitempty"`
	ListPageSize int                 `json:"listPageSize,omitempty"`
}

type SignTaskListFilter struct {
	SignTaskSubject string   `json:"signTaskSubject,omitempty"`
	SignTaskStatus  []string `json:"signTaskStatus,omitempty"`
}
