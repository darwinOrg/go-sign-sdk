package client

import (
	"encoding/json"
	"fmt"
	"github.com/darwinOrg/go-sign-sdk/fasc/common"
	"github.com/darwinOrg/go-sign-sdk/fasc/model/requestModel/signtaskRequestModel"
	"github.com/darwinOrg/go-sign-sdk/fasc/model/responseModel/signtaskResponseModel"
)

const (
	CreateSignTaskUrl                    = "/sign-task/create"                           //创建签署任务
	CreateSignTaskWithTemplateUrl        = "/sign-task/create-with-template"             //创建签署任务基于签署模板
	AddSignTaskDocUrl                    = "/sign-task/doc/add"                          //添加签署任务文档
	DeleteSignTaskDocUrl                 = "/sign-task/doc/delete"                       //移除签署任务文档
	AddSignTaskFieldUrl                  = "/sign-task/field/add"                        //添加签署任务控件
	DeleteSignTaskFieldUrl               = "/sign-task/field/delete"                     //移除签署任务控件
	AddSignTaskAttachUrl                 = "/sign-task/attach/add"                       //添加签署任务附件
	DeleteSignTaskAttachUrl              = "/sign-task/attach/delete"                    //移除签署任务附件
	AddSignTaskActorUrl                  = "/sign-task/actor/add"                        //添加签署任务参与方
	DeleteSignTaskActorUrl               = "/sign-task/actor/delete"                     //移除签署任务参与方
	FillSignTaskFieldValuesUrl           = "/sign-task/field/fill-values"                //填写签署任务控件内容
	StartSignTaskUrl                     = "/sign-task/start"                            //提交签署任务
	GetSignTaskActorUrl                  = "/sign-task/actor/get-url"                    //获取签署任务参与方专属链接
	FinalizeSignTaskDocUrl               = "/sign-task/doc-finalize"                     //定稿签署任务文档
	FinishSignTaskUrl                    = "/sign-task/finish"                           //结束签署任务                   //
	DeleteSignTaskUrl                    = "/sign-task/delete"                           //删除签署任务
	AbolishSignTaskUrl                   = "/sign-task/abolish"                          //作废签署任务
	UrgeSignTaskUrl                      = "/sign-task/urge"                             //催办签署任务
	BlockSignTaskUrl                     = "/sign-task/block"                            //阻塞签署任务
	UnblockSignTaskUrl                   = "/sign-task/unblock"                          //解阻签署任务
	CancelSignTaskUrl                    = "/sign-task/cancel"                           //撤销签署任务
	GetSignTaskEditUrl                   = "/sign-task/get-edit-url"                     //获取签署任务编辑链接
	GetSignTaskPreviewUrl                = "/sign-task/get-preview-url"                  //获取签署任务预览链接
	GetBatchSignUrl                      = "/sign-task/get-batch-sign-url"               //获取签署任务批量签署链接
	GetOwnerSignTaskListUrl              = "/sign-task/owner/get-list"                   //获取指定归属方的签署任务列表
	GetSignTaskDetailUrl                 = "/sign-task/app/get-detail"                   //获取签署任务详情
	GetOwnerSignTaskUrlUrl               = "/sign-task/owner/get-download-url"           //获取指定归属方的签署任务文档下载地址
	GetSignTaskCatalogListUrl            = "/sign-task-catalog/list"                     //查询企业签署任务文件夹
	GetSignTaskFieldListUrl              = "/sign-task/field/list"                       //查询签署任务控件信息
	GetSignTaskActorListUrl              = "/sign-task/actor/list"                       //查询签署任务参与方信息
	GetSignTaskActorFacePictureUrl       = "/sign-task/actor/get-face-picture"           //查询参与方的签署刷脸底图
	GetSignTaskSlicingTicketIdUrl        = "/sign-task/owner/get-slicing-ticket-id"      //签署文档切图
	GetSignTaskPicDownloadUrl            = "/sign-task/owner/get-pic-download-url"       //获取图片版签署文档下载地址
	GetApprovalInfoUrl                   = "/sign-task/get-approval-info"                //查询签署任务审批信息
	GetSignTaskDownLoadEvidenceReportUrl = "/sign-task/evidence-report/get-download-url" //获取签署任务公证处保全报告下载地址
	GetSignTaskBusinessTypeListUrl       = "/sign-task/business-type/get-list"           //查询签署业务类型列表      //
)

// CreateSignTask 创建签署任务
func (o *OpenApiClient) CreateSignTask(createSignTaskReq *signtaskRequestModel.CreateSignTaskReq, accessToken string) signtaskResponseModel.CreateSignTaskRes {
	var createSignTaskRes signtaskResponseModel.CreateSignTaskRes
	reqStr, err := json.Marshal(createSignTaskReq)
	if err != nil {
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken, string(reqStr))
	requestUrl := o.serverUrl + CreateSignTaskUrl                      //接口请求api地址
	rspBody, requestId := common.SendPost(requestUrl, headMap, reqStr) //发送post请求
	err = json.Unmarshal(rspBody, &createSignTaskRes)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}
	createSignTaskRes.RequestId = requestId
	return createSignTaskRes
}

// CreateSignTaskWithTemplate 创建签署任务基于签署模板
func (o *OpenApiClient) CreateSignTaskWithTemplate(addSignTaskReq *signtaskRequestModel.CreateSignTaskWithTemplateReq, accessToken string) signtaskResponseModel.CreateSignTaskWithTemplateRes {
	var addSignTaskRes signtaskResponseModel.CreateSignTaskWithTemplateRes
	reqStr, err := json.Marshal(addSignTaskReq)
	if err != nil {
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken, string(reqStr))
	requestUrl := o.serverUrl + CreateSignTaskWithTemplateUrl          //接口请求api地址
	rspBody, requestId := common.SendPost(requestUrl, headMap, reqStr) //发送post请求
	err = json.Unmarshal(rspBody, &addSignTaskRes)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}
	addSignTaskRes.RequestId = requestId
	return addSignTaskRes
}

// AddSignTaskDoc 添加签署任务文档
func (o *OpenApiClient) AddSignTaskDoc(addSignTaskDocReq *signtaskRequestModel.AddSignTaskDocReq, accessToken string) signtaskResponseModel.AddSignTaskDocRes {
	var addSignTaskDocRes signtaskResponseModel.AddSignTaskDocRes
	reqStr, err := json.Marshal(addSignTaskDocReq)
	if err != nil {
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken, string(reqStr))
	requestUrl := o.serverUrl + AddSignTaskDocUrl                      //接口请求api地址
	rspBody, requestId := common.SendPost(requestUrl, headMap, reqStr) //发送post请求
	err = json.Unmarshal(rspBody, &addSignTaskDocRes)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}
	addSignTaskDocRes.RequestId = requestId
	return addSignTaskDocRes
}

// DeleteSignTaskDoc 移除签署任务文档
func (o *OpenApiClient) DeleteSignTaskDoc(deleteSignTaskDocReq *signtaskRequestModel.DeleteSignTaskDocReq, accessToken string) signtaskResponseModel.DeleteSignTaskDocRes {
	var deleteSignTaskDocRes signtaskResponseModel.DeleteSignTaskDocRes
	reqStr, err := json.Marshal(deleteSignTaskDocReq)
	if err != nil {
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken, string(reqStr))
	requestUrl := o.serverUrl + DeleteSignTaskDocUrl                   //接口请求api地址
	rspBody, requestId := common.SendPost(requestUrl, headMap, reqStr) //发送post请求
	err = json.Unmarshal(rspBody, &deleteSignTaskDocRes)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}
	deleteSignTaskDocRes.RequestId = requestId
	return deleteSignTaskDocRes
}

// AddSignTaskField 添加签署任务控件
func (o *OpenApiClient) AddSignTaskField(addSignTaskFieldReq *signtaskRequestModel.AddSignTaskFieldReq, accessToken string) signtaskResponseModel.AddSignTaskFieldRes {
	var addSignTaskFieldRes signtaskResponseModel.AddSignTaskFieldRes
	reqStr, err := json.Marshal(addSignTaskFieldReq)
	if err != nil {
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken, string(reqStr))
	requestUrl := o.serverUrl + AddSignTaskFieldUrl                    //接口请求api地址
	rspBody, requestId := common.SendPost(requestUrl, headMap, reqStr) //发送post请求
	err = json.Unmarshal(rspBody, &addSignTaskFieldRes)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}
	addSignTaskFieldRes.RequestId = requestId
	return addSignTaskFieldRes
}

// DeleteSignTaskField 移除签署任务控件
func (o *OpenApiClient) DeleteSignTaskField(deleteSignTaskFieldReq *signtaskRequestModel.DeleteSignTaskFieldReq, accessToken string) signtaskResponseModel.DeleteSignTaskFieldRes {
	var deleteSignTaskFieldRes signtaskResponseModel.DeleteSignTaskFieldRes
	reqStr, err := json.Marshal(deleteSignTaskFieldReq)
	if err != nil {
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken, string(reqStr))
	requestUrl := o.serverUrl + DeleteSignTaskFieldUrl                 //接口请求api地址
	rspBody, requestId := common.SendPost(requestUrl, headMap, reqStr) //发送post请求
	err = json.Unmarshal(rspBody, &deleteSignTaskFieldRes)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}
	deleteSignTaskFieldRes.RequestId = requestId
	return deleteSignTaskFieldRes
}

// AddSignTaskAttach 添加签署任务附件
func (o *OpenApiClient) AddSignTaskAttach(addAttachReq *signtaskRequestModel.AddSignTaskAttachReq, accessToken string) signtaskResponseModel.AddSignTaskAttachRes {
	var addAttachRes signtaskResponseModel.AddSignTaskAttachRes
	reqStr, err := json.Marshal(addAttachReq)
	if err != nil {
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken, string(reqStr))
	requestUrl := o.serverUrl + AddSignTaskAttachUrl                   //接口请求api地址
	rspBody, requestId := common.SendPost(requestUrl, headMap, reqStr) //发送post请求
	err = json.Unmarshal(rspBody, &addAttachRes)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}
	addAttachRes.RequestId = requestId
	return addAttachRes
}

// DeleteSignTaskAttach 移除签署任务附件
func (o *OpenApiClient) DeleteSignTaskAttach(deleteAttachReq *signtaskRequestModel.DeleteSignTaskAttachReq, accessToken string) signtaskResponseModel.DeleteSignTaskAttachRes {
	var deleteAttachRes signtaskResponseModel.DeleteSignTaskAttachRes
	reqStr, err := json.Marshal(deleteAttachReq)
	if err != nil {
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken, string(reqStr))
	requestUrl := o.serverUrl + DeleteSignTaskAttachUrl                //接口请求api地址
	rspBody, requestId := common.SendPost(requestUrl, headMap, reqStr) //发送post请求
	err = json.Unmarshal(rspBody, &deleteAttachRes)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}
	deleteAttachRes.RequestId = requestId
	return deleteAttachRes
}

// AddSignTaskActor 添加签署任务参与方
func (o *OpenApiClient) AddSignTaskActor(addActorReq *signtaskRequestModel.AddSignTaskActorReq, accessToken string) signtaskResponseModel.AddSignTaskActorRes {
	var addActorRes signtaskResponseModel.AddSignTaskActorRes
	reqStr, err := json.Marshal(addActorReq)
	if err != nil {
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken, string(reqStr))
	requestUrl := o.serverUrl + AddSignTaskActorUrl                    //接口请求api地址
	rspBody, requestId := common.SendPost(requestUrl, headMap, reqStr) //发送post请求
	err = json.Unmarshal(rspBody, &addActorRes)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}
	addActorRes.RequestId = requestId
	return addActorRes
}

// DeleteSignTaskActor 移除签署任务参与方
func (o *OpenApiClient) DeleteSignTaskActor(deleteActorReq *signtaskRequestModel.DeleteSignTaskActorReq, accessToken string) signtaskResponseModel.DeleteSignTaskActorRes {
	var deleteActorRes signtaskResponseModel.DeleteSignTaskActorRes
	reqStr, err := json.Marshal(deleteActorReq)
	if err != nil {
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken, string(reqStr))
	requestUrl := o.serverUrl + DeleteSignTaskActorUrl                 //接口请求api地址
	rspBody, requestId := common.SendPost(requestUrl, headMap, reqStr) //发送post请求
	err = json.Unmarshal(rspBody, &deleteActorRes)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}
	deleteActorRes.RequestId = requestId
	return deleteActorRes
}

// FillSignTaskFieldValues 填写签署任务控件内容
func (o *OpenApiClient) FillSignTaskFieldValues(fillFieldReq *signtaskRequestModel.FillFieldValuesReq, accessToken string) signtaskResponseModel.FillFieldValuesRes {
	var fillFieldRes signtaskResponseModel.FillFieldValuesRes
	reqStr, err := json.Marshal(fillFieldReq)
	if err != nil {
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken, string(reqStr))
	requestUrl := o.serverUrl + FillSignTaskFieldValuesUrl             //接口请求api地址
	rspBody, requestId := common.SendPost(requestUrl, headMap, reqStr) //发送post请求
	err = json.Unmarshal(rspBody, &fillFieldRes)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}
	fillFieldRes.RequestId = requestId
	return fillFieldRes
}

// StartSignTask 提交签署任务
func (o *OpenApiClient) StartSignTask(req *signtaskRequestModel.StartSignTaskReq, accessToken string) signtaskResponseModel.StartSignTaskRes {
	var startSignTaskRes signtaskResponseModel.StartSignTaskRes
	reqStr, err := json.Marshal(req)
	if err != nil {
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken, string(reqStr))
	requestUrl := o.serverUrl + StartSignTaskUrl                       //接口请求api地址
	rspBody, requestId := common.SendPost(requestUrl, headMap, reqStr) //发送post请求
	err = json.Unmarshal(rspBody, &startSignTaskRes)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}
	startSignTaskRes.RequestId = requestId
	return startSignTaskRes
}

// 获取参与方专属链接
func (o *OpenApiClient) GetSignTaskActorUrl(req *signtaskRequestModel.SignTaskActorGetUrlReq, accessToken string) signtaskResponseModel.SignTaskActorGetUrlRes {
	var res signtaskResponseModel.SignTaskActorGetUrlRes
	reqStr, err := json.Marshal(req)
	if err != nil {
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken, string(reqStr))
	requestUrl := o.serverUrl + GetSignTaskActorUrl                    //接口请求api地址
	rspBody, requestId := common.SendPost(requestUrl, headMap, reqStr) //发送post请求
	err = json.Unmarshal(rspBody, &res)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}
	res.RequestId = requestId
	return res
}

// FinalizeSignTaskDoc 定稿签署任务文档
func (o *OpenApiClient) FinalizeSignTaskDoc(finalizeDocReq *signtaskRequestModel.FinalizeSignTaskReq, accessToken string) signtaskResponseModel.FinalizeSignTaskRes {
	var finalizeDocRes signtaskResponseModel.FinalizeSignTaskRes
	reqStr, err := json.Marshal(finalizeDocReq)
	if err != nil {
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken, string(reqStr))
	requestUrl := o.serverUrl + FinalizeSignTaskDocUrl                 //接口请求api地址
	rspBody, requestId := common.SendPost(requestUrl, headMap, reqStr) //发送post请求
	err = json.Unmarshal(rspBody, &finalizeDocRes)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}
	finalizeDocRes.RequestId = requestId
	return finalizeDocRes
}

// 结束签署任务
func (o *OpenApiClient) FinishSignTask(finalizeDocReq *signtaskRequestModel.FinishSignTaskReq, accessToken string) signtaskResponseModel.FinishSignTaskRes {
	var res signtaskResponseModel.FinishSignTaskRes
	reqStr, err := json.Marshal(finalizeDocReq)
	if err != nil {
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken, string(reqStr))
	requestUrl := o.serverUrl + FinishSignTaskUrl                      //接口请求api地址
	rspBody, requestId := common.SendPost(requestUrl, headMap, reqStr) //发送post请求
	err = json.Unmarshal(rspBody, &res)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}
	res.RequestId = requestId
	return res
}

// 删除签署任务
func (o *OpenApiClient) DeleteSignTask(finalizeDocReq *signtaskRequestModel.DeleteSignTaskReq, accessToken string) signtaskResponseModel.DeleteSignTaskRes {
	var res signtaskResponseModel.DeleteSignTaskRes
	reqStr, err := json.Marshal(finalizeDocReq)
	if err != nil {
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken, string(reqStr))
	requestUrl := o.serverUrl + DeleteSignTaskUrl                      //接口请求api地址
	rspBody, requestId := common.SendPost(requestUrl, headMap, reqStr) //发送post请求
	err = json.Unmarshal(rspBody, &res)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}
	res.RequestId = requestId
	return res
}

// 作废签署任务
func (o *OpenApiClient) AbolishSignTask(finalizeDocReq *signtaskRequestModel.AbolishSignTaskReq, accessToken string) signtaskResponseModel.AbolishSignTaskRes {
	var res signtaskResponseModel.AbolishSignTaskRes
	reqStr, err := json.Marshal(finalizeDocReq)
	if err != nil {
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken, string(reqStr))
	requestUrl := o.serverUrl + AbolishSignTaskUrl                     //接口请求api地址
	rspBody, requestId := common.SendPost(requestUrl, headMap, reqStr) //发送post请求
	err = json.Unmarshal(rspBody, &res)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}
	res.RequestId = requestId
	return res
}

// 催办签署任务
func (o *OpenApiClient) UrgeSignTask(req *signtaskRequestModel.SignTaskUrgeReq, accessToken string) signtaskResponseModel.SignTaskUrgeRes {
	var res signtaskResponseModel.SignTaskUrgeRes
	reqStr, err := json.Marshal(req)
	if err != nil {
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken, string(reqStr))
	requestUrl := o.serverUrl + UrgeSignTaskUrl                        //接口请求api地址
	rspBody, requestId := common.SendPost(requestUrl, headMap, reqStr) //发送post请求
	err = json.Unmarshal(rspBody, &res)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}
	res.RequestId = requestId
	return res
}

// BlockSignTask 阻塞签署任务
func (o *OpenApiClient) BlockSignTask(blockSignTaskReq *signtaskRequestModel.BlockSignTaskReq, accessToken string) signtaskResponseModel.BlockSignTaskRes {
	var blockSignTaskRes signtaskResponseModel.BlockSignTaskRes
	reqStr, err := json.Marshal(blockSignTaskReq)
	if err != nil {
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken, string(reqStr))
	requestUrl := o.serverUrl + BlockSignTaskUrl                       //接口请求api地址
	rspBody, requestId := common.SendPost(requestUrl, headMap, reqStr) //发送post请求
	err = json.Unmarshal(rspBody, &blockSignTaskRes)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}
	blockSignTaskRes.RequestId = requestId
	return blockSignTaskRes
}

// UnBlockSignTask 解阻签署任务
func (o *OpenApiClient) UnBlockSignTask(unBlockSignTaskReq *signtaskRequestModel.UnBlockSignTaskReq, accessToken string) signtaskResponseModel.UnBlockSignTaskRes {
	var unBlockSignTaskRes signtaskResponseModel.UnBlockSignTaskRes
	reqStr, err := json.Marshal(unBlockSignTaskReq)
	if err != nil {
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken, string(reqStr))
	requestUrl := o.serverUrl + UnblockSignTaskUrl                     //接口请求api地址
	rspBody, requestId := common.SendPost(requestUrl, headMap, reqStr) //发送post请求
	err = json.Unmarshal(rspBody, &unBlockSignTaskRes)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}
	unBlockSignTaskRes.RequestId = requestId
	return unBlockSignTaskRes
}

// CancelSignTask 撤销签署任务
func (o *OpenApiClient) CancelSignTask(cancelSignTaskReq *signtaskRequestModel.CancelSignTaskReq, accessToken string) signtaskResponseModel.CancelSignTaskRes {
	var cancelSignTaskRes signtaskResponseModel.CancelSignTaskRes
	reqStr, err := json.Marshal(cancelSignTaskReq)
	if err != nil {
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken, string(reqStr))
	requestUrl := o.serverUrl + CancelSignTaskUrl                      //接口请求api地址
	rspBody, requestId := common.SendPost(requestUrl, headMap, reqStr) //发送post请求
	err = json.Unmarshal(rspBody, &cancelSignTaskRes)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}
	cancelSignTaskRes.RequestId = requestId
	return cancelSignTaskRes
}

// 获取签署任务编辑链接
func (o *OpenApiClient) GetSignTaskEditUrl(req *signtaskRequestModel.GetSignTaskUrlReq, accessToken string) signtaskResponseModel.GetSignTaskEditUrlRes {
	var res signtaskResponseModel.GetSignTaskEditUrlRes
	reqStr, err := json.Marshal(req)
	if err != nil {
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken, string(reqStr))
	requestUrl := o.serverUrl + GetSignTaskEditUrl                     //接口请求api地址
	rspBody, requestId := common.SendPost(requestUrl, headMap, reqStr) //发送post请求
	err = json.Unmarshal(rspBody, &res)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}
	res.RequestId = requestId
	return res
}

// 获取签署任务预览链接
func (o *OpenApiClient) GetSignTaskPreviewUrl(req *signtaskRequestModel.GetSignTaskUrlReq, accessToken string) signtaskResponseModel.GetSignTaskPreviewUrlRes {
	var res signtaskResponseModel.GetSignTaskPreviewUrlRes
	reqStr, err := json.Marshal(req)
	if err != nil {
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken, string(reqStr))
	requestUrl := o.serverUrl + GetSignTaskPreviewUrl                  //接口请求api地址
	rspBody, requestId := common.SendPost(requestUrl, headMap, reqStr) //发送post请求
	err = json.Unmarshal(rspBody, &res)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}
	res.RequestId = requestId
	return res
}

// 获取签署任务批量签署链接
func (o *OpenApiClient) GetBatchSignUrl(req *signtaskRequestModel.GetBatchSignUrlReq, accessToken string) signtaskResponseModel.BatchSignUrlRes {
	var res signtaskResponseModel.BatchSignUrlRes
	reqStr, err := json.Marshal(req)
	if err != nil {
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken, string(reqStr))
	requestUrl := o.serverUrl + GetBatchSignUrl                        //接口请求api地址
	rspBody, requestId := common.SendPost(requestUrl, headMap, reqStr) //发送post请求
	err = json.Unmarshal(rspBody, &res)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}
	res.RequestId = requestId
	return res
}

// GetOwnerSignTaskListResponse 获取指定归属方的签署任务列表
func (o *OpenApiClient) GetOwnerSignTaskListResponse(ownerSignTaskListReq *signtaskRequestModel.GetOwnerSignTaskListReq, accessToken string) signtaskResponseModel.GetOwnerSignTaskListRes {
	var ownerSignTaskListRes signtaskResponseModel.GetOwnerSignTaskListRes
	reqStr, err := json.Marshal(ownerSignTaskListReq)
	if err != nil {
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken, string(reqStr))
	requestUrl := o.serverUrl + GetOwnerSignTaskListUrl                //接口请求api地址
	rspBody, requestId := common.SendPost(requestUrl, headMap, reqStr) //发送post请求
	err = json.Unmarshal(rspBody, &ownerSignTaskListRes)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}
	ownerSignTaskListRes.RequestId = requestId
	return ownerSignTaskListRes
}

// GetSignTaskDetailResponse 获取签署任务详情
func (o *OpenApiClient) GetSignTaskDetailResponse(signTaskDetailReq *signtaskRequestModel.GetSignTaskDetailReq, accessToken string) signtaskResponseModel.GetSignTaskDetailRes {
	var signTaskDetailRes signtaskResponseModel.GetSignTaskDetailRes
	reqStr, err := json.Marshal(signTaskDetailReq)
	if err != nil {
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken, string(reqStr))
	requestUrl := o.serverUrl + GetSignTaskDetailUrl                   //接口请求api地址
	rspBody, requestId := common.SendPost(requestUrl, headMap, reqStr) //发送post请求
	err = json.Unmarshal(rspBody, &signTaskDetailRes)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}
	signTaskDetailRes.RequestId = requestId
	return signTaskDetailRes
}

// GetOwnerSignTaskUrlResponse 获取指定归属方的签署任务文档下载地址
func (o *OpenApiClient) GetOwnerSignTaskUrlResponse(ownerSignTaskUrlReq *signtaskRequestModel.GetOwnerSignTaskUrlReq, accessToken string) signtaskResponseModel.GetSignTaskFieldUrlRes {
	var ownerSignTaskUrlRes signtaskResponseModel.GetSignTaskFieldUrlRes
	reqStr, err := json.Marshal(ownerSignTaskUrlReq)
	if err != nil {
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken, string(reqStr))
	requestUrl := o.serverUrl + GetOwnerSignTaskUrlUrl                 //接口请求api地址
	rspBody, requestId := common.SendPost(requestUrl, headMap, reqStr) //发送post请求
	err = json.Unmarshal(rspBody, &ownerSignTaskUrlRes)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}
	ownerSignTaskUrlRes.RequestId = requestId
	return ownerSignTaskUrlRes
}

// 查询企业签署任务文件夹
func (o *OpenApiClient) GetSignTaskCatalogList(req *signtaskRequestModel.SignTaskCatalogListReq, accessToken string) signtaskResponseModel.SignTaskCatalogListRes {
	var res signtaskResponseModel.SignTaskCatalogListRes
	reqStr, err := json.Marshal(req)
	if err != nil {
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken, string(reqStr))
	requestUrl := o.serverUrl + GetSignTaskCatalogListUrl              //接口请求api地址
	rspBody, requestId := common.SendPost(requestUrl, headMap, reqStr) //发送post请求
	err = json.Unmarshal(rspBody, &res)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}
	res.RequestId = requestId
	return res
}

// 查询签署任务控件信息
func (o *OpenApiClient) GetSignTaskFieldList(req *signtaskRequestModel.ListSignTaskFieldReq, accessToken string) signtaskResponseModel.ListSignTaskFieldRes {
	var res signtaskResponseModel.ListSignTaskFieldRes
	reqStr, err := json.Marshal(req)
	if err != nil {
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken, string(reqStr))
	requestUrl := o.serverUrl + GetSignTaskFieldListUrl                //接口请求api地址
	rspBody, requestId := common.SendPost(requestUrl, headMap, reqStr) //发送post请求
	err = json.Unmarshal(rspBody, &res)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}
	res.RequestId = requestId
	return res
}

// 查询签署任务参与方信息
func (o *OpenApiClient) GetSignTaskActorList(req *signtaskRequestModel.ListSignTaskActorReq, accessToken string) signtaskResponseModel.ListSignTaskActorRes {
	var res signtaskResponseModel.ListSignTaskActorRes
	reqStr, err := json.Marshal(req)
	if err != nil {
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken, string(reqStr))
	requestUrl := o.serverUrl + GetSignTaskActorListUrl                //接口请求api地址
	rspBody, requestId := common.SendPost(requestUrl, headMap, reqStr) //发送post请求
	err = json.Unmarshal(rspBody, &res)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}
	res.RequestId = requestId
	return res
}

// 查询参与方的签署刷脸底图
func (o *OpenApiClient) GetSignTaskActorFacePicture(req *signtaskRequestModel.GetActorFacePictureReq, accessToken string) signtaskResponseModel.GetActorFacePictureRes {
	var res signtaskResponseModel.GetActorFacePictureRes
	reqStr, err := json.Marshal(req)
	if err != nil {
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken, string(reqStr))
	requestUrl := o.serverUrl + GetSignTaskActorFacePictureUrl         //接口请求api地址
	rspBody, requestId := common.SendPost(requestUrl, headMap, reqStr) //发送post请求
	err = json.Unmarshal(rspBody, &res)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}
	res.RequestId = requestId
	return res
}

// 签署任务文档切图
func (o *OpenApiClient) GetSignTaskSlicingTicketId(req *signtaskRequestModel.GetSignTaskSlicingTicketIdReq, accessToken string) signtaskResponseModel.GetSignTaskSlicingTicketIdRes {
	var res signtaskResponseModel.GetSignTaskSlicingTicketIdRes
	reqStr, err := json.Marshal(req)
	if err != nil {
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken, string(reqStr))
	requestUrl := o.serverUrl + GetSignTaskSlicingTicketIdUrl          //接口请求api地址
	rspBody, requestId := common.SendPost(requestUrl, headMap, reqStr) //发送post请求
	err = json.Unmarshal(rspBody, &res)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}
	res.RequestId = requestId
	return res
}

// 获取图片版签署文档下载地址
func (o *OpenApiClient) GetSignTaskPicDownloadUrl(req *signtaskRequestModel.GetSignTaskPicDownloadUrlReq, accessToken string) signtaskResponseModel.GetSignTaskPicDownloadUrlRes {
	var res signtaskResponseModel.GetSignTaskPicDownloadUrlRes
	reqStr, err := json.Marshal(req)
	if err != nil {
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken, string(reqStr))
	requestUrl := o.serverUrl + GetSignTaskPicDownloadUrl              //接口请求api地址
	rspBody, requestId := common.SendPost(requestUrl, headMap, reqStr) //发送post请求
	err = json.Unmarshal(rspBody, &res)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}
	res.RequestId = requestId
	return res
}

// 查询签署任务审批信息
func (o *OpenApiClient) GetApprovalInfo(req *signtaskRequestModel.GetApprovalInfoReq, accessToken string) signtaskResponseModel.GetApprovalInfoRes {
	var res signtaskResponseModel.GetApprovalInfoRes
	reqStr, err := json.Marshal(req)
	if err != nil {
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken, string(reqStr))
	requestUrl := o.serverUrl + GetApprovalInfoUrl                     //接口请求api地址
	rspBody, requestId := common.SendPost(requestUrl, headMap, reqStr) //发送post请求
	err = json.Unmarshal(rspBody, &res)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}
	res.RequestId = requestId
	return res
}

// 获取签署任务公证处保全报告下载地址
func (o *OpenApiClient) GetSignTaskDownLoadEvidenceReport(req *signtaskRequestModel.GetEvidenceReportReq, accessToken string) signtaskResponseModel.GetEvidenceReportRes {
	var res signtaskResponseModel.GetEvidenceReportRes
	reqStr, err := json.Marshal(req)
	if err != nil {
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken, string(reqStr))
	requestUrl := o.serverUrl + GetSignTaskDownLoadEvidenceReportUrl   //接口请求api地址
	rspBody, requestId := common.SendPost(requestUrl, headMap, reqStr) //发送post请求
	err = json.Unmarshal(rspBody, &res)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}
	res.RequestId = requestId
	return res
}

// 查询签署业务类型列表
func (o *OpenApiClient) GetSignTaskBusinessTypeList(req *signtaskRequestModel.GetSignTaskBusinessTypeListReq, accessToken string) signtaskResponseModel.GetSignTaskBusinessTypeListRes {
	var res signtaskResponseModel.GetSignTaskBusinessTypeListRes
	reqStr, err := json.Marshal(req)
	if err != nil {
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken, string(reqStr))
	requestUrl := o.serverUrl + GetSignTaskBusinessTypeListUrl         //接口请求api地址
	rspBody, requestId := common.SendPost(requestUrl, headMap, reqStr) //发送post请求
	err = json.Unmarshal(rspBody, &res)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}
	res.RequestId = requestId
	return res
}
