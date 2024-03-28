package signtaskRequestModel

type GetActorFacePictureReq struct {
	SignTaskId string `json:"signTaskId,omitempty"`
	ActorId    string `json:"actorId,omitempty"`
}
