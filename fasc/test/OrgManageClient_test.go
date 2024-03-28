package test

import (
	"fmt"
	"github.com/darwinOrg/go-sign-sdk/fasc/client"
	"github.com/darwinOrg/go-sign-sdk/fasc/model/requestModel/orgRequestModel"
	"strconv"
	"testing"
	"time"
)

// 获取组织管理链接
func TestGetCorpOrganizationManageUrl(t *testing.T) {
	c := client.NewClient(APPID, APPSECRET, SERVERURL)
	accessTokenRes := c.GetAuthToken()
	accessToken := accessTokenRes.Data.AccessToken

	req := &orgRequestModel.GetCorpOrgManageUrlReq{
		OpenCorpId:   OpenCorpId,
		ClientUserId: "clientUserId" + strconv.FormatInt(time.Now().Unix(), 10),
		Modules:      []string{"member", "connect"},
		RedirectUrl:  "http://www.baidu.com",
	}
	res := c.GetCorpOrganizationManageUrl(req, accessToken)
	fmt.Println(ModelToJsonString(res, false))
}

// 创建部门
func TestCreateCorpDept(t *testing.T) {
	c := client.NewClient(APPID, APPSECRET, SERVERURL)
	accessTokenRes := c.GetAuthToken()
	accessToken := accessTokenRes.Data.AccessToken

	var ParentDeptId int64
	req := &orgRequestModel.CreateCorpDeptReq{
		OpenCorpId:   OpenCorpId,
		ParentDeptId: &ParentDeptId,
		DeptName:     "",
		DeptOrderNum: 0,
	}
	res := c.CreateCorpDept(req, accessToken)
	fmt.Println(ModelToJsonString(res, false))
}

// 查询部门列表
func TestGetCorpDeptList(t *testing.T) {
	c := client.NewClient(APPID, APPSECRET, SERVERURL)
	accessTokenRes := c.GetAuthToken()
	accessToken := accessTokenRes.Data.AccessToken

	req := &orgRequestModel.GetCorpDeptListReq{
		OpenCorpId:   OpenCorpId,
		ParentDeptId: nil, //nil表示空
	}
	res := c.GetCorpDeptList(req, accessToken)
	fmt.Println(ModelToJsonString(res, false))
}

// 查询部门详情
func TestGetCorpDeptDetail(t *testing.T) {
	c := client.NewClient(APPID, APPSECRET, SERVERURL)
	accessTokenRes := c.GetAuthToken()
	accessToken := accessTokenRes.Data.AccessToken

	req := &orgRequestModel.GetCorpDeptDetailReq{
		OpenCorpId: OpenCorpId,
		DeptId:     0,
	}
	res := c.GetCorpDeptDetail(req, accessToken)
	fmt.Println(ModelToJsonString(res, false))
}

// 修改部门基本信息
func TestModifyCorpDeptInfo(t *testing.T) {
	c := client.NewClient(APPID, APPSECRET, SERVERURL)
	accessTokenRes := c.GetAuthToken()
	accessToken := accessTokenRes.Data.AccessToken

	req := &orgRequestModel.ModifyCorpDeptReq{
		OpenCorpId:   OpenCorpId,
		DeptId:       0,
		DeptName:     "",
		DeptOrderNum: 10000000,
	}
	res := c.ModifyCorpDeptInfo(req, accessToken)
	fmt.Println(ModelToJsonString(res, false))
}

// 删除部门
func TestDeleteCorpDept(t *testing.T) {
	c := client.NewClient(APPID, APPSECRET, SERVERURL)
	accessTokenRes := c.GetAuthToken()
	accessToken := accessTokenRes.Data.AccessToken

	req := &orgRequestModel.DeleteCorpDeptReq{
		OpenCorpId: OpenCorpId,
		DeptId:     0,
	}
	res := c.DeleteCorpDept(req, accessToken)
	fmt.Println(ModelToJsonString(res, false))
}

// 创建成员
func TestCreateCorpMember(t *testing.T) {
	c := client.NewClient(APPID, APPSECRET, SERVERURL)
	accessTokenRes := c.GetAuthToken()
	accessToken := accessTokenRes.Data.AccessToken

	infos := []orgRequestModel.EmployeeInfo{
		orgRequestModel.EmployeeInfo{
			MemberName:         "zfcTest003",
			InternalIdentifier: "zfcTest003",
			MemberEmail:        "zfcTest003@qq.com",
			MemberMobile:       "",
			MemberDeptIds:      []int64{},
		},
	}
	req := &orgRequestModel.CreateCorpMemberReq{
		OpenCorpId:          OpenCorpId,
		EmployeeInfos:       infos,
		NotifyActiveByEmail: false,
	}
	res := c.CreateCorpMember(req, accessToken)
	fmt.Println(ModelToJsonString(res, false))
}

// 获取成员激活链接
func TestGetCorpMemberActiveUrl(t *testing.T) {
	c := client.NewClient(APPID, APPSECRET, SERVERURL)
	accessTokenRes := c.GetAuthToken()
	accessToken := accessTokenRes.Data.AccessToken

	req := &orgRequestModel.GetCorpMemberActiveUrlReq{
		OpenCorpId: OpenCorpId,
		MemberIds:  []int64{},
	}
	res := c.GetCorpMemberActiveUrl(req, accessToken)
	fmt.Println(ModelToJsonString(res, false))
}

// 查询企业成员列表
func TestGetCorpMemberList(t *testing.T) {
	c := client.NewClient(APPID, APPSECRET, SERVERURL)
	accessTokenRes := c.GetAuthToken()
	accessToken := accessTokenRes.Data.AccessToken

	req := &orgRequestModel.GetCorpMemberListReq{
		OpenCorpId: OpenCorpId,
		ListFilter: &orgRequestModel.CorpListFilter{
			RoleType:   "",
			DeptId:     nil,
			FetchChild: true,
		},
		ListPageNo:   1,
		ListPageSize: 100,
	}
	res := c.GetCorpMemberList(req, accessToken)
	fmt.Println(ModelToJsonString(res, false))
}

// 查询成员详情
func TestGetCorpMemberDetail(t *testing.T) {
	c := client.NewClient(APPID, APPSECRET, SERVERURL)
	accessTokenRes := c.GetAuthToken()
	accessToken := accessTokenRes.Data.AccessToken

	req := &orgRequestModel.GetCorpMemberDetailReq{
		OpenCorpId: OpenCorpId,
		MemberId:   0,
	}
	res := c.GetCorpMemberDetail(req, accessToken)
	fmt.Println(ModelToJsonString(res, false))
}

// 修改成员基本信息
func TestModifyCorpMemberInfo(t *testing.T) {
	c := client.NewClient(APPID, APPSECRET, SERVERURL)
	accessTokenRes := c.GetAuthToken()
	accessToken := accessTokenRes.Data.AccessToken

	req := &orgRequestModel.ModifyCorpMemberReq{
		OpenCorpId:         OpenCorpId,
		MemberId:           0,
		MemberName:         "",
		InternalIdentifier: "",
		MemberEmail:        "",
		MemberMobile:       "",
		MemberDeptIds:      []int64{},
	}
	res := c.ModifyCorpMemberInfo(req, accessToken)
	fmt.Println(ModelToJsonString(res, false))
}

// 设置成员所属部门
func TestSetCorpMemberDept(t *testing.T) {
	c := client.NewClient(APPID, APPSECRET, SERVERURL)
	accessTokenRes := c.GetAuthToken()
	accessToken := accessTokenRes.Data.AccessToken

	req := &orgRequestModel.SetCorpMemberDeptReq{
		OpenCorpId:    OpenCorpId,
		MemberIds:     []int64{},
		MemberDeptIds: []int64{},
		Model:         "cover",
	}
	res := c.SetCorpMemberDept(req, accessToken)
	fmt.Println(ModelToJsonString(res, false))
}

// 设置成员状态
func TestSetCorpMemberStatus(t *testing.T) {
	c := client.NewClient(APPID, APPSECRET, SERVERURL)
	accessTokenRes := c.GetAuthToken()
	accessToken := accessTokenRes.Data.AccessToken

	req := &orgRequestModel.SetCorpMemberStatusReq{
		OpenCorpId:   OpenCorpId,
		MemberIds:    []int64{},
		MemberStatus: "enable",
	}
	res := c.SetCorpMemberStatus(req, accessToken)
	fmt.Println(ModelToJsonString(res, false))
}

// 删除成员
func TestDeleteCorpMember(t *testing.T) {
	c := client.NewClient(APPID, APPSECRET, SERVERURL)
	accessTokenRes := c.GetAuthToken()
	accessToken := accessTokenRes.Data.AccessToken

	req := &orgRequestModel.DeleteCorpMemberReq{
		OpenCorpId: OpenCorpId,
		MemberIds:  []int64{},
	}
	res := c.DeleteCorpMember(req, accessToken)
	fmt.Println(ModelToJsonString(res, false))
}

// 查询相对方列表
func TestGetCounterPartList(t *testing.T) {
	c := client.NewClient(APPID, APPSECRET, SERVERURL)
	accessTokenRes := c.GetAuthToken()
	accessToken := accessTokenRes.Data.AccessToken

	req := &orgRequestModel.GetCounterPartListReq{OpenCorpId: OpenCorpId}
	res := c.GetCounterPartList(req, accessToken)
	fmt.Println(ModelToJsonString(res, false))
}
