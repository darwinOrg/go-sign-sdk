package test

import (
	"encoding/json"
	"fmt"
	"github.com/darwinOrg/go-sign-sdk/fasc/client"
	"github.com/darwinOrg/go-sign-sdk/fasc/model/commonModel"
	"github.com/darwinOrg/go-sign-sdk/fasc/model/requestModel/userManageRequestModel"
	"strconv"
	"testing"
	"time"
)

func TestGetUserAuthUrl(t *testing.T) {
	c := client.NewClient(APPID, APPSECRET, SERVERURL)
	accessTokenRes := c.GetAuthToken()
	accessToken := accessTokenRes.Data.AccessToken
	req := &userManageRequestModel.GetUserAuthUrlReq{
		ClientUserId:       ClientUserId + strconv.FormatInt(time.Now().Unix(), 10),
		UserName:           "",
		UserIdentType:      "",
		UserIdentNo:        "",
		UserIdentInfoMatch: false,
		AuthScopes:         []string{"ident_info", "signtask_info", "seal_info", "signtask_init", "signtask_file"},
		RedirectUrl:        "http://www.baidu.com",
		AccountName:        "",
		UserIdentInfo: &userManageRequestModel.UserIdentInfo{
			UserName:      "",
			UserIdentType: "",
			UserIdentNo:   "",
			Mobile:        "",
			BankAccountNo: "",
			IdentMethod:   []string{},
		},
		NonEditableInfo: []string{},
	}
	res := c.GetUserAuthUrlResponse(req, accessToken)
	jsonData, _ := json.Marshal(res)
	fmt.Println(string(jsonData))
}

// 禁用个人用户
func TestGetUserDisableResponse(t *testing.T) {
	c := client.NewClient(APPID, APPSECRET, SERVERURL)
	accessTokenRes := c.GetAuthToken()
	accessToken := accessTokenRes.Data.AccessToken
	req := &userManageRequestModel.UserDisableReq{
		OpenUserId: OpenUserId,
	}
	response := c.GetUserDisableResponse(req, accessToken)
	jsonData, _ := json.Marshal(response)
	fmt.Println(string(jsonData))
}

// 回复个人用户
func TestGetUserEnableResponse(t *testing.T) {
	c := client.NewClient(APPID, APPSECRET, SERVERURL)
	accessTokenRes := c.GetAuthToken()
	accessToken := accessTokenRes.Data.AccessToken
	req := &userManageRequestModel.UserEnableReq{
		OpenUserId: OpenUserId,
	}
	response := c.GetUserEnableResponse(req, accessToken)
	jsonData, _ := json.Marshal(response)
	fmt.Println(string(jsonData))
}

// 解绑个人用户
func TestGetUnBindUserResponse(t *testing.T) {
	c := client.NewClient(APPID, APPSECRET, SERVERURL)
	accessTokenRes := c.GetAuthToken()
	accessToken := accessTokenRes.Data.AccessToken
	req := &userManageRequestModel.UnBindUserReq{
		OpenUserId: OpenUserId,
	}
	response := c.GetUnBindUserResponse(req, accessToken)
	jsonData, _ := json.Marshal(response)
	fmt.Println(string(jsonData))
}

// 查询个人用户基本信息
func TestGetUserInfoResponse(t *testing.T) {
	c := client.NewClient(APPID, APPSECRET, SERVERURL)
	accessTokenRes := c.GetAuthToken()
	accessToken := accessTokenRes.Data.AccessToken
	req := &userManageRequestModel.GetUserReq{
		OpenUserId:   OpenUserId,
		ClientUserId: "",
	}
	response := c.GetUserInfoResponse(req, accessToken)
	jsonData, _ := json.Marshal(response)
	fmt.Println(string(jsonData))
}

// 获取个人用户身份信息
func TestGetUserIdentityResponse(t *testing.T) {
	c := client.NewClient(APPID, APPSECRET, SERVERURL)
	accessTokenRes := c.GetAuthToken()
	accessToken := accessTokenRes.Data.AccessToken
	req := &userManageRequestModel.GetUserIdentifyInfoReq{
		OpenUserId: OpenUserId,
	}
	response := c.GetUserIdentityResponse(req, accessToken)
	jsonResData, _ := json.Marshal(response)
	fmt.Println(string(jsonResData))
}

// 个人三要素
func TestGetThreeElementVerifyUrl(t *testing.T) {
	c := client.NewClient(APPID, APPSECRET, SERVERURL)
	accessTokenRes := c.GetAuthToken()
	accessToken := accessTokenRes.Data.AccessToken
	req := &userManageRequestModel.GetThreeElementVerifyUrlReq{
		ClientUserId: "ClientUserId34304924",
		Initiator: &commonModel.OpenId{
			IdType: "corp",
			OpenId: OpenCorpId,
		},
		UserName:    "周福成",
		UserIdentNo: "440981199904038619",
		Mobile:      "13138799005",
		RedirectUrl: "http://www.baidu.com",
	}
	response := c.GetThreeElementVerifyUrl(req, accessToken)
	fmt.Println(ModelToJsonString(response, false))
}

// 个人四要素
func TestGetFourElementVerifyUrl(t *testing.T) {
	c := client.NewClient(APPID, APPSECRET, SERVERURL)
	accessTokenRes := c.GetAuthToken()
	accessToken := accessTokenRes.Data.AccessToken
	req := &userManageRequestModel.GetFourElementVerifyUrlReq{
		ClientUserId: "ClientUserId34304924",
		Initiator: &commonModel.OpenId{
			IdType: "corp",
			OpenId: OpenCorpId,
		},
		UserName:      "周福成",
		UserIdentNo:   "440981199904038619",
		IdCardImage:   true,
		BankAccountNo: "",
		Mobile:        "13138799005",
		RedirectUrl:   "http://www.baidu.com",
	}
	response := c.GetFourElementVerifyUrl(req, accessToken)
	fmt.Println(ModelToJsonString(response, false))
}

// 三四要素下载图片
func TestGetIdCardImageDownloadUrl(t *testing.T) {
	c := client.NewClient(APPID, APPSECRET, SERVERURL)
	accessTokenRes := c.GetAuthToken()
	accessToken := accessTokenRes.Data.AccessToken
	req := &userManageRequestModel.GetIdCardImageDownloadUrlReq{VerifyId: "62645ca9defc44b4ac91ce6f6de05f4c"}
	response := c.GetIdCardImageDownloadUrl(req, accessToken)
	fmt.Println(ModelToJsonString(response, false))
}
