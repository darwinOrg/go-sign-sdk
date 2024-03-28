package commonModel

// Actor 签署任务参与方
type Actor struct {
	ActorType         string            `json:"actorType,omitempty"`
	ActorId           string            `json:"actorId,omitempty"`
	ActorName         string            `json:"actorName,omitempty"`
	Permissions       []string          `json:"permissions,omitempty"`
	ActorOpenId       string            `json:"actorOpenId,omitempty"`
	ActorFDDId        string            `json:"actorFDDId,omitempty"`
	IdentNameForMatch string            `json:"identNameForMatch,omitempty"`
	CertNoForMatch    string            `json:"certNoForMatch,omitempty"`
	ActorCorpMember   []ActorCorpMember `json:"actorCorpMembers,omitempty"`
	Notification      *Notification     `json:"notification,omitempty"`
}
