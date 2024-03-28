package signtaskResponseModel

type GetApprovalInfoRes struct {
	RequestId string              `json:"requestId"`
	Code      string              `json:"code"`
	Msg       string              `json:"msg"`
	Data      GetApprovalInfoData `json:"data"`
}

type GetApprovalInfoData struct {
	SignTaskId      string         `json:"signTaskId,omitempty"`
	SignTaskSubject string         `json:"signTaskSubject,omitempty"`
	ApprovalId      string         `json:"approvalId,omitempty"`
	ApprovalSubject string         `json:"approvalSubject,omitempty"`
	ApplicantName   string         `json:"applicantName,omitempty"`
	ApplicationTime string         `json:"applicationTime,omitempty"`
	ApprovalStatus  string         `json:"approvalStatus,omitempty"`
	ApprovalNode    []ApprovalNode `json:"approvalNode"`
}

type ApprovalNode struct {
	NodeId        int64           `json:"nodeId,string"`
	ApprovalType  string          `json:"approvalType,omitempty"`
	NodeStatus    string          `json:"nodeStatus,omitempty"`
	ApproversInfo []ApproversInfo `json:"approversInfo"`
}

type ApproversInfo struct {
	ApproverName   string `json:"approverName,omitempty"`
	ApproverStatus string `json:"approverStatus,omitempty"`
	OperateTime    string `json:"operateTime,omitempty"`
	RejectNote     string `json:"rejectNote,omitempty"`
}
