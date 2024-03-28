package signtaskRequestModel

import (
	"github.com/darwinOrg/go-sign-sdk/fasc/model/commonModel"
)

// CreateSignTaskReq 创建签署任务
type CreateSignTaskReq struct {
	SignTaskSubject  string                      `json:"signTaskSubject,omitempty"`
	Initiator        *commonModel.OpenId         `json:"initiator,omitempty"`
	SignDocType      string                      `json:"signDocType,omitempty"`
	ExpiresTime      string                      `json:"expiresTime,omitempty"`
	AutoStart        bool                        `json:"autoStart"`
	AutoFillFinalize bool                        `json:"autoFillFinalize"`
	BusinessTypeId   *int64                      `json:"businessTypeId"`
	SignInOrder      bool                        `json:"signInOrder"`
	AutoFinish       bool                        `json:"autoFinish"`
	CertCAOrg        string                      `json:"certCAOrg,omitempty"`
	CatalogId        string                      `json:"catalogId,omitempty"`
	BusinessScene    *BusinessSceneInfo          `json:"businessScene,omitempty"`
	Docs             []commonModel.Doc           `json:"docs,omitempty"`
	Attachs          []commonModel.Attach        `json:"attachs,omitempty"`
	Actors           []commonModel.SignTaskActor `json:"actors,omitempty"`
}
