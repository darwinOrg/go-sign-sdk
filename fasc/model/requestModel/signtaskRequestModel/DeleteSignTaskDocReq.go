package signtaskRequestModel

// DeleteSignTaskDocReq 移除签署任务文档
type DeleteSignTaskDocReq struct {
	SignTaskId string   `json:"signTaskId,omitempty"`
	DocIds     []string `json:"docIds,omitempty"`
}
