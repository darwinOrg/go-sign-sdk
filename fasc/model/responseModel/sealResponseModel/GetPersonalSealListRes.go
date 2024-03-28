package sealResponseModel

type GetPersonalSealListRes struct {
	RequestId string              `json:"requestId"`
	Code      string              `json:"code"`
	Msg       string              `json:"msg"`
	Data      GetPersonalSealList `json:"data"`
}

type GetPersonalSealList struct {
	SealInfos []PersonalSealInfo `json:"sealInfos,omitempty"`
}

type PersonalSealInfo struct {
	SealId          int64          `json:"sealId,string"`
	SealName        string         `json:"sealName,omitempty"`
	SealStatus      string         `json:"sealStatus,omitempty"`
	CategoryType    string         `json:"categoryType,omitempty"`
	PicFileUrl      string         `json:"picFileUrl,omitempty"`
	CreateTime      string         `json:"createTime,omitempty"`
	CertCAOrg       string         `json:"certCAOrg,omitempty"`
	CertEncryptType string         `json:"certEncryptType,omitempty"`
	FreeSignInfos   []FreeSignInfo `json:"freeSignInfos,omitempty"`
}
