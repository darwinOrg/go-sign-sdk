package client

import (
	"encoding/json"
	"fmt"
	"github.com/darwinOrg/go-sign-sdk/fasc/common"
	"github.com/darwinOrg/go-sign-sdk/fasc/model/requestModel/billManageRequestModel"
	"github.com/darwinOrg/go-sign-sdk/fasc/model/responseModel/billManageResponseModel"
)

const (
	GetBillUrlReqUrl = "/billing/get-bill-url" //获取计费链接
)

// GetBillUrlResponse 获取计费链接
func (o *OpenApiClient) GetBillUrl(billUrlReq *billManageRequestModel.GetBillingUrlReq, accessToken string) billManageResponseModel.GetBillingUrlRes {
	var billUrlRes billManageResponseModel.GetBillingUrlRes
	reqStr, err := json.Marshal(billUrlReq)
	if err != nil {
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken, string(reqStr))
	requestUrl := o.serverUrl + GetBillUrlReqUrl                       //接口请求api地址
	rspBody, requestId := common.SendPost(requestUrl, headMap, reqStr) //发送post请求
	err = json.Unmarshal(rspBody, &billUrlRes)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}
	billUrlRes.RequestId = requestId
	return billUrlRes
}
