package test

import (
	"encoding/json"
	"fmt"
	"github.com/darwinOrg/go-sign-sdk/fasc/client"
	"github.com/darwinOrg/go-sign-sdk/fasc/model/requestModel/corpManageRequestModel"
	"strconv"
	"testing"
	"time"
)

// 获取企业授权链接
func TestGetCorpAuthUrlResponse(t *testing.T) {
	c := client.NewClient(APPID, APPSECRET, SERVERURL)
	accessTokenRes := c.GetAuthToken()
	accessToken := accessTokenRes.Data.AccessToken
	req := &corpManageRequestModel.GetCorpAuthUrlReq{
		ClientCorpId: "clientCorpId" + strconv.FormatInt(time.Now().Unix(), 10),
		ClientUserId: "clientUserId" + strconv.FormatInt(time.Now().Unix(), 10),
		AccountName:  "13138799005",
		CorpIdentInfo: &corpManageRequestModel.CorpIdentInfoReq{
			CorpName:        "",
			CorpIdentType:   "",
			CorpIdentNo:     "",
			LegalRepName:    "",
			License:         "",
			CorpIdentMethod: []string{"legalRep"},
		},
		CorpNonEditableInfo: nil,
		OprIdentInfo: &corpManageRequestModel.OprIdentInfoReq{
			UserName:       "",
			UserIdentType:  "",
			UserIdentNo:    "",
			Mobile:         "",
			BankAccountNo:  "",
			OprIdentMethod: []string{"mobile"},
		},
		OprNonEditableInfo: nil,
		CorpIdentInfoMatch: false,
		AuthScopes: []string{
			"ident_info", "seal_info", "signtask_info", "signtask_init", "organization", "template", "signtask_file",
		},
		RedirectUrl: "http://www.baidu.com",
	}
	response := c.GetCorpAuthUrlResponse(req, accessToken)
	jsonResData, _ := json.Marshal(response)
	fmt.Println(string(jsonResData))
}

// 禁用企业
func TestGetCorpDisableResponse(t *testing.T) {
	c := client.NewClient(APPID, APPSECRET, SERVERURL)
	accessTokenRes := c.GetAuthToken()
	accessToken := accessTokenRes.Data.AccessToken
	req := &corpManageRequestModel.CorpDisableReq{
		OpenCorpId: OpenCorpId,
	}
	response := c.GetCorpDisableResponse(req, accessToken)
	jsonResData, _ := json.Marshal(response)
	fmt.Println(string(jsonResData))
}

// 恢复企业
func TestGetCorpEnableResponse(t *testing.T) {
	c := client.NewClient(APPID, APPSECRET, SERVERURL)
	accessTokenRes := c.GetAuthToken()
	accessToken := accessTokenRes.Data.AccessToken
	req := &corpManageRequestModel.CorpEnableReq{
		OpenCorpId: OpenCorpId,
	}
	response := c.GetCorpEnableResponse(req, accessToken)
	jsonResData, _ := json.Marshal(response)
	fmt.Println(string(jsonResData))
}

// 解绑企业用户
func TestGetCorpUnBindResponse(t *testing.T) {
	c := client.NewClient(APPID, APPSECRET, SERVERURL)
	accessTokenRes := c.GetAuthToken()
	accessToken := accessTokenRes.Data.AccessToken
	req := &corpManageRequestModel.UnBindCorpReq{
		OpenCorpId: OpenCorpId,
	}
	response := c.GetCorpUnBindResponse(req, accessToken)
	jsonResData, _ := json.Marshal(response)
	fmt.Println(string(jsonResData))
}

// 查询企业用户基本信息
func TestGetCorpInfoResponse(t *testing.T) {
	c := client.NewClient(APPID, APPSECRET, SERVERURL)
	accessTokenRes := c.GetAuthToken()
	accessToken := accessTokenRes.Data.AccessToken
	req := &corpManageRequestModel.GetCorpReq{
		OpenCorpId:   OpenCorpId,
		ClientCorpId: "",
	}
	response := c.GetCorpInfoResponse(req, accessToken)
	jsonResData, _ := json.Marshal(response)
	fmt.Println(string(jsonResData))
}

// 获取企业用户身份信息
func TestGetCorpIdentityResponse(t *testing.T) {
	c := client.NewClient(APPID, APPSECRET, SERVERURL)
	accessTokenRes := c.GetAuthToken()
	accessToken := accessTokenRes.Data.AccessToken
	req := &corpManageRequestModel.GetCorpIdentityInfoReq{
		OpenCorpId: OpenCorpId,
	}
	response := c.GetCorpIdentityResponse(req, accessToken)
	jsonResData, _ := json.Marshal(response)
	fmt.Println(string(jsonResData))
}

// 查询企业实名认证状态
func TestGetCorpIdentifiedStatus(t *testing.T) {
	c := client.NewClient(APPID, APPSECRET, SERVERURL)
	accessTokenRes := c.GetAuthToken()
	accessToken := accessTokenRes.Data.AccessToken
	req := &corpManageRequestModel.GetIdentifiedStatusReq{
		CorpName:    "",
		CorpIdentNo: "",
	}
	response := c.GetCorpIdentifiedStatus(req, accessToken)
	jsonResData, _ := json.Marshal(response)
	fmt.Println(string(jsonResData))
}
