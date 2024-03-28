package sealResponseModel

type GetSealVerifyListRes struct {
	RequestId string            `json:"requestId"`
	Code      string            `json:"code"`
	Msg       string            `json:"msg"`
	Data      GetSealVerifyList `json:"data"`
}

type GetSealVerifyList struct {
	VerifyInfos []VerifyInfo `json:"verifyInfos,omitempty"`
}

type VerifyInfo struct {
	VerifyId     int64  `json:"verifyId,string"`
	SealName     string `json:"sealName,omitempty"`
	CategoryType string `json:"categoryType,omitempty"`
	PicFileId    string `json:"picFileId,omitempty"`
	PicFileUrl   string `json:"picFileUrl,omitempty"`
	CreateTime   string `json:"createTime,omitempty"`
	VerifyStatus string `json:"verifyStatus,omitempty"`
}
