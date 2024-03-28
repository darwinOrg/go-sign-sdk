package test

import (
	"fmt"
	"github.com/darwinOrg/go-sign-sdk/fasc/client"
	"github.com/darwinOrg/go-sign-sdk/fasc/model/commonModel"
	"github.com/darwinOrg/go-sign-sdk/fasc/model/requestModel/docProcessRequestModel"
	"testing"
)

// 通过网络文件地址上传
func TestGetFddFileUploadUrl(t *testing.T) {
	c := client.NewClient(APPID, APPSECRET, SERVERURL)
	accessTokenRes := c.GetAuthToken()
	accessToken := accessTokenRes.Data.AccessToken

	req := &docProcessRequestModel.UploadFileByUrlReq{
		FileType: "doc",
		FileUrl:  "http://static.fadada.com/fortest20170901/协议书范本.pdf",
	}

	res := c.GetFddFileUploadUrl(req, accessToken)
	fmt.Println(ModelToJsonString(res, false))
}

// 获取本地文件上传url
func TestGetLocalFileUploadUrl(t *testing.T) {
	c := client.NewClient(APPID, APPSECRET, SERVERURL)
	accessTokenRes := c.GetAuthToken()
	accessToken := accessTokenRes.Data.AccessToken

	req := &docProcessRequestModel.GetLocalUploadFileUrlReq{
		FileType: "doc",
	}

	res := c.GetLocalFileUploadUrl(req, accessToken)
	fmt.Println(ModelToJsonString(res, false))
}

// 本地文件上传
func TestPutFileUpload(t *testing.T) {
	c := client.NewClient(APPID, APPSECRET, SERVERURL)
	er := c.PutFileUpload("", "")
	if er != nil {
		fmt.Println("文件put失败")
	}
}

// 文件处理
func TestFileProcess(t *testing.T) {
	c := client.NewClient(APPID, APPSECRET, SERVERURL)
	accessTokenRes := c.GetAuthToken()
	accessToken := accessTokenRes.Data.AccessToken

	req := &docProcessRequestModel.FileProcessReq{FddFileUrlList: []docProcessRequestModel.FddUploadFile{
		{FileType: "attach",
			FddFileUrl: "https://attach-test-os1.fadada.com/655c454fec864ad0bf9e52e3cb45d4ec",
			FileName:   "协议书范本.pdf",
		},
	}}

	res := c.FileProcess(req, accessToken)
	fmt.Println(ModelToJsonString(res, false))
}

// 文档验签
func TestVerifyFileSign(t *testing.T) {
	c := client.NewClient(APPID, APPSECRET, SERVERURL)
	accessTokenRes := c.GetAuthToken()
	accessToken := accessTokenRes.Data.AccessToken

	req := &docProcessRequestModel.VerifyFileSignReq{
		FileId:   DocFileId,
		FileHash: "fhaoifhiowhofowfohwfohowohfwhohwfo",
	}
	res := c.VerifyFileSign(req, accessToken)
	fmt.Println(ModelToJsonString(res, false))
}

// 获取文件对比页面链接
func TestGetCompareUrl(t *testing.T) {
	c := client.NewClient(APPID, APPSECRET, SERVERURL)
	accessTokenRes := c.GetAuthToken()
	accessToken := accessTokenRes.Data.AccessToken

	req := &docProcessRequestModel.GetCompareUrlReq{
		Initiator: &commonModel.OpenId{
			IdType: "corp",
			OpenId: OpenCorpId,
		},
		OriginFileId: "",
		TargetFileId: "",
	}

	res := c.GetCompareUrl(req, accessToken)
	fmt.Println(ModelToJsonString(res, false))
}

// 获取历史文件对比页面链接
func TestGetCompareResultUrl(t *testing.T) {
	c := client.NewClient(APPID, APPSECRET, SERVERURL)
	accessTokenRes := c.GetAuthToken()
	accessToken := accessTokenRes.Data.AccessToken

	req := &docProcessRequestModel.GetCompareResultUrlReq{
		Initiator: &commonModel.OpenId{
			IdType: "corp",
			OpenId: OpenCorpId,
		},
		CompareId: "",
	}

	res := c.GetCompareResultUrl(req, accessToken)
	fmt.Println(ModelToJsonString(res, false))
}

// 获取合同智审页面链接
func TestGetExamineUrl(t *testing.T) {
	c := client.NewClient(APPID, APPSECRET, SERVERURL)
	accessTokenRes := c.GetAuthToken()
	accessToken := accessTokenRes.Data.AccessToken

	req := &docProcessRequestModel.GetExamineUrlReq{
		Initiator: &commonModel.OpenId{
			IdType: "corp",
			OpenId: OpenCorpId,
		},
		FileId: "",
	}

	res := c.GetExamineUrl(req, accessToken)
	fmt.Println(ModelToJsonString(res, false))
}

// 获取历史合同智审页面链接
func TestGetExamineResultUrl(t *testing.T) {
	c := client.NewClient(APPID, APPSECRET, SERVERURL)
	accessTokenRes := c.GetAuthToken()
	accessToken := accessTokenRes.Data.AccessToken

	req := &docProcessRequestModel.GetExamineResultUrlReq{
		Initiator: &commonModel.OpenId{
			IdType: "corp",
			OpenId: OpenCorpId,
		},
		ExamineId: "",
	}

	res := c.GetExamineResultUrl(req, accessToken)
	fmt.Println(ModelToJsonString(res, false))
}
