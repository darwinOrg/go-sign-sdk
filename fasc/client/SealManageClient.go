package client

import (
	"encoding/json"
	"fmt"
	"github.com/darwinOrg/go-sign-sdk/fasc/common"
	"github.com/darwinOrg/go-sign-sdk/fasc/model/requestModel/sealRequestModel"
	"github.com/darwinOrg/go-sign-sdk/fasc/model/responseModel/sealResponseModel"
)

const (
	GetSealListUrl             = "/seal/get-list"                      //查询印章列表
	GetSealDetailUrl           = "/seal/get-detail"                    //查询印章详情
	GetAppointedSealUrl        = "/seal/manage/get-appointed-seal-url" //获取指定印章详情链接
	GetSealUserListUrl         = "/seal/get-user-list"                 //查询企业用印员列表
	GetUserSealListUrl         = "/seal/user/get-list"                 //查询指定成员的印章列表
	GetSealCreateUrl           = "/seal/create/get-url"                //获取印章创建链接
	GetVerifySealListUrl       = "/seal/verify/get-list"               //查询审核中的印章列表
	ModifySealInfoUrl          = "/seal/modify"                        //修改印章基本信息
	GetSealGrantUrl            = "/seal/grant/get-url"                 //获取设置用印员链接
	GetSealFreeSignUrl         = "/seal/free-sign/get-url"             //获取设置企业印章免验证签链接
	CancelSealGrantUrl         = "/seal/grant/cancel"                  //解除印章授权
	SetSealStatusUrl           = "/seal/set-status"                    //设置印章状态
	DeleteSealUrl              = "/seal/delete"                        //删除印章
	GetSealManageUrl           = "/seal/manage/get-url"                //获取印章管理链接
	GetPersonalSealListUrl     = "/personal-seal/get-list"             //查询个人签名列表
	GetPersonalSealFreeSignUrl = "/personal-seal/free-sign/get-url"    //获取设置个人签名免验证签链接
)

// GetSealListResponse 查询印章列表
func (o *OpenApiClient) GetSealList(sealListReq *sealRequestModel.GetSealListReq, accessToken string) sealResponseModel.GetSealListRes {
	var sealListRes sealResponseModel.GetSealListRes
	reqStr, err := json.Marshal(sealListReq)
	if err != nil {
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken, string(reqStr))
	requestUrl := o.serverUrl + GetSealListUrl                         //接口请求api地址
	rspBody, requestId := common.SendPost(requestUrl, headMap, reqStr) //发送post请求
	err = json.Unmarshal(rspBody, &sealListRes)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}
	sealListRes.RequestId = requestId
	return sealListRes
}

// 查询印章详情
func (o *OpenApiClient) GetSealDetail(req *sealRequestModel.GetSealDetailReq, accessToken string) sealResponseModel.GetSealDetailRes {
	var res sealResponseModel.GetSealDetailRes
	reqStr, err := json.Marshal(req)
	if err != nil {
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken, string(reqStr))
	requestUrl := o.serverUrl + GetSealDetailUrl                       //接口请求api地址
	rspBody, requestId := common.SendPost(requestUrl, headMap, reqStr) //发送post请求
	err = json.Unmarshal(rspBody, &res)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}
	res.RequestId = requestId
	return res
}

// 获取指定印章详情链接
func (o *OpenApiClient) GetAppointedSeal(req *sealRequestModel.GetAppointedSealUrlReq, accessToken string) sealResponseModel.GetAppointedSealUrlRes {
	var res sealResponseModel.GetAppointedSealUrlRes
	reqStr, err := json.Marshal(req)
	if err != nil {
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken, string(reqStr))
	requestUrl := o.serverUrl + GetAppointedSealUrl                    //接口请求api地址
	rspBody, requestId := common.SendPost(requestUrl, headMap, reqStr) //发送post请求
	err = json.Unmarshal(rspBody, &res)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}
	res.RequestId = requestId
	return res
}

// GetSealUserListResponse 查询企业用印员列表
func (o *OpenApiClient) GetSealUserList(sealUserListReq *sealRequestModel.GetSealUserListReq, accessToken string) sealResponseModel.GetSealUserListRes {
	var sealUserListRes sealResponseModel.GetSealUserListRes
	reqStr, err := json.Marshal(sealUserListReq)
	if err != nil {
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken, string(reqStr))
	requestUrl := o.serverUrl + GetSealUserListUrl                     //接口请求api地址
	rspBody, requestId := common.SendPost(requestUrl, headMap, reqStr) //发送post请求
	err = json.Unmarshal(rspBody, &sealUserListRes)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}
	sealUserListRes.RequestId = requestId
	return sealUserListRes
}

// 查询指定成员的印章列表
func (o *OpenApiClient) GetUserSealList(req *sealRequestModel.GetUserSealListReq, accessToken string) sealResponseModel.GetUserSealListRes {
	var res sealResponseModel.GetUserSealListRes
	reqStr, err := json.Marshal(req)
	if err != nil {
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken, string(reqStr))
	requestUrl := o.serverUrl + GetUserSealListUrl                     //接口请求api地址
	rspBody, requestId := common.SendPost(requestUrl, headMap, reqStr) //发送post请求
	err = json.Unmarshal(rspBody, &res)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}
	res.RequestId = requestId
	return res
}

// 获取印章创建链接
func (o *OpenApiClient) GetSealCreate(req *sealRequestModel.GetSealCreateUrlReq, accessToken string) sealResponseModel.GetSealCreateUrlRes {
	var res sealResponseModel.GetSealCreateUrlRes
	reqStr, err := json.Marshal(req)
	if err != nil {
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken, string(reqStr))
	requestUrl := o.serverUrl + GetSealCreateUrl                       //接口请求api地址
	rspBody, requestId := common.SendPost(requestUrl, headMap, reqStr) //发送post请求
	err = json.Unmarshal(rspBody, &res)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}
	res.RequestId = requestId
	return res
}

// 查询审核中的印章列表
func (o *OpenApiClient) GetVerifySealList(req *sealRequestModel.GetSealVerifyListReq, accessToken string) sealResponseModel.GetSealVerifyListRes {
	var res sealResponseModel.GetSealVerifyListRes
	reqStr, err := json.Marshal(req)
	if err != nil {
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken, string(reqStr))
	requestUrl := o.serverUrl + GetVerifySealListUrl                   //接口请求api地址
	rspBody, requestId := common.SendPost(requestUrl, headMap, reqStr) //发送post请求
	err = json.Unmarshal(rspBody, &res)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}
	res.RequestId = requestId
	return res
}

// 修改印章基本信息
func (o *OpenApiClient) ModifySealInfo(req *sealRequestModel.ModifySealReq, accessToken string) sealResponseModel.ModifySealRes {
	var res sealResponseModel.ModifySealRes
	reqStr, err := json.Marshal(req)
	if err != nil {
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken, string(reqStr))
	requestUrl := o.serverUrl + ModifySealInfoUrl                      //接口请求api地址
	rspBody, requestId := common.SendPost(requestUrl, headMap, reqStr) //发送post请求
	err = json.Unmarshal(rspBody, &res)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}
	res.RequestId = requestId
	return res
}

// 获取设置用印员链接
func (o *OpenApiClient) GetSealGrant(req *sealRequestModel.GetSealGrantUrlReq, accessToken string) sealResponseModel.GetSealGrantUrlRes {
	var res sealResponseModel.GetSealGrantUrlRes
	reqStr, err := json.Marshal(req)
	if err != nil {
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken, string(reqStr))
	requestUrl := o.serverUrl + GetSealGrantUrl                        //接口请求api地址
	rspBody, requestId := common.SendPost(requestUrl, headMap, reqStr) //发送post请求
	err = json.Unmarshal(rspBody, &res)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}
	res.RequestId = requestId
	return res
}

// 获取设置企业印章免验证签链接
func (o *OpenApiClient) GetSealFreeSign(req *sealRequestModel.GetSealFreeSignUrlReq, accessToken string) sealResponseModel.GetSealFreeSignUrlRes {
	var res sealResponseModel.GetSealFreeSignUrlRes
	reqStr, err := json.Marshal(req)
	if err != nil {
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken, string(reqStr))
	requestUrl := o.serverUrl + GetSealFreeSignUrl                     //接口请求api地址
	rspBody, requestId := common.SendPost(requestUrl, headMap, reqStr) //发送post请求
	err = json.Unmarshal(rspBody, &res)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}
	res.RequestId = requestId
	return res
}

// 解除印章授权
func (o *OpenApiClient) CancelSealGrant(req *sealRequestModel.CancelSealGrantReq, accessToken string) sealResponseModel.CancelSealGrantRes {
	var res sealResponseModel.CancelSealGrantRes
	reqStr, err := json.Marshal(req)
	if err != nil {
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken, string(reqStr))
	requestUrl := o.serverUrl + CancelSealGrantUrl                     //接口请求api地址
	rspBody, requestId := common.SendPost(requestUrl, headMap, reqStr) //发送post请求
	err = json.Unmarshal(rspBody, &res)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}
	res.RequestId = requestId
	return res
}

// 设置印章状态
func (o *OpenApiClient) SetSealStatus(req *sealRequestModel.SetSealStatusReq, accessToken string) sealResponseModel.SetSealStatusRes {
	var res sealResponseModel.SetSealStatusRes
	reqStr, err := json.Marshal(req)
	if err != nil {
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken, string(reqStr))
	requestUrl := o.serverUrl + SetSealStatusUrl                       //接口请求api地址
	rspBody, requestId := common.SendPost(requestUrl, headMap, reqStr) //发送post请求
	err = json.Unmarshal(rspBody, &res)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}
	res.RequestId = requestId
	return res
}

// 删除印章
func (o *OpenApiClient) DeleteSeal(req *sealRequestModel.SealDeleteReq, accessToken string) sealResponseModel.SealDeleteRes {
	var res sealResponseModel.SealDeleteRes
	reqStr, err := json.Marshal(req)
	if err != nil {
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken, string(reqStr))
	requestUrl := o.serverUrl + DeleteSealUrl                          //接口请求api地址
	rspBody, requestId := common.SendPost(requestUrl, headMap, reqStr) //发送post请求
	err = json.Unmarshal(rspBody, &res)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}
	res.RequestId = requestId
	return res
}

// 获取印章管理链接
func (o *OpenApiClient) GetSealManage(req *sealRequestModel.GetSealManageUrlReq, accessToken string) sealResponseModel.GetSealManageUrlRes {
	var res sealResponseModel.GetSealManageUrlRes
	reqStr, err := json.Marshal(req)
	if err != nil {
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken, string(reqStr))
	requestUrl := o.serverUrl + GetSealManageUrl                       //接口请求api地址
	rspBody, requestId := common.SendPost(requestUrl, headMap, reqStr) //发送post请求
	err = json.Unmarshal(rspBody, &res)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}
	res.RequestId = requestId
	return res
}

// 查询个人签名列表
func (o *OpenApiClient) GetPersonalSealList(req *sealRequestModel.GetPersonalSealListReq, accessToken string) sealResponseModel.GetPersonalSealListRes {
	var res sealResponseModel.GetPersonalSealListRes
	reqStr, err := json.Marshal(req)
	if err != nil {
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken, string(reqStr))
	requestUrl := o.serverUrl + GetPersonalSealListUrl                 //接口请求api地址
	rspBody, requestId := common.SendPost(requestUrl, headMap, reqStr) //发送post请求
	err = json.Unmarshal(rspBody, &res)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}
	res.RequestId = requestId
	return res
}

// 获取设置个人签名免验证签链接
func (o *OpenApiClient) GetPersonalSealFreeSign(req *sealRequestModel.GetPersonalSealFreeSignUrlReq, accessToken string) sealResponseModel.GetSealFreeSignUrlRes {
	var res sealResponseModel.GetSealFreeSignUrlRes
	reqStr, err := json.Marshal(req)
	if err != nil {
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken, string(reqStr))
	requestUrl := o.serverUrl + GetPersonalSealFreeSignUrl             //接口请求api地址
	rspBody, requestId := common.SendPost(requestUrl, headMap, reqStr) //发送post请求
	err = json.Unmarshal(rspBody, &res)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}
	res.RequestId = requestId
	return res
}
