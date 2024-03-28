package client

import (
	"encoding/json"
	"fmt"
	common "github.com/darwinOrg/go-sign-sdk/fasc/common"
	"github.com/darwinOrg/go-sign-sdk/fasc/model/requestModel/euiRequestModel"
	"github.com/darwinOrg/go-sign-sdk/fasc/model/responseModel/euiResponseModel"
)

const (
	AppPageResourceReqUrl  = "/app-page-resource/get-url"  //获取应用级资源访问链接
	UserPageResourceReqUrl = "/user-page-resource/get-url" //获取用户级资源访问链接
)

// GetAppResourceResponse 获取应用级资源访问链接
func (o *OpenApiClient) GetAppResourceResponse(appResourceReq *euiRequestModel.GetAppResourceUrlReq, accessToken string) euiResponseModel.GetAppResourceUrlRes {
	var appResourceRes euiResponseModel.GetAppResourceUrlRes
	reqStr, err := json.Marshal(appResourceReq)
	if err != nil {
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken, string(reqStr))
	requestUrl := o.serverUrl + AppPageResourceReqUrl                  //接口请求api地址
	rspBody, requestId := common.SendPost(requestUrl, headMap, reqStr) //发送post请求
	err = json.Unmarshal(rspBody, &appResourceRes)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}
	appResourceRes.RequestId = requestId
	return appResourceRes
}

// GetUserResourceResponse 获取用户级资源访问链接
func (o *OpenApiClient) GetUserResourceResponse(userResourceReq *euiRequestModel.GetUserResourceUrlReq, accessToken string) euiResponseModel.GetUserResourceUrlRes {
	var userResourceRes euiResponseModel.GetUserResourceUrlRes
	reqStr, err := json.Marshal(userResourceReq)
	if err != nil {
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken, string(reqStr))
	requestUrl := o.serverUrl + UserPageResourceReqUrl                 //接口请求api地址
	rspBody, requestId := common.SendPost(requestUrl, headMap, reqStr) //发送post请求
	err = json.Unmarshal(rspBody, &userResourceRes)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}
	userResourceRes.RequestId = requestId
	return userResourceRes
}
