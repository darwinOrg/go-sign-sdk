package test

import (
	"fmt"
	"github.com/darwinOrg/go-sign-sdk/fasc/client"
	"github.com/darwinOrg/go-sign-sdk/fasc/model/commonModel"
	"github.com/darwinOrg/go-sign-sdk/fasc/model/requestModel/sealRequestModel"
	"strconv"
	"testing"
	"time"
)

// 查询印章列表
func TestGetSealList(t *testing.T) {
	c := client.NewClient(APPID, APPSECRET, SERVERURL)
	accessTokenRes := c.GetAuthToken()
	accessToken := accessTokenRes.Data.AccessToken

	req := &sealRequestModel.GetSealListReq{
		OpenCorpId: OpenCorpId,
		ListFilter: &sealRequestModel.SealListFilter{CategoryType: []string{}},
	}
	res := c.GetSealList(req, accessToken)
	fmt.Println(ModelToJsonString(res, false))
}

// 查询印章详情
func TestGetSealDetail(t *testing.T) {
	c := client.NewClient(APPID, APPSECRET, SERVERURL)
	accessTokenRes := c.GetAuthToken()
	accessToken := accessTokenRes.Data.AccessToken

	req := &sealRequestModel.GetSealDetailReq{
		OpenCorpId: OpenCorpId,
		SealId:     0,
	}

	res := c.GetSealDetail(req, accessToken)
	fmt.Println(ModelToJsonString(res, false))
}

// 获取指定印章详情链接
func TestGetAppointedSeal(t *testing.T) {
	c := client.NewClient(APPID, APPSECRET, SERVERURL)
	accessTokenRes := c.GetAuthToken()
	accessToken := accessTokenRes.Data.AccessToken

	req := &sealRequestModel.GetAppointedSealUrlReq{
		OpenCorpId:   OpenCorpId,
		SealId:       0,
		ClientUserId: "clientUserId" + strconv.FormatInt(time.Now().Unix(), 10),
		RedirectUrl:  "http://www.baidu.com",
	}

	res := c.GetAppointedSeal(req, accessToken)
	fmt.Println(ModelToJsonString(res, false))
}

// 查询企业用印员列表
func TestGetSealUserList(t *testing.T) {
	c := client.NewClient(APPID, APPSECRET, SERVERURL)
	accessTokenRes := c.GetAuthToken()
	accessToken := accessTokenRes.Data.AccessToken

	req := &sealRequestModel.GetSealUserListReq{
		OpenCorpId: OpenCorpId,
		ListFilter: &sealRequestModel.SealListFilter{[]string{
			"official_seal",
		}},
	}

	res := c.GetSealUserList(req, accessToken)
	fmt.Println(ModelToJsonString(res, false))
}

// 查询指定成员的印章列表
func TestGetUserSealList(t *testing.T) {
	c := client.NewClient(APPID, APPSECRET, SERVERURL)
	accessTokenRes := c.GetAuthToken()
	accessToken := accessTokenRes.Data.AccessToken

	req := &sealRequestModel.GetUserSealListReq{
		OpenCorpId: OpenCorpId,
		MemberId:   0,
	}

	res := c.GetUserSealList(req, accessToken)
	fmt.Println(ModelToJsonString(res, false))
}

// 获取印章创建链接
func TestGetSealCreate(t *testing.T) {
	c := client.NewClient(APPID, APPSECRET, SERVERURL)
	accessTokenRes := c.GetAuthToken()
	accessToken := accessTokenRes.Data.AccessToken

	req := &sealRequestModel.GetSealCreateUrlReq{
		OpenCorpId:     OpenCorpId,
		SealName:       "",
		CategoryType:   "",
		SealTag:        "",
		CreateSerialNo: "",
		ClientUserId:   "clientUserId" + strconv.FormatInt(time.Now().Unix(), 10),
		RedirectUrl:    "",
	}

	res := c.GetSealCreate(req, accessToken)
	fmt.Println(ModelToJsonString(res, false))
}

// 查询审核中的印章列表
func TestGetVerifySealList(t *testing.T) {
	c := client.NewClient(APPID, APPSECRET, SERVERURL)
	accessTokenRes := c.GetAuthToken()
	accessToken := accessTokenRes.Data.AccessToken

	req := &sealRequestModel.GetSealVerifyListReq{
		OpenCorpId: OpenCorpId,
	}

	res := c.GetVerifySealList(req, accessToken)
	fmt.Println(ModelToJsonString(res, false))
}

// 修改印章基本信息
func TestModifySealInfo(t *testing.T) {
	c := client.NewClient(APPID, APPSECRET, SERVERURL)
	accessTokenRes := c.GetAuthToken()
	accessToken := accessTokenRes.Data.AccessToken

	req := &sealRequestModel.ModifySealReq{
		OpenCorpId:   OpenCorpId,
		SealId:       0,
		SealName:     "修改后的印章名称",
		SealTag:      "",
		CategoryType: "",
	}

	res := c.ModifySealInfo(req, accessToken)
	fmt.Println(ModelToJsonString(res, false))
}

// 获取设置用印员链接
func TestGetSealGrant(t *testing.T) {
	c := client.NewClient(APPID, APPSECRET, SERVERURL)
	accessTokenRes := c.GetAuthToken()
	accessToken := accessTokenRes.Data.AccessToken

	req := &sealRequestModel.GetSealGrantUrlReq{
		OpenCorpId: OpenCorpId,
		SealId:     0,
		MemberInfo: &commonModel.MemberInfo{
			MemberIds:      []int64{},
			GrantStartTime: "",
			GrantEndTime:   "",
		},
		ClientUserId: "clientUserId" + strconv.FormatInt(time.Now().Unix(), 10),
		RedirectUrl:  "http://www.baidu.com",
	}

	res := c.GetSealGrant(req, accessToken)
	fmt.Println(ModelToJsonString(res, false))
}

// 获取设置企业印章免验证签链接
func TestGetSealFreeSign(t *testing.T) {
	c := client.NewClient(APPID, APPSECRET, SERVERURL)
	accessTokenRes := c.GetAuthToken()
	accessToken := accessTokenRes.Data.AccessToken

	req := &sealRequestModel.GetSealFreeSignUrlReq{
		OpenCorpId:   OpenCorpId,
		SealIds:      []int64{},
		BusinessId:   "",
		Email:        "",
		ClientUserId: "clientUserId" + strconv.FormatInt(time.Now().Unix(), 10),
		RedirectUrl:  "",
	}

	res := c.GetSealFreeSign(req, accessToken)
	fmt.Println(ModelToJsonString(res, false))
}

// 解除印章授权
func TestCancelSealGrant(t *testing.T) {
	c := client.NewClient(APPID, APPSECRET, SERVERURL)
	accessTokenRes := c.GetAuthToken()
	accessToken := accessTokenRes.Data.AccessToken

	req := &sealRequestModel.CancelSealGrantReq{
		OpenCorpId: OpenCorpId,
		SealId:     0,
		MemberId:   0,
	}

	res := c.CancelSealGrant(req, accessToken)
	fmt.Println(ModelToJsonString(res, false))
}

// 设置印章状态
func TestSetSealStatus(t *testing.T) {
	c := client.NewClient(APPID, APPSECRET, SERVERURL)
	accessTokenRes := c.GetAuthToken()
	accessToken := accessTokenRes.Data.AccessToken

	req := &sealRequestModel.SetSealStatusReq{
		OpenCorpId: OpenCorpId,
		SealId:     0,
		SealStatus: "disable",
	}

	res := c.SetSealStatus(req, accessToken)
	fmt.Println(ModelToJsonString(res, false))
}

// 删除印章
func TestDeleteSeal(t *testing.T) {
	c := client.NewClient(APPID, APPSECRET, SERVERURL)
	accessTokenRes := c.GetAuthToken()
	accessToken := accessTokenRes.Data.AccessToken

	req := &sealRequestModel.SealDeleteReq{
		OpenCorpId: OpenCorpId,
		SealId:     0,
	}

	res := c.DeleteSeal(req, accessToken)
	fmt.Println(ModelToJsonString(res, false))
}

// 获取印章管理链接
func TestGetSealManage(t *testing.T) {
	c := client.NewClient(APPID, APPSECRET, SERVERURL)
	accessTokenRes := c.GetAuthToken()
	accessToken := accessTokenRes.Data.AccessToken

	req := &sealRequestModel.GetSealManageUrlReq{
		OpenCorpId:   OpenCorpId,
		ClientUserId: "clientUserId" + strconv.FormatInt(time.Now().Unix(), 10),
		RedirectUrl:  "",
	}

	res := c.GetSealManage(req, accessToken)
	fmt.Println(ModelToJsonString(res, false))
}

// 查询个人签名列表
func TestGetPersonalSealList(t *testing.T) {
	c := client.NewClient(APPID, APPSECRET, SERVERURL)
	accessTokenRes := c.GetAuthToken()
	accessToken := accessTokenRes.Data.AccessToken

	req := &sealRequestModel.GetPersonalSealListReq{
		OpenUserId: OpenUserId,
	}

	res := c.GetPersonalSealList(req, accessToken)
	fmt.Println(ModelToJsonString(res, false))
}

// 获取设置个人签名免验证签链接
func TestGetPersonalSealFreeSign(t *testing.T) {
	c := client.NewClient(APPID, APPSECRET, SERVERURL)
	accessTokenRes := c.GetAuthToken()
	accessToken := accessTokenRes.Data.AccessToken

	req := &sealRequestModel.GetPersonalSealFreeSignUrlReq{
		OpenUserId:  OpenUserId,
		SealIds:     []int64{},
		BusinessId:  "",
		Email:       "",
		RedirectUrl: "",
	}

	res := c.GetPersonalSealFreeSign(req, accessToken)
	fmt.Println(ModelToJsonString(res, false))
}
