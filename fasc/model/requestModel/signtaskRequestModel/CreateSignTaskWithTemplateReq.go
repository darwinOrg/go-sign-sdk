package signtaskRequestModel

import (
	"github.com/darwinOrg/go-sign-sdk/fasc/model/commonModel"
)

// CreateSignTaskWithTemplateReq 创建签署任务基于签署模板
type CreateSignTaskWithTemplateReq struct {
	SignTaskSubject  string                      `json:"signTaskSubject,omitempty"`
	Initiator        *commonModel.OpenId         `json:"initiator,omitempty"`
	SignDocType      string                      `json:"signDocType,omitempty"`
	SignTemplateId   string                      `json:"signTemplateId,omitempty"`
	ExpiresTime      string                      `json:"expiresTime,omitempty"`
	AutoStart        bool                        `json:"autoStart"`
	AutoFillFinalize bool                        `json:"autoFillFinalize"`
	AutoFinish       bool                        `json:"autoFinish"`
	CertCAOrg        string                      `json:"certCAOrg,omitempty"`
	CatalogId        string                      `json:"catalogId,omitempty"`
	BusinessScene    *BusinessSceneInfo          `json:"businessScene,omitempty"`
	Actors           []commonModel.SignTaskActor `json:"actors,omitempty"`
}

type BusinessSceneInfo struct {
	BusinessId       string `json:"businessId,omitempty"`
	TransReferenceId string `json:"transReferenceId,omitempty"`
}
