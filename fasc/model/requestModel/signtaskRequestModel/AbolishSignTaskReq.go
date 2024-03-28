package signtaskRequestModel

import "github.com/darwinOrg/go-sign-sdk/fasc/model/commonModel"

type AbolishSignTaskReq struct {
	SignTaskId         string              `json:"signTaskId,omitempty"`
	AbolishedInitiator *AbolishedInitiator `json:"abolishedInitiator,omitempty"`
	DocSource          string              `json:"docSource,omitempty"`
	Reason             string              `json:"reason,omitempty"`
	AutoStart          bool                `json:"autoStart"`
	SignTaskSubject    string              `json:"signTaskSubject,omitempty"`
	ExpiresTime        string              `json:"expiresTime,omitempty"`
	BusinessTypeId     *int64              `json:"businessTypeId,omitempty"`
	CatalogId          string              `json:"catalogId,omitempty"`
	BusinessScene      *BusinessSceneInfo  `json:"businessScene,omitempty"`
	CertCAOrg          string              `json:"certCAOrg,omitempty"`
	Docs               []Doc               `json:"docs,omitempty"`
	Actors             []SignTaskActor     `json:"actors,omitempty"`
}

type AbolishedInitiator struct {
	InitiatorId string `json:"initiatorId,omitempty"`
	ActorId     string `json:"actorId,omitempty"`
}

type Doc struct {
	DocId         string              `json:"docId,omitempty"`
	DocName       string              `json:"docName,omitempty"`
	DocFileId     string              `json:"docFileId,omitempty"`
	DocTemplateId string              `json:"docTemplateId,omitempty"`
	DocFields     []commonModel.Field `json:"docFields,omitempty"`
}

type SignTaskActor struct {
	Actor          *commonModel.Actor `json:"actor"`
	SignFields     []SignField        `json:"signFields"`
	SignConfigInfo *SignConfigInfo    `json:"signConfigInfo"`
}

type SignField struct {
	FieldDocId string `json:"fieldDocId,omitempty"`
	FieldId    string `json:"fieldId,omitempty"`
	FieldName  string `json:"fieldName,omitempty"`
	SealId     *int64 `json:"sealId,omitempty"`
}

type SignConfigInfo struct {
	RequestVerifyFree bool `json:"requestVerifyFree"`
}
