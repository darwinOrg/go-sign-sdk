package sealResponseModel

type GetSealDetailRes struct {
	RequestId string        `json:"requestId"`
	Code      string        `json:"code"`
	Msg       string        `json:"msg"`
	Data      GetSealDetail `json:"data"`
}

type GetSealDetail struct {
	SealInfo *SealDetailInfo `json:"sealInfo,omitempty"`
}

type SealDetailInfo struct {
	SealUsers     []SealUserDetail `json:"sealUsers,omitempty"`
	FreeSignInfos []FreeSignInfo   `json:"freeSignInfos,omitempty"`
}

type SealUserDetail struct {
	MemberId           int64  `json:"memberId,string"`
	MemberName         string `json:"memberName,omitempty"`
	InternalIdentifier string `json:"internalIdentifier,omitempty"`
	MemberEmail        string `json:"memberEmail,omitempty"`
	GrantTime          string `json:"grantTime,omitempty"`
	GrantStartTime     string `json:"grantStartTime,omitempty"`
	GrantEndTime       string `json:"grantEndTime,omitempty"`
	GrantStatus        string `json:"grantStatus,omitempty"`
}

type FreeSignInfo struct {
	BusinessId    string `json:"businessId,omitempty"`
	BusinessScene string `json:"businessScene,omitempty"`
}
