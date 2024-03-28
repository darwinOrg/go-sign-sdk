package test

import (
	"fmt"
	"github.com/darwinOrg/go-sign-sdk/fasc/client"
	"github.com/darwinOrg/go-sign-sdk/fasc/model/commonModel"
	"github.com/darwinOrg/go-sign-sdk/fasc/model/requestModel/euiRequestModel"
	"testing"
)

// 获取应用级资源访问链接
func TestGetAppResourceResponse(t *testing.T) {
	c := client.NewClient(APPID, APPSECRET, SERVERURL)
	accessTokenRes := c.GetAuthToken()
	accessToken := accessTokenRes.Data.AccessToken
	req := &euiRequestModel.GetAppResourceUrlReq{
		OpenId: &commonModel.OpenId{
			IdType: "corp",
			OpenId: OpenCorpId,
		},
		Resource: &euiRequestModel.AppResource{
			ResourceId: "SIGNTASK",
			Action:     "EDIT",
			Params:     "signTaskId=15379579395",
		},
	}
	response := c.GetAppResourceResponse(req, accessToken)
	fmt.Println(ModelToJsonString(response, false))
}

// 获取用户级资源访问链接
func TestGetUserResourceResponse(t *testing.T) {
	c := client.NewClient(APPID, APPSECRET, SERVERURL)
	accessTokenRes := c.GetAuthToken()
	accessToken := accessTokenRes.Data.AccessToken
	req := &euiRequestModel.GetUserResourceUrlReq{
		OpenCorpId:   OpenCorpId,
		ClientUserId: "",
		Resource: &euiRequestModel.UserResource{
			ResourceId: "SEAL",
			Action:     "MANAGE",
			Params:     "",
		},
	}
	response := c.GetUserResourceResponse(req, accessToken)
	fmt.Println(ModelToJsonString(response, false))
}
