package templateResponseModel

import "github.com/darwinOrg/go-sign-sdk/fasc/model/commonModel"

// GetSignTemplateDetailRes 获取签署模板详情
type GetSignTemplateDetailRes struct {
	RequestId string                       `json:"requestId"`
	Code      string                       `json:"code"`
	Msg       string                       `json:"msg"`
	Data      GetSignTemplateDetailResData `json:"data"`
}

type Attach struct {
	AttachId   int    `json:"attachId,omitempty"`
	AttachName string `json:"attachName,omitempty"`
}

type GetSignTemplateDetailResData struct {
	SignTemplateId     string             `json:"signTemplateId"`
	SignTemplateName   string             `json:"signTemplateName"`
	BusinessTypeName   string             `json:"businessTypeName"`
	CatalogName        string             `json:"catalogName,omitempty"`
	SignTemplateStatus string             `json:"signTemplateStatus"`
	CertCAOrg          string             `json:"certCAOrg"`
	Docs               []Doc              `json:"docs"`
	Attachs            []Attach           `json:"attachs"`
	SignInOrder        bool               `json:"signInOrder"`
	Actors             []SignTaskTemActor `json:"actors,omitempty"`
}

type Doc struct {
	DocId     int                 `json:"docId,omitempty"`
	DocName   string              `json:"docName,omitempty"`
	DocFields []commonModel.Field `json:"docFields,omitempty"`
}

type SignTaskTemActor struct {
	ActorInfo      *ActorInfo                  `json:"actorInfo"`
	FillFields     []FillField                 `json:"fillFields"`
	SignFields     []SignField                 `json:"signFields"`
	SignConfigInfo *commonModel.SignConfigInfo `json:"signConfigInfo"`
}

type ActorInfo struct {
	ActorType   string   `json:"actorType,omitempty"`
	ActorId     string   `json:"actorId,omitempty"`
	IsInitiator bool     `json:"isInitiator,omitempty"`
	Permissions []string `json:"permissions,omitempty"`
}

type FillField struct {
	FieldDocId string `json:"fieldDocId,omitempty"`
	FieldId    string `json:"fieldId,omitempty"`
}

type SignField struct {
	FieldDocId string `json:"fieldDocId,omitempty"`
	FieldId    string `json:"fieldId,omitempty"`
}
