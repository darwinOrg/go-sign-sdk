package signtaskRequestModel

// BlockSignTaskReq 阻塞签署任务
type BlockSignTaskReq struct {
	SignTaskId string `json:"signTaskId,omitempty"`
	ActorId    string `json:"actorId,omitempty"`
}
