package sealResponseModel

// GetSealListRes 获取印章列表
type GetSealListRes struct {
	RequestId string          `json:"requestId"`
	Code      string          `json:"code"`
	Msg       string          `json:"msg"`
	Data      SealListResData `json:"data"`
}

type SealListResData struct {
	SealInfos []SealInfo `json:"sealInfos,omitempty"`
}

type SealInfo struct {
	SealId          string     `json:"sealId"`
	SealName        string     `json:"sealName"`
	SealStatus      string     `json:"sealStatus"`
	CategoryType    string     `json:"categoryType"`
	PicFileUrl      string     `json:"picFileUrl"`
	CreateTime      string     `json:"createTime"`
	CertCAOrg       string     `json:"certCAOrg"`
	CertEncryptType string     `json:"certEncryptType"`
	SealUsers       []SealUser `json:"sealUsers"`
}

type SealUser struct {
	MemberId           string `json:"memberId"`
	MemberName         string `json:"memberName"`
	InternalIdentifier string `json:"internalIdentifier"`
	MemberEmail        string `json:"memberEmail"`
	GrantTime          string `json:"grantTime,omitempty"`
	GrantStartTime     string `json:"grantStartTime,omitempty"`
	GrantEndTime       string `json:"grantEndTime,omitempty"`
	GrantStatus        string `json:"grantStatus,omitempty"`
}
