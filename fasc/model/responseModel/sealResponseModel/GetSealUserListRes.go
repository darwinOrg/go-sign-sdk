package sealResponseModel

// GetSealUserListRes 获取企业用印员列表
type GetSealUserListRes struct {
	RequestId string              `json:"requestId"`
	Code      string              `json:"code"`
	Msg       string              `json:"msg"`
	Data      SealUserListResData `json:"data"`
}

type SealUserListResData struct {
	SealUsers []SealUserInfo `json:"sealUsers,omitempty"`
}

type SealUserInfo struct {
	MemberId           string        `json:"memberId,omitempty"`
	MemberName         string        `json:"memberName,omitempty"`
	InternalIdentifier string        `json:"internalIdentifier,omitempty"`
	MemberEmail        string        `json:"memberEmail,omitempty"`
	MemberMobile       string        `json:"memberMobile,omitempty"`
	SealInfos          []SealArrInfo `json:"sealInfos,omitempty"`
}

type SealArrInfo struct {
	SealId          string `json:"sealId,omitempty"`
	SealName        string `json:"sealName,omitempty"`
	SealStatus      string `json:"sealStatus,omitempty"`
	CategoryType    string `json:"categoryType,omitempty"`
	PicFileUrl      string `json:"picFileUrl,omitempty"`
	CreateTime      string `json:"createTime,omitempty"`
	CertCAOrg       string `json:"certCAOrg,omitempty"`
	CertEncryptType string `json:"certEncryptType,omitempty"`
	GrantTime       string `json:"grantTime,omitempty"`
	GrantStartTime  string `json:"grantStartTime,omitempty"`
	GrantEndTime    string `json:"grantEndTime,omitempty"`
	GrantStatus     string `json:"grantStatus,omitempty"`
}
