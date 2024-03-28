package client

import (
	"encoding/json"
	"fmt"
	"github.com/darwinOrg/go-sign-sdk/fasc/common"
	"github.com/darwinOrg/go-sign-sdk/fasc/model/requestModel/corpManageRequestModel"
	"github.com/darwinOrg/go-sign-sdk/fasc/model/responseModel/corpManageResponseModel"
)

const (
	GetCorpAuthReqUrl          = "/corp/get-auth-url"          //获取企业用户授权链接
	CorpDisableReqUrl          = "/corp/disable"               //禁用企业用户
	CorpEnableReqUrl           = "/corp/enable"                //恢复企业用户
	CorpUnBindReqUrl           = "/corp/unbind"                //解绑企业用户
	GetCorpInfoReqUrl          = "/corp/get"                   //查询企业用户基本信息
	GetCorpIdentityInfoReqUrl  = "/corp/get-identity-info"     //获取企业用户身份信息
	GetCorpIdentifiedStatusUrl = "/corp/get-identified-status" //查询企业实名认证状态
)

// GetCorpAuthUrlResponse 获取企业用户授权链接
func (o *OpenApiClient) GetCorpAuthUrlResponse(gerCorpAuthUrlReq *corpManageRequestModel.GetCorpAuthUrlReq, accessToken string) corpManageResponseModel.GetCorpAuthUrlRes {
	var gerCorpAuthUrlRes corpManageResponseModel.GetCorpAuthUrlRes
	reqStr, err := json.Marshal(gerCorpAuthUrlReq)
	if err != nil {
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken, string(reqStr))
	requestUrl := o.serverUrl + GetCorpAuthReqUrl                      //接口请求api地址
	rspBody, requestId := common.SendPost(requestUrl, headMap, reqStr) //发送post请求
	err = json.Unmarshal(rspBody, &gerCorpAuthUrlRes)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}
	gerCorpAuthUrlRes.RequestId = requestId
	return gerCorpAuthUrlRes
}

// GetCorpDisableResponse 禁用企业用户
func (o *OpenApiClient) GetCorpDisableResponse(corpDisableReq *corpManageRequestModel.CorpDisableReq, accessToken string) corpManageResponseModel.CorpDisableRes {
	var CorpDisableRes corpManageResponseModel.CorpDisableRes
	reqStr, err := json.Marshal(corpDisableReq)
	if err != nil {
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken, string(reqStr))
	requestUrl := o.serverUrl + CorpDisableReqUrl                      //接口请求api地址
	rspBody, requestId := common.SendPost(requestUrl, headMap, reqStr) //发送post请求
	err = json.Unmarshal(rspBody, &CorpDisableRes)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}
	CorpDisableRes.RequestId = requestId
	return CorpDisableRes
}

// GetCorpEnableResponse 恢复企业用户
func (o *OpenApiClient) GetCorpEnableResponse(corpEnableReq *corpManageRequestModel.CorpEnableReq, accessToken string) corpManageResponseModel.CorpEnableRes {
	var corpEnableRes corpManageResponseModel.CorpEnableRes
	reqStr, err := json.Marshal(corpEnableReq)
	if err != nil {
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken, string(reqStr))
	requestUrl := o.serverUrl + CorpEnableReqUrl                       //接口请求api地址
	rspBody, requestId := common.SendPost(requestUrl, headMap, reqStr) //发送post请求
	err = json.Unmarshal(rspBody, &corpEnableRes)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}
	corpEnableRes.RequestId = requestId
	return corpEnableRes
}

// GetCorpUnBindResponse 解绑企业用户
func (o *OpenApiClient) GetCorpUnBindResponse(corpUnBindReq *corpManageRequestModel.UnBindCorpReq, accessToken string) corpManageResponseModel.UnBindCorpRes {
	var unBindCorpRes corpManageResponseModel.UnBindCorpRes
	reqStr, err := json.Marshal(corpUnBindReq)
	if err != nil {
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken, string(reqStr))
	requestUrl := o.serverUrl + CorpUnBindReqUrl                       //接口请求api地址
	rspBody, requestId := common.SendPost(requestUrl, headMap, reqStr) //发送post请求
	err = json.Unmarshal(rspBody, &unBindCorpRes)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}
	unBindCorpRes.RequestId = requestId
	return unBindCorpRes
}

// GetCorpInfoResponse 查询企业用户基本信息
func (o *OpenApiClient) GetCorpInfoResponse(corpInfoReq *corpManageRequestModel.GetCorpReq, accessToken string) corpManageResponseModel.GetCorpRes {
	var corpInfoRes corpManageResponseModel.GetCorpRes
	reqStr, err := json.Marshal(corpInfoReq)
	if err != nil {
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken, string(reqStr))
	requestUrl := o.serverUrl + GetCorpInfoReqUrl                      //接口请求api地址
	rspBody, requestId := common.SendPost(requestUrl, headMap, reqStr) //发送post请求
	err = json.Unmarshal(rspBody, &corpInfoRes)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}
	corpInfoRes.RequestId = requestId
	return corpInfoRes
}

// GetCorpIdentityResponse 获取企业用户身份信息
func (o *OpenApiClient) GetCorpIdentityResponse(corpIdentityReq *corpManageRequestModel.GetCorpIdentityInfoReq, accessToken string) corpManageResponseModel.GetCorpIdentityInfoRes {
	var corpIdentityRes corpManageResponseModel.GetCorpIdentityInfoRes
	reqStr, err := json.Marshal(corpIdentityReq)
	if err != nil {
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken, string(reqStr))
	requestUrl := o.serverUrl + GetCorpIdentityInfoReqUrl              //接口请求api地址
	rspBody, requestId := common.SendPost(requestUrl, headMap, reqStr) //发送post请求
	err = json.Unmarshal(rspBody, &corpIdentityRes)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}
	corpIdentityRes.RequestId = requestId
	return corpIdentityRes
}

// 查询企业实名认证状态
func (o *OpenApiClient) GetCorpIdentifiedStatus(getIdentifiedStatusReq *corpManageRequestModel.GetIdentifiedStatusReq, accessToken string) corpManageResponseModel.GetIdentifiedStatusRes {
	var response corpManageResponseModel.GetIdentifiedStatusRes
	reqStr, err := json.Marshal(getIdentifiedStatusReq)
	if err != nil {
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken, string(reqStr))
	requestUrl := o.serverUrl + GetCorpIdentifiedStatusUrl             //接口请求api地址
	rspBody, requestId := common.SendPost(requestUrl, headMap, reqStr) //发送post请求
	err = json.Unmarshal(rspBody, &response)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}
	response.RequestId = requestId
	return response
}
