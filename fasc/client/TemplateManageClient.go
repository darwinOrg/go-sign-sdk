package client

import (
	"encoding/json"
	"fmt"
	"github.com/darwinOrg/go-sign-sdk/fasc/common"
	"github.com/darwinOrg/go-sign-sdk/fasc/model/requestModel/templateRequestModel"
	"github.com/darwinOrg/go-sign-sdk/fasc/model/responseModel/templateResponseModel"
)

const (
	GetDocTemplateListReqUrl    = "/doc-template/get-list"        //查询文档模板列表
	GetDocTemplateDetailReqUrl  = "/doc-template/get-detail"      //获取文档模板详情
	GetSignTemplateListReqUrl   = "/sign-template/get-list"       //获取签署模板列表
	GetSignTemplateDetailReqUrl = "/sign-template/get-detail"     //获取签署模板详情
	GetTemplateCreateUrl        = "/template/create/get-url"      //获取模板新增链接
	GetTemplateEditUrl          = "/template/edit/get-url"        //获取模板编辑链接
	GetTemplatePreviewUrl       = "/template/preview/get-url"     //获取模板预览链接
	GetTemplateManageUrl        = "/template/manage/get-url"      //获取模板管理链接
	GetAppDocTemplatesUrl       = "/app-doc-template/get-list"    //查询应用文档模板列表
	GetAppDocTemplateDetailUrl  = "/app-doc-template/get-detail"  //获取应用文档模板详情
	SetAppDocTemplateStatusUrl  = "/app-doc-template/set-status"  //启用或停用应用文档模板
	DeleteAppDocTemplateUrl     = "/app-doc-template/delete"      //删除应用文档模板
	GetAppSignTemplatesUrl      = "/app-sign-template/get-list"   //查询应用签署任务模板列表
	GetAppSignTemplateDetailUrl = "/app-sign-template/get-detail" //获取应用签署任务模板详情
	SetAppSignTemplateStatusUrl = "/app-sign-template/set-status" //启用或停用应用签署任务模板
	DeleteAppSignTemplateUrl    = "/app-sign-template/delete"     //删除应用签署任务模板
	GetAppTemplateCreateUrl     = "/app-template/create/get-url"  //获取应用模板新增链接
	GetAppTemplateEditUrl       = "/app-template/edit/get-url"    //获取应用模板编辑链接
	GetAppTemplatePreviewUrl    = "/app-template/preview/get-url" //获取应用模板预览链接
	CreateAppFieldUrl           = "/app-field/create"             //创建业务控件
	ModifyAppFieldUrl           = "/app-field/modify"             //修改业务控件
	SetAppFieldStatusUrl        = "/app-field/set-status"         //设置业务控件状态
	GetAppFieldListUrl          = "/app-field/get-list"           //查询业务控件列表
)

// GetDocTemplateListResponse 查询文档模板列表
func (o *OpenApiClient) GetDocTemplateList(docTemListReq *templateRequestModel.GetDocTemplateListReq, accessToken string) templateResponseModel.GetDocTemplateListRes {
	var docTemListRes templateResponseModel.GetDocTemplateListRes
	reqStr, err := json.Marshal(docTemListReq)
	if err != nil {
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken, string(reqStr))
	requestUrl := o.serverUrl + GetDocTemplateListReqUrl               //接口请求api地址
	rspBody, requestId := common.SendPost(requestUrl, headMap, reqStr) //发送post请求
	err = json.Unmarshal(rspBody, &docTemListRes)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}
	docTemListRes.RequestId = requestId
	return docTemListRes
}

// GetDocTemplateDetailResponse 获取文档模板详情
func (o *OpenApiClient) GetDocTemplateDetail(docTemDetailReq *templateRequestModel.GetDocTemplateDetailReq, accessToken string) templateResponseModel.GetDocTemplateDetailRes {
	var docTemDetailRes templateResponseModel.GetDocTemplateDetailRes
	reqStr, err := json.Marshal(docTemDetailReq)
	if err != nil {
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken, string(reqStr))
	requestUrl := o.serverUrl + GetDocTemplateDetailReqUrl             //接口请求api地址
	rspBody, requestId := common.SendPost(requestUrl, headMap, reqStr) //发送post请求
	err = json.Unmarshal(rspBody, &docTemDetailRes)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}
	docTemDetailRes.RequestId = requestId
	return docTemDetailRes
}

// GetSignTemplateListResponse 查询签署模板列表
func (o *OpenApiClient) GetSignTemplateList(signTemplateReq *templateRequestModel.GetSignTemplateListReq, accessToken string) templateResponseModel.GetSignTemplateListRes {
	var signTemplateRes templateResponseModel.GetSignTemplateListRes
	reqStr, err := json.Marshal(signTemplateReq)
	if err != nil {
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken, string(reqStr))
	requestUrl := o.serverUrl + GetSignTemplateListReqUrl              //接口请求api地址
	rspBody, requestId := common.SendPost(requestUrl, headMap, reqStr) //发送post请求
	err = json.Unmarshal(rspBody, &signTemplateRes)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}
	signTemplateRes.RequestId = requestId
	return signTemplateRes
}

// GetSignTemplateDetailResponse 获取签署模板详情
func (o *OpenApiClient) GetSignTemplateDetail(signTemDetailReq *templateRequestModel.GetSignTemplateDetailReq, accessToken string) templateResponseModel.GetSignTemplateDetailRes {
	var signTemDetailRes templateResponseModel.GetSignTemplateDetailRes
	reqStr, err := json.Marshal(signTemDetailReq)
	if err != nil {
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken, string(reqStr))
	requestUrl := o.serverUrl + GetSignTemplateDetailReqUrl            //接口请求api地址
	rspBody, requestId := common.SendPost(requestUrl, headMap, reqStr) //发送post请求
	err = json.Unmarshal(rspBody, &signTemDetailRes)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}
	signTemDetailRes.RequestId = requestId
	return signTemDetailRes
}

// 获取模板新增链接
func (o *OpenApiClient) GetTemplateCreate(req *templateRequestModel.GetTemplateCreateUrlReq, accessToken string) templateResponseModel.GetTemplateCreateUrlRes {
	var res templateResponseModel.GetTemplateCreateUrlRes
	reqStr, err := json.Marshal(req)
	if err != nil {
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken, string(reqStr))
	requestUrl := o.serverUrl + GetTemplateCreateUrl                   //接口请求api地址
	rspBody, requestId := common.SendPost(requestUrl, headMap, reqStr) //发送post请求
	err = json.Unmarshal(rspBody, &res)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}
	res.RequestId = requestId
	return res
}

// 获取模板编辑链接
func (o *OpenApiClient) GetTemplateEdit(req *templateRequestModel.GetTemplateEditUrlReq, accessToken string) templateResponseModel.GetTemplateEditUrlRes {
	var res templateResponseModel.GetTemplateEditUrlRes
	reqStr, err := json.Marshal(req)
	if err != nil {
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken, string(reqStr))
	requestUrl := o.serverUrl + GetTemplateEditUrl                     //接口请求api地址
	rspBody, requestId := common.SendPost(requestUrl, headMap, reqStr) //发送post请求
	err = json.Unmarshal(rspBody, &res)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}
	res.RequestId = requestId
	return res
}

// 获取模板预览链接
func (o *OpenApiClient) GetTemplatePreview(req *templateRequestModel.GetTemplatePreviewUrlReq, accessToken string) templateResponseModel.GetTemplatePreviewUrlRes {
	var res templateResponseModel.GetTemplatePreviewUrlRes
	reqStr, err := json.Marshal(req)
	if err != nil {
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken, string(reqStr))
	requestUrl := o.serverUrl + GetTemplatePreviewUrl                  //接口请求api地址
	rspBody, requestId := common.SendPost(requestUrl, headMap, reqStr) //发送post请求
	err = json.Unmarshal(rspBody, &res)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}
	res.RequestId = requestId
	return res
}

// 获取模板管理链接
func (o *OpenApiClient) GetTemplateManage(req *templateRequestModel.GetTemplateManageUrlReq, accessToken string) templateResponseModel.GetTemplateManageUrlRes {
	var res templateResponseModel.GetTemplateManageUrlRes
	reqStr, err := json.Marshal(req)
	if err != nil {
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken, string(reqStr))
	requestUrl := o.serverUrl + GetTemplateManageUrl                   //接口请求api地址
	rspBody, requestId := common.SendPost(requestUrl, headMap, reqStr) //发送post请求
	err = json.Unmarshal(rspBody, &res)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}
	res.RequestId = requestId
	return res
}

// 查询应用文档模板列表
func (o *OpenApiClient) GetAppDocTemplates(req *templateRequestModel.GetAppDocTemplateListReq, accessToken string) templateResponseModel.AppDocTemplatePageListRes {
	var res templateResponseModel.AppDocTemplatePageListRes
	reqStr, err := json.Marshal(req)
	if err != nil {
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken, string(reqStr))
	requestUrl := o.serverUrl + GetAppDocTemplatesUrl                  //接口请求api地址
	rspBody, requestId := common.SendPost(requestUrl, headMap, reqStr) //发送post请求
	err = json.Unmarshal(rspBody, &res)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}
	res.RequestId = requestId
	return res
}

// 获取应用文档模板详情
func (o *OpenApiClient) GetAppDocTemplateDetail(req *templateRequestModel.GetAppDocTemplateDetailReq, accessToken string) templateResponseModel.AppDocTemplateDetailRes {
	var res templateResponseModel.AppDocTemplateDetailRes
	reqStr, err := json.Marshal(req)
	if err != nil {
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken, string(reqStr))
	requestUrl := o.serverUrl + GetAppDocTemplateDetailUrl             //接口请求api地址
	rspBody, requestId := common.SendPost(requestUrl, headMap, reqStr) //发送post请求
	err = json.Unmarshal(rspBody, &res)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}
	res.RequestId = requestId
	return res
}

// 启用/停用应用文档模板
func (o *OpenApiClient) SetAppDocTemplateStatus(req *templateRequestModel.SetAppDocTemplateStatusReq, accessToken string) templateResponseModel.SetAppDocTemplateStatusRes {
	var res templateResponseModel.SetAppDocTemplateStatusRes
	reqStr, err := json.Marshal(req)
	if err != nil {
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken, string(reqStr))
	requestUrl := o.serverUrl + SetAppDocTemplateStatusUrl             //接口请求api地址
	rspBody, requestId := common.SendPost(requestUrl, headMap, reqStr) //发送post请求
	err = json.Unmarshal(rspBody, &res)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}
	res.RequestId = requestId
	return res
}

// 删除应用文档模板
func (o *OpenApiClient) DeleteAppDocTemplate(req *templateRequestModel.DeleteAppDocTemplateReq, accessToken string) templateResponseModel.DeleteAppDocTemplateRes {
	var res templateResponseModel.DeleteAppDocTemplateRes
	reqStr, err := json.Marshal(req)
	if err != nil {
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken, string(reqStr))
	requestUrl := o.serverUrl + DeleteAppDocTemplateUrl                //接口请求api地址
	rspBody, requestId := common.SendPost(requestUrl, headMap, reqStr) //发送post请求
	err = json.Unmarshal(rspBody, &res)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}
	res.RequestId = requestId
	return res
}

// 查询应用签署任务模板列表
func (o *OpenApiClient) GetAppSignTemplates(req *templateRequestModel.GetAppSignTemplateListReq, accessToken string) templateResponseModel.AppSignTemplatePageListRes {
	var res templateResponseModel.AppSignTemplatePageListRes
	reqStr, err := json.Marshal(req)
	if err != nil {
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken, string(reqStr))
	requestUrl := o.serverUrl + GetAppSignTemplatesUrl                 //接口请求api地址
	rspBody, requestId := common.SendPost(requestUrl, headMap, reqStr) //发送post请求
	err = json.Unmarshal(rspBody, &res)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}
	res.RequestId = requestId
	return res
}

// 获取应用签署任务模板详情
func (o *OpenApiClient) GetAppSignTemplateDetail(req *templateRequestModel.GetAppSignTemplateDetailReq, accessToken string) templateResponseModel.AppSignTemplateDetailRes {
	var res templateResponseModel.AppSignTemplateDetailRes
	reqStr, err := json.Marshal(req)
	if err != nil {
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken, string(reqStr))
	requestUrl := o.serverUrl + GetAppSignTemplateDetailUrl            //接口请求api地址
	rspBody, requestId := common.SendPost(requestUrl, headMap, reqStr) //发送post请求
	err = json.Unmarshal(rspBody, &res)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}
	res.RequestId = requestId
	return res
}

// 启用/停用应用签署任务模板
func (o *OpenApiClient) SetAppSignTemplateStatus(req *templateRequestModel.SetAppSignTemplateStatusReq, accessToken string) templateResponseModel.SetAppSignTemplateStatusRes {
	var res templateResponseModel.SetAppSignTemplateStatusRes
	reqStr, err := json.Marshal(req)
	if err != nil {
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken, string(reqStr))
	requestUrl := o.serverUrl + SetAppSignTemplateStatusUrl            //接口请求api地址
	rspBody, requestId := common.SendPost(requestUrl, headMap, reqStr) //发送post请求
	err = json.Unmarshal(rspBody, &res)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}
	res.RequestId = requestId
	return res
}

// 删除应用签署任务模板
func (o *OpenApiClient) DeleteAppSignTemplate(req *templateRequestModel.DeleteAppSignTemplateReq, accessToken string) templateResponseModel.DeleteAppSignTemplateRes {
	var res templateResponseModel.DeleteAppSignTemplateRes
	reqStr, err := json.Marshal(req)
	if err != nil {
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken, string(reqStr))
	requestUrl := o.serverUrl + DeleteAppSignTemplateUrl               //接口请求api地址
	rspBody, requestId := common.SendPost(requestUrl, headMap, reqStr) //发送post请求
	err = json.Unmarshal(rspBody, &res)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}
	res.RequestId = requestId
	return res
}

// 获取应用模板新增链接
func (o *OpenApiClient) GetAppTemplateCreate(req *templateRequestModel.GetAppTemplateCreateUrlReq, accessToken string) templateResponseModel.GetAppTemplateCreateUrlRes {
	var res templateResponseModel.GetAppTemplateCreateUrlRes
	reqStr, err := json.Marshal(req)
	if err != nil {
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken, string(reqStr))
	requestUrl := o.serverUrl + GetAppTemplateCreateUrl                //接口请求api地址
	rspBody, requestId := common.SendPost(requestUrl, headMap, reqStr) //发送post请求
	err = json.Unmarshal(rspBody, &res)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}
	res.RequestId = requestId
	return res
}

// 获取应用模板编辑链接
func (o *OpenApiClient) GetAppTemplateEdit(req *templateRequestModel.GetAppTemplateEditUrlReq, accessToken string) templateResponseModel.GetAppTemplateEditUrlRes {
	var res templateResponseModel.GetAppTemplateEditUrlRes
	reqStr, err := json.Marshal(req)
	if err != nil {
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken, string(reqStr))
	requestUrl := o.serverUrl + GetAppTemplateEditUrl                  //接口请求api地址
	rspBody, requestId := common.SendPost(requestUrl, headMap, reqStr) //发送post请求
	err = json.Unmarshal(rspBody, &res)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}
	res.RequestId = requestId
	return res
}

// 获取应用模板预览链接
func (o *OpenApiClient) GetAppTemplatePreview(req *templateRequestModel.GetAppTemplatePreviewUrlReq, accessToken string) templateResponseModel.GetAppTemplatePreviewUrlRes {
	var res templateResponseModel.GetAppTemplatePreviewUrlRes
	reqStr, err := json.Marshal(req)
	if err != nil {
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken, string(reqStr))
	requestUrl := o.serverUrl + GetAppTemplatePreviewUrl               //接口请求api地址
	rspBody, requestId := common.SendPost(requestUrl, headMap, reqStr) //发送post请求
	err = json.Unmarshal(rspBody, &res)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}
	res.RequestId = requestId
	return res
}

// 创建业务控件
func (o *OpenApiClient) CreateAppField(req *templateRequestModel.CreateAppFieldReq, accessToken string) templateResponseModel.CreateAppFieldRes {
	var res templateResponseModel.CreateAppFieldRes
	reqStr, err := json.Marshal(req)
	if err != nil {
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken, string(reqStr))
	requestUrl := o.serverUrl + CreateAppFieldUrl                      //接口请求api地址
	rspBody, requestId := common.SendPost(requestUrl, headMap, reqStr) //发送post请求
	err = json.Unmarshal(rspBody, &res)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}
	res.RequestId = requestId
	return res
}

// 修改业务控件
func (o *OpenApiClient) ModifyAppField(req *templateRequestModel.ModifyAppFieldReq, accessToken string) templateResponseModel.ModifyAppFieldRes {
	var res templateResponseModel.ModifyAppFieldRes
	reqStr, err := json.Marshal(req)
	if err != nil {
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken, string(reqStr))
	requestUrl := o.serverUrl + ModifyAppFieldUrl                      //接口请求api地址
	rspBody, requestId := common.SendPost(requestUrl, headMap, reqStr) //发送post请求
	err = json.Unmarshal(rspBody, &res)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}
	res.RequestId = requestId
	return res
}

// 设置业务控件状态
func (o *OpenApiClient) SetAppFieldStatus(req *templateRequestModel.SetAppFieldStatusReq, accessToken string) templateResponseModel.SetAppFieldStatusRes {
	var res templateResponseModel.SetAppFieldStatusRes
	reqStr, err := json.Marshal(req)
	if err != nil {
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken, string(reqStr))
	requestUrl := o.serverUrl + SetAppFieldStatusUrl                   //接口请求api地址
	rspBody, requestId := common.SendPost(requestUrl, headMap, reqStr) //发送post请求
	err = json.Unmarshal(rspBody, &res)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}
	res.RequestId = requestId
	return res
}

// 查询业务控件列表
func (o *OpenApiClient) GetAppFieldList(req *templateRequestModel.GetAppFieldListReq, accessToken string) templateResponseModel.GetAppFieldListRes {
	var res templateResponseModel.GetAppFieldListRes
	reqStr, err := json.Marshal(req)
	if err != nil {
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken, string(reqStr))
	requestUrl := o.serverUrl + GetAppFieldListUrl                     //接口请求api地址
	rspBody, requestId := common.SendPost(requestUrl, headMap, reqStr) //发送post请求
	err = json.Unmarshal(rspBody, &res)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}
	res.RequestId = requestId
	return res
}
