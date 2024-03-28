package client

import (
	"encoding/json"
	"fmt"
	"github.com/darwinOrg/go-sign-sdk/fasc/common"
	"github.com/darwinOrg/go-sign-sdk/fasc/model/requestModel/orgRequestModel"
	"github.com/darwinOrg/go-sign-sdk/fasc/model/responseModel/orgResponseModel"
)

const (
	CorpOrganizationManageUrl = "/corp/organization/manage/get-url" //获取组织管理链接
	CorpDeptCreateUrl         = "/corp/dept/create"                 //创建部门
	CorpDeptGetListUrl        = "/corp/dept/get-list"               //查询部门列表
	CorpDeptGetDetailUrl      = "/corp/dept/get-detail"             //查询部门详情
	CorpDeptModifyUrl         = "/corp/dept/modify"                 //修改部门基本信息
	CorpDeptDeleteUrl         = "/corp/dept/delete"                 //删除部门
	CorpMemberCreateUrl       = "/corp/member/create"               //创建成员
	CorpMemberGetActiveUrl    = "/corp/member/get-active-url"       //获取成员激活链接
	CorpMemberGetListUrl      = "/corp/member/get-list"             //查询企业成员列表
	CorpMemberGetDetailUrl    = "/corp/member/get-detail"           //查询成员详情
	CorpMemberModifyUrl       = "/corp/member/modify"               //修改成员基本信息
	CorpMemberSetDeptUrl      = "/corp/member/set-dept"             //设置成员所属部门
	CorpMemberSetStatusUrl    = "/corp/member/set-status"           //设置成员状态
	CorpMemberDeleteUrl       = "/corp/member/delete"               //删除成员
	GetCounterPartListUrl     = "/counterpart/get-list"             //查询相对方列表
)

// 获取组织管理链接
func (o *OpenApiClient) GetCorpOrganizationManageUrl(getCorpOrgManageUrlReq *orgRequestModel.GetCorpOrgManageUrlReq, accessToken string) orgResponseModel.GetCorpOrgManageUrlRes {
	var GetCorpOrgManageUrlRes orgResponseModel.GetCorpOrgManageUrlRes
	reqStr, err := json.Marshal(getCorpOrgManageUrlReq)
	if err != nil {
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken, string(reqStr))
	requestUrl := o.serverUrl + CorpOrganizationManageUrl              //接口请求api地址
	rspBody, requestId := common.SendPost(requestUrl, headMap, reqStr) //发送post请求
	err = json.Unmarshal(rspBody, &GetCorpOrgManageUrlRes)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}
	GetCorpOrgManageUrlRes.RequestId = requestId
	return GetCorpOrgManageUrlRes
}

// 创建部门
func (o *OpenApiClient) CreateCorpDept(req *orgRequestModel.CreateCorpDeptReq, accessToken string) orgResponseModel.CreateCorpDeptRes {
	var res orgResponseModel.CreateCorpDeptRes
	reqStr, err := json.Marshal(req)
	if err != nil {
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken, string(reqStr))
	requestUrl := o.serverUrl + CorpDeptCreateUrl                      //接口请求api地址
	rspBody, requestId := common.SendPost(requestUrl, headMap, reqStr) //发送post请求
	err = json.Unmarshal(rspBody, &res)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}
	res.RequestId = requestId
	return res
}

// 查询部门列表
func (o *OpenApiClient) GetCorpDeptList(req *orgRequestModel.GetCorpDeptListReq, accessToken string) orgResponseModel.GetCorpDeptListRes {
	var res orgResponseModel.GetCorpDeptListRes
	reqStr, err := json.Marshal(req)
	if err != nil {
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken, string(reqStr))
	requestUrl := o.serverUrl + CorpDeptGetListUrl                     //接口请求api地址
	rspBody, requestId := common.SendPost(requestUrl, headMap, reqStr) //发送post请求
	err = json.Unmarshal(rspBody, &res)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}
	res.RequestId = requestId
	return res
}

// 查询部门详情
func (o *OpenApiClient) GetCorpDeptDetail(req *orgRequestModel.GetCorpDeptDetailReq, accessToken string) orgResponseModel.GetCorpDeptDetailRes {
	var res orgResponseModel.GetCorpDeptDetailRes
	reqStr, err := json.Marshal(req)
	if err != nil {
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken, string(reqStr))
	requestUrl := o.serverUrl + CorpDeptGetDetailUrl                   //接口请求api地址
	rspBody, requestId := common.SendPost(requestUrl, headMap, reqStr) //发送post请求
	err = json.Unmarshal(rspBody, &res)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}
	res.RequestId = requestId
	return res
}

// 修改部门基本信息
func (o *OpenApiClient) ModifyCorpDeptInfo(req *orgRequestModel.ModifyCorpDeptReq, accessToken string) orgResponseModel.ModifyCorpDeptRes {
	var res orgResponseModel.ModifyCorpDeptRes
	reqStr, err := json.Marshal(req)
	if err != nil {
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken, string(reqStr))
	requestUrl := o.serverUrl + CorpDeptModifyUrl                      //接口请求api地址
	rspBody, requestId := common.SendPost(requestUrl, headMap, reqStr) //发送post请求
	err = json.Unmarshal(rspBody, &res)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}
	res.RequestId = requestId
	return res
}

// 删除部门
func (o *OpenApiClient) DeleteCorpDept(req *orgRequestModel.DeleteCorpDeptReq, accessToken string) orgResponseModel.DeleteCorpDeptRes {
	var res orgResponseModel.DeleteCorpDeptRes
	reqStr, err := json.Marshal(req)
	if err != nil {
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken, string(reqStr))
	requestUrl := o.serverUrl + CorpDeptDeleteUrl                      //接口请求api地址
	rspBody, requestId := common.SendPost(requestUrl, headMap, reqStr) //发送post请求
	err = json.Unmarshal(rspBody, &res)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}
	res.RequestId = requestId
	return res
}

// 创建成员
func (o *OpenApiClient) CreateCorpMember(req *orgRequestModel.CreateCorpMemberReq, accessToken string) orgResponseModel.CreateCorpMemberRes {
	var res orgResponseModel.CreateCorpMemberRes
	reqStr, err := json.Marshal(req)
	if err != nil {
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken, string(reqStr))
	requestUrl := o.serverUrl + CorpMemberCreateUrl                    //接口请求api地址
	rspBody, requestId := common.SendPost(requestUrl, headMap, reqStr) //发送post请求
	err = json.Unmarshal(rspBody, &res)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}
	res.RequestId = requestId
	return res
}

// 获取成员激活链接
func (o *OpenApiClient) GetCorpMemberActiveUrl(req *orgRequestModel.GetCorpMemberActiveUrlReq, accessToken string) orgResponseModel.GetCorpMemberActiveUrlRes {
	var res orgResponseModel.GetCorpMemberActiveUrlRes
	reqStr, err := json.Marshal(req)
	if err != nil {
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken, string(reqStr))
	requestUrl := o.serverUrl + CorpMemberGetActiveUrl                 //接口请求api地址
	rspBody, requestId := common.SendPost(requestUrl, headMap, reqStr) //发送post请求
	err = json.Unmarshal(rspBody, &res)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}
	res.RequestId = requestId
	return res
}

// GetCorpMemberListResponse 查询企业成员列表
func (o *OpenApiClient) GetCorpMemberList(corpMemberListReq *orgRequestModel.GetCorpMemberListReq, accessToken string) orgResponseModel.GetCorpMemberListRes {
	var CorpMemberListRes orgResponseModel.GetCorpMemberListRes
	reqStr, err := json.Marshal(corpMemberListReq)
	if err != nil {
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken, string(reqStr))
	requestUrl := o.serverUrl + CorpMemberGetListUrl                   //接口请求api地址
	rspBody, requestId := common.SendPost(requestUrl, headMap, reqStr) //发送post请求
	err = json.Unmarshal(rspBody, &CorpMemberListRes)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}
	CorpMemberListRes.RequestId = requestId
	return CorpMemberListRes
}

// 查询成员详情
func (o *OpenApiClient) GetCorpMemberDetail(req *orgRequestModel.GetCorpMemberDetailReq, accessToken string) orgResponseModel.GetCorpMemberDetailRes {
	var res orgResponseModel.GetCorpMemberDetailRes
	reqStr, err := json.Marshal(req)
	if err != nil {
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken, string(reqStr))
	requestUrl := o.serverUrl + CorpMemberGetDetailUrl                 //接口请求api地址
	rspBody, requestId := common.SendPost(requestUrl, headMap, reqStr) //发送post请求
	err = json.Unmarshal(rspBody, &res)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}
	res.RequestId = requestId
	return res
}

// 修改成员基本信息
func (o *OpenApiClient) ModifyCorpMemberInfo(req *orgRequestModel.ModifyCorpMemberReq, accessToken string) orgResponseModel.ModifyCorpMemberRes {
	var res orgResponseModel.ModifyCorpMemberRes
	reqStr, err := json.Marshal(req)
	if err != nil {
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken, string(reqStr))
	requestUrl := o.serverUrl + CorpMemberModifyUrl                    //接口请求api地址
	rspBody, requestId := common.SendPost(requestUrl, headMap, reqStr) //发送post请求
	err = json.Unmarshal(rspBody, &res)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}
	res.RequestId = requestId
	return res
}

// 设置成员所属部门
func (o *OpenApiClient) SetCorpMemberDept(req *orgRequestModel.SetCorpMemberDeptReq, accessToken string) orgResponseModel.SetCorpMemberDeptRes {
	var res orgResponseModel.SetCorpMemberDeptRes
	reqStr, err := json.Marshal(req)
	if err != nil {
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken, string(reqStr))
	requestUrl := o.serverUrl + CorpMemberSetDeptUrl                   //接口请求api地址
	rspBody, requestId := common.SendPost(requestUrl, headMap, reqStr) //发送post请求
	err = json.Unmarshal(rspBody, &res)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}
	res.RequestId = requestId
	return res
}

// 设置成员状态
func (o *OpenApiClient) SetCorpMemberStatus(req *orgRequestModel.SetCorpMemberStatusReq, accessToken string) orgResponseModel.SetCorpMemberStatusRes {
	var res orgResponseModel.SetCorpMemberStatusRes
	reqStr, err := json.Marshal(req)
	if err != nil {
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken, string(reqStr))
	requestUrl := o.serverUrl + CorpMemberSetStatusUrl                 //接口请求api地址
	rspBody, requestId := common.SendPost(requestUrl, headMap, reqStr) //发送post请求
	err = json.Unmarshal(rspBody, &res)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}
	res.RequestId = requestId
	return res
}

// 删除成员
func (o *OpenApiClient) DeleteCorpMember(req *orgRequestModel.DeleteCorpMemberReq, accessToken string) orgResponseModel.DeleteCorpMemberRes {
	var res orgResponseModel.DeleteCorpMemberRes
	reqStr, err := json.Marshal(req)
	if err != nil {
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken, string(reqStr))
	requestUrl := o.serverUrl + CorpMemberDeleteUrl                    //接口请求api地址
	rspBody, requestId := common.SendPost(requestUrl, headMap, reqStr) //发送post请求
	err = json.Unmarshal(rspBody, &res)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}
	res.RequestId = requestId
	return res
}

// 查询相对方列表
func (o *OpenApiClient) GetCounterPartList(req *orgRequestModel.GetCounterPartListReq, accessToken string) orgResponseModel.GetCounterPartListRes {
	var res orgResponseModel.GetCounterPartListRes
	reqStr, err := json.Marshal(req)
	if err != nil {
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken, string(reqStr))
	requestUrl := o.serverUrl + GetCounterPartListUrl                  //接口请求api地址
	rspBody, requestId := common.SendPost(requestUrl, headMap, reqStr) //发送post请求
	err = json.Unmarshal(rspBody, &res)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}
	res.RequestId = requestId
	return res
}
