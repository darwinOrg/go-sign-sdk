package corpManageResponseModel

import "github.com/darwinOrg/go-sign-sdk/fasc/model/commonModel"

type GetCorpIdentityInfoRes struct {
	RequestId string                     `json:"requestId"`
	Code      string                     `json:"code"`
	Msg       string                     `json:"msg"`
	Data      GetCorpIdentityInfoResData `json:"data"`
}

type GetCorpIdentityInfoResData struct {
	OpenCorpId          string                      `json:"openCorpId,omitempty"`
	CorpIdentStatus     string                      `json:"corpIdentStatus,omitempty"`
	CorpIdentInfo       *commonModel.CorpIdentInfo  `json:"corpIdentInfo,omitempty"`
	CorpIdentInfoExtend *commonModel.CorpInfoExtend `json:"corpIdentInfoExtend,omitempty"`
	CorpIdentMethod     string                      `json:"corpIdentMethod,omitempty"`
	IdentSubmitTime     string                      `json:"identSubmitTime,omitempty"`
	IdentSuccessTime    string                      `json:"identSuccessTime,omitempty"`
	FddId               string                      `json:"fddId,omitempty"`
}
