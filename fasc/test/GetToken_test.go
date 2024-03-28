package test

import (
	"encoding/json"
	"fmt"
	"github.com/darwinOrg/go-sign-sdk/fasc/client"
	"testing"
)

const (
	ContentTypeJson = "application/json"
	ContentTypeZip  = "application/zip"
	ContentTypePdf  = "application/pdf"
	ContentTypeFile = "application/*"
	AppId           = APPID
	AppSecret       = APPSECRET
	ServerUrl       = SERVERURL
)

// v5.1测试示例
// fasc_openapi_sdk 单元测试
func TestOpenApiClient_GetAuthToken(t *testing.T) {
	//获取服务访问凭证
	c := client.NewClient(AppId, AppSecret, ServerUrl)
	accessTokenRes := c.GetAuthToken()
	jsonData, err := json.Marshal(accessTokenRes)
	if err != nil {

	}
	fmt.Println(string(jsonData))
}
