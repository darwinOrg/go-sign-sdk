package client

import (
	"encoding/json"
	"fmt"
	"github.com/darwinOrg/go-sign-sdk/fasc/common"
	"github.com/darwinOrg/go-sign-sdk/fasc/model/requestModel/docProcessRequestModel"
	"github.com/darwinOrg/go-sign-sdk/fasc/model/responseModel/docProcessResponseModel"
)

const (
	GetFddFileUrlUrl               = "/file/upload-by-url"           //通过网络文件地址上传
	GetLocalUploadUrlUrl           = "/file/get-upload-url"          //获取本地文件上传url
	FileProcessUrl                 = "/file/process"                 //文件处理
	FileVerifySignUrl              = "/file/verify-sign"             //文档验签
	GetOcrEditCompareUrl           = "/ocr/edit/get-compare-url"     //获取文件对比页面链接
	GetOcrEditCompareResultUrl     = "/ocr/edit/compare-result-url"  //获取历史文件对比页面链接
	GetOcrEditExamineUrl           = "/ocr/edit/get-examine-url"     //获取合同智审页面链接
	GetOcrEditExamineResultUrl     = "/ocr/edit/examine-result-url"  //获取历史合同智审页面链接
	GetOcrEditExamineResultDataUrl = "/ocr/edit/examine-result-data" //获取历史合同智审结果数据
)

// GetFddFileUploadUrl 通过网络文件地址上传
func (o *OpenApiClient) GetFddFileUploadUrl(uploadFileReq *docProcessRequestModel.UploadFileByUrlReq, accessToken string) docProcessResponseModel.UploadFileByUrlRes {
	var uploadFileRes docProcessResponseModel.UploadFileByUrlRes
	reqStr, err := json.Marshal(uploadFileReq)
	if err != nil {
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken, string(reqStr))
	requestUrl := o.serverUrl + GetFddFileUrlUrl                       //接口请求api地址
	rspBody, requestId := common.SendPost(requestUrl, headMap, reqStr) //发送post请求
	err = json.Unmarshal(rspBody, &uploadFileRes)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}
	uploadFileRes.RequestId = requestId
	return uploadFileRes
}

// GetLocalFileUploadUrl 获取文件本地上传地址
func (o *OpenApiClient) GetLocalFileUploadUrl(localUploadUrlReq *docProcessRequestModel.GetLocalUploadFileUrlReq, accessToken string) docProcessResponseModel.GetLocalUploadFileUrlRes {
	var localUploadUrlRes docProcessResponseModel.GetLocalUploadFileUrlRes
	reqStr, err := json.Marshal(localUploadUrlReq)
	if err != nil {
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken, string(reqStr))
	requestUrl := o.serverUrl + GetLocalUploadUrlUrl                   //接口请求api地址
	rspBody, requestId := common.SendPost(requestUrl, headMap, reqStr) //发送post请求
	err = json.Unmarshal(rspBody, &localUploadUrlRes)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}
	localUploadUrlRes.RequestId = requestId
	return localUploadUrlRes
}

// PutFileUpload put上传本地
func (o *OpenApiClient) PutFileUpload(filePath string, uploadUrl string) error {
	return common.PutFileUpload(uploadUrl, filePath)
}

// FileProcess 文件处理
func (o *OpenApiClient) FileProcess(fileProcessReq *docProcessRequestModel.FileProcessReq, accessToken string) docProcessResponseModel.FileProcessRes {
	var fileProcessRes docProcessResponseModel.FileProcessRes
	reqStr, err := json.Marshal(fileProcessReq)
	if err != nil {
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken, string(reqStr))
	requestUrl := o.serverUrl + FileProcessUrl                         //接口请求api地址
	rspBody, requestId := common.SendPost(requestUrl, headMap, reqStr) //发送post请求
	err = json.Unmarshal(rspBody, &fileProcessRes)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}
	fileProcessRes.RequestId = requestId
	return fileProcessRes
}

// 文档验签
func (o *OpenApiClient) VerifyFileSign(req *docProcessRequestModel.VerifyFileSignReq, accessToken string) docProcessResponseModel.VerifyFileSignRes {
	var res docProcessResponseModel.VerifyFileSignRes
	reqStr, err := json.Marshal(req)
	if err != nil {
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken, string(reqStr))
	requestUrl := o.serverUrl + FileVerifySignUrl                      //接口请求api地址
	rspBody, requestId := common.SendPost(requestUrl, headMap, reqStr) //发送post请求
	err = json.Unmarshal(rspBody, &res)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}
	res.RequestId = requestId
	return res
}

// 获取文件对比页面链接
func (o *OpenApiClient) GetCompareUrl(req *docProcessRequestModel.GetCompareUrlReq, accessToken string) docProcessResponseModel.GetCompareUrlRes {
	var res docProcessResponseModel.GetCompareUrlRes
	reqStr, err := json.Marshal(req)
	if err != nil {
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken, string(reqStr))
	requestUrl := o.serverUrl + GetOcrEditCompareUrl                   //接口请求api地址
	rspBody, requestId := common.SendPost(requestUrl, headMap, reqStr) //发送post请求
	err = json.Unmarshal(rspBody, &res)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}
	res.RequestId = requestId
	return res
}

// 获取历史文件对比页面链接
func (o *OpenApiClient) GetCompareResultUrl(req *docProcessRequestModel.GetCompareResultUrlReq, accessToken string) docProcessResponseModel.GetCompareResultUrlRes {
	var res docProcessResponseModel.GetCompareResultUrlRes
	reqStr, err := json.Marshal(req)
	if err != nil {
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken, string(reqStr))
	requestUrl := o.serverUrl + GetOcrEditCompareResultUrl             //接口请求api地址
	rspBody, requestId := common.SendPost(requestUrl, headMap, reqStr) //发送post请求
	err = json.Unmarshal(rspBody, &res)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}
	res.RequestId = requestId
	return res
}

// 获取合同智审页面链接
func (o *OpenApiClient) GetExamineUrl(req *docProcessRequestModel.GetExamineUrlReq, accessToken string) docProcessResponseModel.GetExamineUrlRes {
	var res docProcessResponseModel.GetExamineUrlRes
	reqStr, err := json.Marshal(req)
	if err != nil {
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken, string(reqStr))
	requestUrl := o.serverUrl + GetOcrEditExamineUrl                   //接口请求api地址
	rspBody, requestId := common.SendPost(requestUrl, headMap, reqStr) //发送post请求
	err = json.Unmarshal(rspBody, &res)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}
	res.RequestId = requestId
	return res
}

// 获取历史合同智审页面链接
func (o *OpenApiClient) GetExamineResultUrl(req *docProcessRequestModel.GetExamineResultUrlReq, accessToken string) docProcessResponseModel.GetExamineResultUrlRes {
	var res docProcessResponseModel.GetExamineResultUrlRes
	reqStr, err := json.Marshal(req)
	if err != nil {
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken, string(reqStr))
	requestUrl := o.serverUrl + GetOcrEditExamineResultUrl             //接口请求api地址
	rspBody, requestId := common.SendPost(requestUrl, headMap, reqStr) //发送post请求
	err = json.Unmarshal(rspBody, &res)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}
	res.RequestId = requestId
	return res
}

// 获取历史合同智审结果数据
func (o *OpenApiClient) GetOcrEditExamineResultData(req *docProcessRequestModel.GetOcrEditExamineResultDataReq, accessToken string) docProcessResponseModel.GetOcrEditExamineResultDataRes {
	var res docProcessResponseModel.GetOcrEditExamineResultDataRes
	reqStr, err := json.Marshal(req)
	if err != nil {
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken, string(reqStr))
	requestUrl := o.serverUrl + GetOcrEditExamineResultDataUrl         //接口请求api地址
	rspBody, requestId := common.SendPost(requestUrl, headMap, reqStr) //发送post请求
	err = json.Unmarshal(rspBody, &res)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}
	res.RequestId = requestId
	return res
}
