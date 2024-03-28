package signtaskRequestModel

// DeleteSignTaskAttachReq 移除签署任务附件
type DeleteSignTaskAttachReq struct {
	SignTaskId string   `json:"signTaskId,omitempty"`
	AttachIds  []string `json:"attachIds,omitempty"`
}
