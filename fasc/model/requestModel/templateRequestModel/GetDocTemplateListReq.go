package templateRequestModel

import "github.com/darwinOrg/go-sign-sdk/fasc/model/commonModel"

// GetDocTemplateListReq 查询文档模板列表
type GetDocTemplateListReq struct {
	OwnerId      *commonModel.OpenId   `json:"ownerId,omitempty"`
	ListFilter   *DocTemplateLisFilter `json:"listFilter,omitempty"`
	ListPageNo   int                   `json:"listPageNo,omitempty"`
	ListPageSize int                   `json:"listPageSize,omitempty"`
}

type DocTemplateLisFilter struct {
	DocTemplateName string `json:"docTemplateName,omitempty"`
}
