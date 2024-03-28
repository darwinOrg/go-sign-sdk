package client

import (
	"encoding/json"
	"fmt"
	"github.com/darwinOrg/go-sign-sdk/fasc/common"
	"github.com/darwinOrg/go-sign-sdk/fasc/model/requestModel/userManageRequestModel"
	"github.com/darwinOrg/go-sign-sdk/fasc/model/responseModel/userManageResponseModel"
)

const (
	GetUserAuthUrl            = "/user/get-auth-url"                                 //获取个人授权链接
	UserDisableUrl            = "/user/disable"                                      //禁用个人用户
	UserEnableUrl             = "/user/enable"                                       //恢复个人用户
	UserUnBindUrl             = "/user/unbind"                                       //解绑个人用户
	GetUserInfoUrl            = "/user/get"                                          //查询个人用户基本信息
	GetIdentityInfoUrl        = "/user/get-identity-info"                            //获取个人用户身份信息
	GetThreeElementVerifyUrl  = "/user/three-element-verify/get-url"                 //个人三要素校验
	GetFourElementVerifyUrl   = "/user/four-element-verify/get-url"                  //个人四要素校验
	GetIdCardImageDownloadUrl = "/user/element-verify/get-idcard-image-download-url" //根据verifyId获取下载链接
)

// 获取个人用户授权链接
func (o *OpenApiClient) GetUserAuthUrlResponse(authUrlReq *userManageRequestModel.GetUserAuthUrlReq, accessToken string) userManageResponseModel.GetUserAuthUrlRes {
	var getAuthUrlRes userManageResponseModel.GetUserAuthUrlRes
	reqStr, err := json.Marshal(authUrlReq)
	if err != nil {
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken, string(reqStr))            //拼接post提交body参数
	requestUrl := o.serverUrl + GetUserAuthUrl                         //接口请求api地址
	rspBody, requestId := common.SendPost(requestUrl, headMap, reqStr) //发送post请求
	err = json.Unmarshal(rspBody, &getAuthUrlRes)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}
	getAuthUrlRes.RequestId = requestId
	return getAuthUrlRes
}

// GetUserDisableResponse 禁用个人用户
func (o *OpenApiClient) GetUserDisableResponse(userDisableReq *userManageRequestModel.UserDisableReq, accessToken string) userManageResponseModel.UserDisableRes {
	var userDisableRes userManageResponseModel.UserDisableRes
	reqStr, err := json.Marshal(userDisableReq)
	if err != nil {
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken, string(reqStr))
	requestUrl := o.serverUrl + UserDisableUrl                         //接口请求api地址
	rspBody, requestId := common.SendPost(requestUrl, headMap, reqStr) //发送post请求
	err = json.Unmarshal(rspBody, &userDisableRes)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}
	userDisableRes.RequestId = requestId
	return userDisableRes
}

// GetUserEnableResponse 恢复个人用户
func (o *OpenApiClient) GetUserEnableResponse(userEnableReq *userManageRequestModel.UserEnableReq, accessToken string) userManageResponseModel.UserEnableRes {
	var userEnableRes userManageResponseModel.UserEnableRes
	reqStr, err := json.Marshal(userEnableReq)
	if err != nil {
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken, string(reqStr))
	requestUrl := o.serverUrl + UserEnableUrl                          //接口请求api地址
	rspBody, requestId := common.SendPost(requestUrl, headMap, reqStr) //发送post请求
	err = json.Unmarshal(rspBody, &userEnableRes)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}
	userEnableRes.RequestId = requestId
	return userEnableRes
}

// GetUnBindUserResponse 解绑个人用户
func (o *OpenApiClient) GetUnBindUserResponse(unBindUserReq *userManageRequestModel.UnBindUserReq, accessToken string) userManageResponseModel.UnBindUserRes {
	var unBindUserRes userManageResponseModel.UnBindUserRes
	reqStr, err := json.Marshal(unBindUserReq)
	if err != nil {
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken, string(reqStr))
	requestUrl := o.serverUrl + UserUnBindUrl                          //接口请求api地址
	rspBody, requestId := common.SendPost(requestUrl, headMap, reqStr) //发送post请求
	err = json.Unmarshal(rspBody, &unBindUserRes)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}
	unBindUserRes.RequestId = requestId
	return unBindUserRes
}

// GetUserInfoResponse 查询个人用户基本信息
func (o *OpenApiClient) GetUserInfoResponse(userInfoReq *userManageRequestModel.GetUserReq, accessToken string) userManageResponseModel.GetUserRes {
	var userInfoRes userManageResponseModel.GetUserRes
	reqStr, err := json.Marshal(userInfoReq)
	if err != nil {
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken, string(reqStr))
	requestUrl := o.serverUrl + GetUserInfoUrl                         //接口请求api地址
	rspBody, requestId := common.SendPost(requestUrl, headMap, reqStr) //发送post请求
	err = json.Unmarshal(rspBody, &userInfoRes)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}
	userInfoRes.RequestId = requestId
	return userInfoRes
}

// GetUserIdentityResponse 获取个人用户身份信息
func (o *OpenApiClient) GetUserIdentityResponse(userIdentityReq *userManageRequestModel.GetUserIdentifyInfoReq, accessToken string) userManageResponseModel.GetUserIdentifyInfoRes {
	var userIdentityRes userManageResponseModel.GetUserIdentifyInfoRes
	reqStr, err := json.Marshal(userIdentityReq)
	if err != nil {
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken, string(reqStr))
	requestUrl := o.serverUrl + GetIdentityInfoUrl                     //接口请求api地址
	rspBody, requestId := common.SendPost(requestUrl, headMap, reqStr) //发送post请求
	err = json.Unmarshal(rspBody, &userIdentityRes)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}
	userIdentityRes.RequestId = requestId
	return userIdentityRes
}

// 个人三要素校验
func (o *OpenApiClient) GetThreeElementVerifyUrl(req *userManageRequestModel.GetThreeElementVerifyUrlReq, accessToken string) userManageResponseModel.GetThreeElementVerifyUrlRes {
	var res userManageResponseModel.GetThreeElementVerifyUrlRes
	reqStr, err := json.Marshal(req)
	if err != nil {
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken, string(reqStr))
	requestUrl := o.serverUrl + GetThreeElementVerifyUrl               //接口请求api地址
	rspBody, requestId := common.SendPost(requestUrl, headMap, reqStr) //发送post请求
	err = json.Unmarshal(rspBody, &res)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}
	res.RequestId = requestId
	return res
}

// 个人四要素校验
func (o *OpenApiClient) GetFourElementVerifyUrl(req *userManageRequestModel.GetFourElementVerifyUrlReq, accessToken string) userManageResponseModel.GetFourElementVerifyUrlRes {
	var res userManageResponseModel.GetFourElementVerifyUrlRes
	reqStr, err := json.Marshal(req)
	if err != nil {
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken, string(reqStr))
	requestUrl := o.serverUrl + GetFourElementVerifyUrl                //接口请求api地址
	rspBody, requestId := common.SendPost(requestUrl, headMap, reqStr) //发送post请求
	err = json.Unmarshal(rspBody, &res)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}
	res.RequestId = requestId
	return res
}

// 获取要素校验身份证图片下载链接
func (o *OpenApiClient) GetIdCardImageDownloadUrl(req *userManageRequestModel.GetIdCardImageDownloadUrlReq, accessToken string) userManageResponseModel.GetIdCardImageDownloadUrlRes {
	var res userManageResponseModel.GetIdCardImageDownloadUrlRes
	reqStr, err := json.Marshal(req)
	if err != nil {
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken, string(reqStr))
	requestUrl := o.serverUrl + GetIdCardImageDownloadUrl              //接口请求api地址
	rspBody, requestId := common.SendPost(requestUrl, headMap, reqStr) //发送post请求
	err = json.Unmarshal(rspBody, &res)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}
	res.RequestId = requestId
	return res
}
