package signtaskRequestModel

// UnBlockSignTaskReq 解除阻塞签署任务
type UnBlockSignTaskReq struct {
	SignTaskId string `json:"signTaskId,omitempty"`
	ActorId    string `json:"actorId,omitempty"`
}
