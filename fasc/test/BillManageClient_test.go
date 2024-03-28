package test

import (
	"fmt"
	"github.com/darwinOrg/go-sign-sdk/fasc/client"
	"github.com/darwinOrg/go-sign-sdk/fasc/model/commonModel"
	"github.com/darwinOrg/go-sign-sdk/fasc/model/requestModel/billManageRequestModel"
	"testing"
)

// 查询企业用印员列表
func TestGetBillUrl(t *testing.T) {
	c := client.NewClient(APPID, APPSECRET, SERVERURL)
	accessTokenRes := c.GetAuthToken()
	accessToken := accessTokenRes.Data.AccessToken

	req := &billManageRequestModel.GetBillingUrlReq{
		OpenId: &commonModel.OpenId{
			IdType: "corp",
			OpenId: OpenCorpId,
		},
		UrlType:     "",
		RedirectUrl: "",
	}

	res := c.GetBillUrl(req, accessToken)
	fmt.Println(ModelToJsonString(res, false))
}
