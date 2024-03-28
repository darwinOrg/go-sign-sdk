package test

import (
	"fmt"
	"github.com/darwinOrg/go-sign-sdk/fasc/client"
	"github.com/darwinOrg/go-sign-sdk/fasc/model/commonModel"
	"github.com/darwinOrg/go-sign-sdk/fasc/model/requestModel/templateRequestModel"
	"testing"
)

// 查询文档模板列表
func TestGetDocTemplateList(t *testing.T) {
	c := client.NewClient(APPID, APPSECRET, SERVERURL)
	accessTokenRes := c.GetAuthToken()
	accessToken := accessTokenRes.Data.AccessToken

	req := &templateRequestModel.GetDocTemplateListReq{
		OwnerId: &commonModel.OpenId{
			IdType: "corp",
			OpenId: OpenCorpId,
		},
		ListFilter: &templateRequestModel.DocTemplateLisFilter{
			DocTemplateName: "",
		},
		ListPageNo:   1,
		ListPageSize: 100,
	}
	res := c.GetDocTemplateList(req, accessToken)
	fmt.Println(ModelToJsonString(res, false))
}

// 获取文档模板详情
func TestGetDocTemplateDetail(t *testing.T) {
	c := client.NewClient(APPID, APPSECRET, SERVERURL)
	accessTokenRes := c.GetAuthToken()
	accessToken := accessTokenRes.Data.AccessToken

	req := &templateRequestModel.GetDocTemplateDetailReq{
		OwnerId: &commonModel.OpenId{
			IdType: "corp",
			OpenId: OpenCorpId,
		},
		DocTemplateId: "",
	}
	res := c.GetDocTemplateDetail(req, accessToken)
	fmt.Println(ModelToJsonString(res, false))
}

// 获取签署模板列表
func TestGetSignTemplateList(t *testing.T) {
	c := client.NewClient(APPID, APPSECRET, SERVERURL)
	accessTokenRes := c.GetAuthToken()
	accessToken := accessTokenRes.Data.AccessToken

	req := &templateRequestModel.GetSignTemplateListReq{
		OwnerId: &commonModel.OpenId{
			IdType: "corp",
			OpenId: OpenCorpId,
		},
		ListFilter:   &templateRequestModel.SignTemplateListFilterInfo{SignTemplateName: ""},
		ListPageNo:   1,
		ListPageSize: 100,
	}
	res := c.GetSignTemplateList(req, accessToken)
	fmt.Println(ModelToJsonString(res, false))
}

// 获取签署模板详情
func TestGetSignTemplateDetail(t *testing.T) {
	c := client.NewClient(APPID, APPSECRET, SERVERURL)
	accessTokenRes := c.GetAuthToken()
	accessToken := accessTokenRes.Data.AccessToken

	req := &templateRequestModel.GetSignTemplateDetailReq{
		OwnerId: &commonModel.OpenId{
			IdType: "corp",
			OpenId: OpenCorpId,
		},
		SignTemplateId: SignTemplateId,
	}
	res := c.GetSignTemplateDetail(req, accessToken)
	fmt.Println(ModelToJsonString(res, false))
}

// 获取模板新增链接
func TestGetTemplateCreate(t *testing.T) {
	c := client.NewClient(APPID, APPSECRET, SERVERURL)
	accessTokenRes := c.GetAuthToken()
	accessToken := accessTokenRes.Data.AccessToken

	req := &templateRequestModel.GetTemplateCreateUrlReq{
		OpenCorpId:  OpenCorpId,
		Type:        "sign",
		RedirectUrl: "",
	}
	res := c.GetTemplateCreate(req, accessToken)
	fmt.Println(ModelToJsonString(res, false))
}

// 获取模板编辑链接
func TestGetTemplateEdit(t *testing.T) {
	c := client.NewClient(APPID, APPSECRET, SERVERURL)
	accessTokenRes := c.GetAuthToken()
	accessToken := accessTokenRes.Data.AccessToken

	req := &templateRequestModel.GetTemplateEditUrlReq{
		OpenCorpId:  OpenCorpId,
		TemplateId:  "",
		RedirectUrl: "",
	}
	res := c.GetTemplateEdit(req, accessToken)
	fmt.Println(ModelToJsonString(res, false))
}

// 获取模板预览链接
func TestGetTemplatePreview(t *testing.T) {
	c := client.NewClient(APPID, APPSECRET, SERVERURL)
	accessTokenRes := c.GetAuthToken()
	accessToken := accessTokenRes.Data.AccessToken

	req := &templateRequestModel.GetTemplatePreviewUrlReq{
		OpenCorpId:  OpenCorpId,
		TemplateId:  "",
		RedirectUrl: "",
	}
	res := c.GetTemplatePreview(req, accessToken)
	fmt.Println(ModelToJsonString(res, false))
}

// 获取模板管理链接
func TestGetTemplateManage(t *testing.T) {
	c := client.NewClient(APPID, APPSECRET, SERVERURL)
	accessTokenRes := c.GetAuthToken()
	accessToken := accessTokenRes.Data.AccessToken

	req := &templateRequestModel.GetTemplateManageUrlReq{
		OpenCorpId:  OpenCorpId,
		RedirectUrl: "",
	}
	res := c.GetTemplateManage(req, accessToken)
	fmt.Println(ModelToJsonString(res, false))
}

// 查询应用文档模板列表
func TestGetAppDocTemplates(t *testing.T) {
	c := client.NewClient(APPID, APPSECRET, SERVERURL)
	accessTokenRes := c.GetAuthToken()
	accessToken := accessTokenRes.Data.AccessToken

	req := &templateRequestModel.GetAppDocTemplateListReq{
		&templateRequestModel.DocListFilterInfo{AppDocTemplateName: ""},
	}
	res := c.GetAppDocTemplates(req, accessToken)
	fmt.Println(ModelToJsonString(res, false))
}

// 获取应用文档模板详情
func TestGetAppDocTemplateDetail(t *testing.T) {
	c := client.NewClient(APPID, APPSECRET, SERVERURL)
	accessTokenRes := c.GetAuthToken()
	accessToken := accessTokenRes.Data.AccessToken

	req := &templateRequestModel.GetAppDocTemplateDetailReq{AppDocTemplateId: ""}
	res := c.GetAppDocTemplateDetail(req, accessToken)
	fmt.Println(ModelToJsonString(res, false))
}

// 启用停用应用文档模板
func TestSetAppDocTemplateStatus(t *testing.T) {
	c := client.NewClient(APPID, APPSECRET, SERVERURL)
	accessTokenRes := c.GetAuthToken()
	accessToken := accessTokenRes.Data.AccessToken

	req := &templateRequestModel.SetAppDocTemplateStatusReq{
		AppDocTemplateId:     AppDocTemplateId,
		AppDocTemplateStatus: "valid",
	}
	res := c.SetAppDocTemplateStatus(req, accessToken)
	fmt.Println(ModelToJsonString(res, false))
}

// 删除应用文档模板
func TestDeleteAppDocTemplate(t *testing.T) {
	c := client.NewClient(APPID, APPSECRET, SERVERURL)
	accessTokenRes := c.GetAuthToken()
	accessToken := accessTokenRes.Data.AccessToken

	req := &templateRequestModel.DeleteAppDocTemplateReq{AppDocTemplateId: AppSignTemplateId}
	res := c.DeleteAppDocTemplate(req, accessToken)
	fmt.Println(ModelToJsonString(res, false))
}

// 查询应用签署任务模板列表
func TestGetAppSignTemplates(t *testing.T) {
	c := client.NewClient(APPID, APPSECRET, SERVERURL)
	accessTokenRes := c.GetAuthToken()
	accessToken := accessTokenRes.Data.AccessToken

	req := &templateRequestModel.GetAppSignTemplateListReq{
		ListFilter: &templateRequestModel.SignListFilterInfo{AppSignTemplateName: ""},
	}
	res := c.GetAppSignTemplates(req, accessToken)
	fmt.Println(ModelToJsonString(res, false))
}

// 获取应用签署任务模板详情
func TestGetAppSignTemplateDetail(t *testing.T) {
	c := client.NewClient(APPID, APPSECRET, SERVERURL)
	accessTokenRes := c.GetAuthToken()
	accessToken := accessTokenRes.Data.AccessToken

	req := &templateRequestModel.GetAppSignTemplateDetailReq{AppSignTemplateId: ""}
	res := c.GetAppSignTemplateDetail(req, accessToken)
	fmt.Println(ModelToJsonString(res, false))
}

// 启用停用应用签署模板
func TestSetAppSignTaskTemplateStatus(t *testing.T) {
	c := client.NewClient(APPID, APPSECRET, SERVERURL)
	accessTokenRes := c.GetAuthToken()
	accessToken := accessTokenRes.Data.AccessToken

	req := &templateRequestModel.SetAppSignTemplateStatusReq{
		AppSignTemplateId:     AppSignTemplateId,
		AppSignTemplateStatus: "invalid",
	}
	res := c.SetAppSignTemplateStatus(req, accessToken)
	fmt.Println(ModelToJsonString(res, false))
}

// 删除应用签署模板
func TestDeleteAppSignTaskTemplate(t *testing.T) {
	c := client.NewClient(APPID, APPSECRET, SERVERURL)
	accessTokenRes := c.GetAuthToken()
	accessToken := accessTokenRes.Data.AccessToken

	req := &templateRequestModel.DeleteAppSignTemplateReq{AppSignTemplateId: AppSignTemplateId}
	res := c.DeleteAppSignTemplate(req, accessToken)
	fmt.Println(ModelToJsonString(res, false))
}

// 获取应用模板新增链接
func TestGetAppTemplateCreate(t *testing.T) {
	c := client.NewClient(APPID, APPSECRET, SERVERURL)
	accessTokenRes := c.GetAuthToken()
	accessToken := accessTokenRes.Data.AccessToken

	req := &templateRequestModel.GetAppTemplateCreateUrlReq{
		Type:           "sign",
		CreateSerialNo: "",
		RedirectUrl:    "",
	}
	res := c.GetAppTemplateCreate(req, accessToken)
	fmt.Println(ModelToJsonString(res, false))
}

// 获取应用模板编辑链接
func TestGetAppTemplateEdit(t *testing.T) {
	c := client.NewClient(APPID, APPSECRET, SERVERURL)
	accessTokenRes := c.GetAuthToken()
	accessToken := accessTokenRes.Data.AccessToken

	req := &templateRequestModel.GetAppTemplateEditUrlReq{
		AppTemplateId: "",
		RedirectUrl:   "",
	}
	res := c.GetAppTemplateEdit(req, accessToken)
	fmt.Println(ModelToJsonString(res, false))
}

// 获取应用模板预览链接
func TestGetAppTemplatePreview(t *testing.T) {
	c := client.NewClient(APPID, APPSECRET, SERVERURL)
	accessTokenRes := c.GetAuthToken()
	accessToken := accessTokenRes.Data.AccessToken

	req := &templateRequestModel.GetAppTemplatePreviewUrlReq{
		AppTemplateId: "",
		RedirectUrl:   "",
	}
	res := c.GetAppTemplatePreview(req, accessToken)
	fmt.Println(ModelToJsonString(res, false))
}

// 创建业务控件
func TestCreateAppField(t *testing.T) {
	c := client.NewClient(APPID, APPSECRET, SERVERURL)
	accessTokenRes := c.GetAuthToken()
	accessToken := accessTokenRes.Data.AccessToken

	req := &templateRequestModel.CreateAppFieldReq{
		FieldKey:  "fieldKey001",
		FieldName: "fieldKey001",
		FieldType: "text_single_line",
	}
	res := c.CreateAppField(req, accessToken)
	fmt.Println(ModelToJsonString(res, false))
}

// 修改业务控件
func TestModifyAppField(t *testing.T) {
	c := client.NewClient(APPID, APPSECRET, SERVERURL)
	accessTokenRes := c.GetAuthToken()
	accessToken := accessTokenRes.Data.AccessToken

	req := &templateRequestModel.ModifyAppFieldReq{
		FieldKey:  "fieldKey001",
		FieldName: "fieldKey001名称",
	}
	res := c.ModifyAppField(req, accessToken)
	fmt.Println(ModelToJsonString(res, false))
}

// 设置业务控件状态
func TestSetAppFieldStatus(t *testing.T) {
	c := client.NewClient(APPID, APPSECRET, SERVERURL)
	accessTokenRes := c.GetAuthToken()
	accessToken := accessTokenRes.Data.AccessToken

	req := &templateRequestModel.SetAppFieldStatusReq{
		FieldKey:    "fieldKey001",
		FieldStatus: "enable",
	}
	res := c.SetAppFieldStatus(req, accessToken)
	fmt.Println(ModelToJsonString(res, false))
}

// 查询业务控件列表
func TestGetAppFieldList(t *testing.T) {
	c := client.NewClient(APPID, APPSECRET, SERVERURL)
	accessTokenRes := c.GetAuthToken()
	accessToken := accessTokenRes.Data.AccessToken

	req := &templateRequestModel.GetAppFieldListReq{}
	res := c.GetAppFieldList(req, accessToken)
	fmt.Println(ModelToJsonString(res, false))
}
