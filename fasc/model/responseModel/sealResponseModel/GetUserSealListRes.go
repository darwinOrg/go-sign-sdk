package sealResponseModel

type GetUserSealListRes struct {
	RequestId string          `json:"requestId"`
	Code      string          `json:"code"`
	Msg       string          `json:"msg"`
	Data      GetUserSealList `json:"data"`
}

type GetUserSealList struct {
	SealInfos []GetUserListSealInfo `json:"sealInfos,omitempty"`
}

type GetUserListSealInfo struct {
	SealId          int64  `json:"sealId,string"`
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
