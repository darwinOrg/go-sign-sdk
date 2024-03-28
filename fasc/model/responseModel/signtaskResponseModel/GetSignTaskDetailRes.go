package signtaskResponseModel

import (
	"github.com/darwinOrg/go-sign-sdk/fasc/model/commonModel"
)

type GetSignTaskDetailRes struct {
	RequestId string                `json:"requestId"`
	Code      string                `json:"code"`
	Msg       string                `json:"msg"`
	Data      SignTaskDetailResData `json:"data"`
}

type SignTaskDetailResData struct {
	Initiator           *commonModel.OpenId   `json:"initiator"`
	SignTaskId          string                `json:"signTaskId"`
	SignTaskSubject     string                `json:"signTaskSubject"`
	StorageType         string                `json:"storageType,omitempty"`
	AutoFinish          bool                  `json:"autoFinish"`
	SignDocType         string                `json:"signDocType,omitempty"`
	SignTaskSource      string                `json:"signTaskSource,omitempty"`
	CertCAOrg           string                `json:"certCAOrg,omitempty"`
	AutoFillFinalize    bool                  `json:"autoFillFinalize"`
	SignInOrder         bool                  `json:"signInOrder"`
	ApprovalStatus      string                `json:"approvalStatus,omitempty"`
	RejectNote          string                `json:"rejectNote,omitempty"`
	BusinessTypeId      int64                 `json:"businessTypeId,omitempty"`
	BusinessTypeName    string                `json:"businessTypeName,omitempty"`
	BusinessCode        string                `json:"businessCode,omitempty"`
	TemplateId          string                `json:"templateId,omitempty"`
	SignTaskStatus      string                `json:"signTaskStatus"`
	AbolishedSignTaskId string                `json:"abolishedSignTaskId,omitempty"`
	OriginalSignTaskId  string                `json:"originalSignTaskId,omitempty"`
	RevokeReason        string                `json:"revokeReason,omitempty"`
	TerminationNote     string                `json:"terminationNote,omitempty"`
	CreateTime          string                `json:"createTime,omitempty"`
	StartTime           string                `json:"startTime,omitempty"`
	FinishTime          string                `json:"finishTime,omitempty"`
	DeadlineTime        string                `json:"deadlineTime,omitempty"`
	Docs                []Doc                 `json:"docs"`
	Attachs             []Attach              `json:"attachs"`
	Actors              []SignTaskDetailActor `json:"actors"`
}

type Doc struct {
	DocId   string `json:"docId,omitempty"`
	DocName string `json:"docName,omitempty"`
}

type Attach struct {
	AttachId   string `json:"attachId,omitempty"`
	AttachName string `json:"attachName,omitempty"`
}

type SignTaskDetailActor struct {
	ActorInfo             Actor       `json:"actorInfo,omitempty"`
	SignFields            []SignField `json:"signFields,omitempty"`
	ReadStatus            string      `json:"readStatus,omitempty"`
	ReadTime              string      `json:"readTime,omitempty"`
	JoinStatus            string      `json:"joinStatus,omitempty"`
	JoinTime              string      `json:"joinTime,omitempty"`
	FillStatus            string      `json:"fillStatus,omitempty"`
	FillTime              string      `json:"fillTime,omitempty"`
	SignStatus            string      `json:"signStatus,omitempty"`
	ActorNote             string      `json:"actorNote,omitempty"`
	SignTime              string      `json:"signTime,omitempty"`
	SignOrderNo           int         `json:"signOrderNo,omitempty"`
	BlockStatus           string      `json:"blockStatus,omitempty"`
	ActorSignTaskUrl      string      `json:"actorSignTaskUrl,omitempty"`
	ActorSignTaskEmbedUrl string      `json:"actorSignTaskEmbedUrl,omitempty"`
}

type Actor struct {
	ActorType   string   `json:"actorType,omitempty"`
	ActorId     string   `json:"actorId,omitempty"`
	ActorName   string   `json:"actorName,omitempty"`
	Permissions []string `json:"permissions,omitempty"`
	IsInitiator bool     `json:"isInitiator"`
}

type SignField struct {
	SignFieldStatus string                     `json:"signFieldStatus,omitempty"`
	FieldDocId      string                     `json:"fieldDocId,omitempty"`
	FieldId         string                     `json:"fieldId,omitempty"`
	FieldName       string                     `json:"fieldName,omitempty"`
	SealId          *int64                     `json:"sealId,omitempty"`
	Moveable        bool                       `json:"moveable"`
	SignRemark      string                     `json:"signRemark"`
	Position        *commonModel.FieldPosition `json:"position,omitempty"`
}
