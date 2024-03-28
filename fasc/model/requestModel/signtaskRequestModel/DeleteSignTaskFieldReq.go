package signtaskRequestModel

type DeleteSignTaskField struct {
	DocId     string   `json:"docId,omitempty"`
	DocFields []string `json:"fieldIds,omitempty"`
}

// DeleteSignTaskFieldReq 移除签署任务文档
type DeleteSignTaskFieldReq struct {
	SignTaskId string                `json:"signTaskId,omitempty"`
	Fields     []DeleteSignTaskField `json:"fields,omitempty"`
}
