package userManageResponseModel

import "github.com/darwinOrg/go-sign-sdk/fasc/model/commonModel"

type GetUserIdentifyInfoRes struct {
	RequestId string                     `json:"requestId"`
	Code      string                     `json:"code"`
	Msg       string                     `json:"msg"`
	Data      GetUserIdentifyInfoResData `json:"data"`
}

type GetUserIdentifyInfoResData struct {
	OpenUserId          string                      `json:"openUserId"`
	IdentStatus         string                      `json:"identStatus"`
	UserIdentInfo       *commonModel.UserIdentInfo  `json:"userIdentInfo"`
	UserIdentInfoExtend *commonModel.UserInfoExtend `json:"userIdentInfoExtend"`
	IdentMethod         string                      `json:"identMethod"`
	IdentSubmitTime     string                      `json:"identSubmitTime"`
	IdentSuccessTime    string                      `json:"identSuccessTime"`
	FddId               string                      `json:"fddId"`
}
