package signtaskResponseModel

// GetOwnerSignTaskListRes 获取指定归属方的签署任务列表
type GetOwnerSignTaskListRes struct {
	RequestId string                   `json:"requestId"`
	Code      string                   `json:"code"`
	Msg       string                   `json:"msg"`
	Data      OwnerSignTaskListResData `json:"data"`
}

type OwnerSignTaskListResData struct {
	SignTasks     []SignTask `json:"signTasks,omitempty"`
	ListPageNo    int        `json:"listPageNo,omitempty"`
	CountInPage   int        `json:"countInPage,omitempty"`
	ListPageCount int        `json:"listPageCount,omitempty"`
	TotalCount    int        `json:"totalCount,omitempty"`
}

type SignTask struct {
	SignTaskId          string `json:"signTaskId,omitempty"`
	TransReferenceId    string `json:"transReferenceId,omitempty"`
	SignTaskSource      string `json:"signTaskSource,omitempty"`
	SignTaskSubject     string `json:"signTaskSubject,omitempty"`
	ApprovalStatus      string `json:"approvalStatus,omitempty"`
	RejectNote          string `json:"rejectNote,omitempty"`
	CatalogId           string `json:"catalogId,omitempty"`
	CatalogName         string `json:"catalogName,omitempty"`
	BusinessTypeId      int64  `json:"businessTypeId,omitempty"`
	BusinessTypeName    string `json:"businessTypeName,omitempty"`
	BusinessCode        string `json:"businessCode,omitempty"`
	TemplateId          string `json:"templateId,omitempty"`
	SignTaskStatus      string `json:"signTaskStatus,omitempty"`
	AbolishedSignTaskId string `json:"abolishedSignTaskId,omitempty"`
	OriginalSignTaskId  string `json:"originalSignTaskId,omitempty"`
	TerminationNote     string `json:"terminationNote,omitempty"`
	RevokeReason        string `json:"revokeReason,omitempty"`
	InitiatorName       string `json:"initiatorName,omitempty"`
	CreateTime          string `json:"createTime,omitempty"`
	StartTime           string `json:"startTime,omitempty"`
	FinishTime          string `json:"finishTime,omitempty"`
	DeadlineTime        string `json:"deadlineTime,omitempty"`
}
