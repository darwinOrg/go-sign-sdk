package signtaskResponseModel

type ListSignTaskActorRes struct {
	RequestId string                  `json:"requestId"`
	Code      string                  `json:"code"`
	Msg       string                  `json:"msg"`
	Data      []ListSignTaskActorData `json:"data"`
}

type ListSignTaskActorData struct {
	ActorType         string        `json:"actorType,omitempty"`
	ActorId           string        `json:"actorId,omitempty"`
	ActorName         string        `json:"actorName,omitempty"`
	CorpName          string        `json:"corpName,omitempty"`
	CorpIdentNo       string        `json:"corpIdentNo,omitempty"`
	MemberInfo        *MemberInfo   `json:"memberInfo,omitempty"`
	Permissions       []string      `json:"permissions,omitempty"`
	IsInitiator       bool          `json:"isInitiator"`
	SignOrderNo       int           `json:"signOrderNo,omitempty"`
	IdentNameForMatch string        `json:"identNameForMatch,omitempty"`
	CertNoForMatch    string        `json:"certNoForMatch,omitempty"`
	ReadStatus        string        `json:"readStatus,omitempty"`
	ReadTime          string        `json:"readTime,omitempty"`
	JoinStatus        string        `json:"joinStatus,omitempty"`
	JoinTime          string        `json:"joinTime,omitempty"`
	JoinName          string        `json:"joinName,omitempty"`
	JoinIdentNo       string        `json:"joinIdentNo,omitempty"`
	FillStatus        string        `json:"fillStatus,omitempty"`
	FillTime          string        `json:"fillTime,omitempty"`
	SignStatus        string        `json:"signStatus,omitempty"`
	ActorNote         string        `json:"actorNote,omitempty"`
	SignTime          string        `json:"signTime,omitempty"`
	BlockStatus       string        `json:"blockStatus,omitempty"`
	Notification      *Notification `json:"notification,omitempty"`
}

type MemberInfo struct {
	MemberName string `json:"memberName,omitempty"`
	Mobile     string `json:"mobile,omitempty"`
}

type Notification struct {
	SendNotification bool   `json:"sendNotification,omitempty"`
	NotifyWay        string `json:"notifyWay,omitempty"`
	NotifyAddress    string `json:"notifyAddress,omitempty"`
}
