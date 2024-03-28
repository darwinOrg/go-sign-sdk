package signtaskRequestModel

type SignTaskActorGetUrlReq struct {
	SignTaskId   string `json:"signTaskId,omitempty'"`
	ActorId      string `json:"actorId,omitempty"`
	ClientUserId string `json:"clientUserId,omitempty"`
	RedirectUrl  string `json:"redirectUrl,omitempty"`
}
