package test

import (
	"fmt"
	"github.com/darwinOrg/go-sign-sdk/fasc/client"
	"github.com/darwinOrg/go-sign-sdk/fasc/model/commonModel"
	"github.com/darwinOrg/go-sign-sdk/fasc/model/requestModel/signtaskRequestModel"
	"strconv"
	"testing"
	"time"
)

/*
*
创建签署任务:
*/
func TestCreateSignTask(t *testing.T) {
	c := client.NewClient(APPID, APPSECRET, SERVERURL)
	accessTokenRes := c.GetAuthToken()
	accessToken := accessTokenRes.Data.AccessToken
	req := signtaskRequestModel.CreateSignTaskReq{
		SignTaskSubject: "签署任务主题" + strconv.FormatInt(time.Now().Unix(), 10),
		Initiator: &commonModel.OpenId{
			IdType: "corp",
			OpenId: OpenCorpId,
		},
		SignDocType:      "",
		ExpiresTime:      "",
		AutoStart:        true,
		AutoFillFinalize: true,
		AutoFinish:       true,
		BusinessTypeId:   nil,
		SignInOrder:      false,
		CertCAOrg:        "",
		BusinessScene:    nil,
		Docs: []commonModel.Doc{
			commonModel.Doc{
				DocId:         "doc1",
				DocName:       "doc1",
				DocFileId:     DocFileId,
				DocTemplateId: "",
				DocFields: []commonModel.Field{
					commonModel.Field{
						FieldId:   "FieldId1",
						FieldName: "FieldId1",
						Position: &commonModel.FieldPosition{
							PositionMode:    "pixel",
							PositionPageNo:  1,
							PositionX:       "200",
							PositionY:       "200",
							PositionKeyword: "",
						},
						FieldType: "corp_seal",
					},
					commonModel.Field{
						FieldId:   "FieldId2",
						FieldName: "FieldId2",
						Position: &commonModel.FieldPosition{
							PositionMode:    "pixel",
							PositionPageNo:  1,
							PositionX:       "100",
							PositionY:       "100",
							PositionKeyword: "",
						},
						FieldType: "person_sign",
					},
				},
			},
		},
		Attachs: nil,
		Actors: []commonModel.SignTaskActor{
			commonModel.SignTaskActor{
				Actor: &commonModel.Actor{
					ActorType:         "corp",
					ActorId:           "企业方",
					ActorName:         "企业方",
					Permissions:       []string{"sign"},
					ActorOpenId:       OpenCorpId,
					ActorFDDId:        "",
					IdentNameForMatch: "",
					CertNoForMatch:    "",
					ActorCorpMember:   nil,
					Notification:      nil,
				},
				FillFields: nil,
				SignFields: []commonModel.SignField{
					commonModel.SignField{
						FieldDocId: "doc1",
						FieldId:    "FieldId1",
						FieldName:  "FieldId1",
						SealId:     nil,
						Moveable:   false,
					},
				},
				SignConfigInfo: nil,
			},
			commonModel.SignTaskActor{
				Actor: &commonModel.Actor{
					ActorType:         "person",
					ActorId:           "个人方",
					ActorName:         "个人方",
					Permissions:       []string{"sign"},
					ActorOpenId:       OpenUserId,
					ActorFDDId:        "",
					IdentNameForMatch: "",
					CertNoForMatch:    "",
					ActorCorpMember:   nil,
					Notification:      nil,
				},
				FillFields: nil,
				SignFields: []commonModel.SignField{
					commonModel.SignField{
						FieldDocId: "doc1",
						FieldId:    "FieldId2",
						FieldName:  "FieldId2",
						SealId:     nil,
						Moveable:   false,
					},
				},
				SignConfigInfo: &commonModel.SignConfigInfo{
					OrderNo:           0,
					RequestMemberSign: false,
					VerifyMethods:     nil,
					ReadingToEnd:      false,
					ReadingTime:       "",
					BlockHere:         false,
					RequestVerifyFree: false,
					SignerSignMethod:  "",
					JoinByLink:        true,
					IdentifiedView:    false,
				},
			},
		},
	}
	res := c.CreateSignTask(&req, accessToken)
	fmt.Println(ModelToJsonString(res, false))
}

// 创建签署任务基于签署模板
func TestCreateSignTaskWithTemplate(t *testing.T) {
	c := client.NewClient(APPID, APPSECRET, SERVERURL)
	accessTokenRes := c.GetAuthToken()
	accessToken := accessTokenRes.Data.AccessToken
	req := &signtaskRequestModel.CreateSignTaskWithTemplateReq{
		SignTaskSubject: "签署任务主题" + strconv.FormatInt(time.Now().Unix(), 10),
		Initiator: &commonModel.OpenId{
			IdType: "corp",
			OpenId: OpenCorpId,
		},
		SignDocType:      "",
		SignTemplateId:   SignTemplateId,
		ExpiresTime:      "",
		AutoStart:        true,
		AutoFillFinalize: true,
		CertCAOrg:        "",
		BusinessScene:    nil,
		Actors: []commonModel.SignTaskActor{
			commonModel.SignTaskActor{
				Actor: &commonModel.Actor{
					ActorType:         "corp",
					ActorId:           "企业方",
					ActorName:         "企业方",
					Permissions:       []string{"sign"},
					ActorOpenId:       OpenCorpId,
					ActorFDDId:        "",
					IdentNameForMatch: "",
					CertNoForMatch:    "",
					ActorCorpMember:   nil,
					Notification:      nil,
				},
				FillFields: nil,
				SignFields: nil,
				SignConfigInfo: &commonModel.SignConfigInfo{
					OrderNo:           0,
					RequestMemberSign: false,
					VerifyMethods:     nil,
					ReadingToEnd:      false,
					ReadingTime:       "",
					BlockHere:         false,
					RequestVerifyFree: false,
					SignerSignMethod:  "",
					JoinByLink:        false,
					IdentifiedView:    false,
				},
			},
			commonModel.SignTaskActor{
				Actor: &commonModel.Actor{
					ActorType:         "person",
					ActorId:           "个人方",
					ActorName:         "个人方",
					Permissions:       []string{"sign"},
					ActorOpenId:       OpenUserId,
					ActorFDDId:        "",
					IdentNameForMatch: "",
					CertNoForMatch:    "",
					ActorCorpMember:   nil,
					Notification:      nil,
				},
				FillFields: nil,
				SignFields: nil,
				SignConfigInfo: &commonModel.SignConfigInfo{
					OrderNo:           0,
					RequestMemberSign: false,
					VerifyMethods:     nil,
					ReadingToEnd:      false,
					ReadingTime:       "",
					BlockHere:         false,
					RequestVerifyFree: false,
					SignerSignMethod:  "",
					JoinByLink:        false,
					IdentifiedView:    false,
				},
			},
		},
	}
	res := c.CreateSignTaskWithTemplate(req, accessToken)
	fmt.Println(ModelToJsonString(res, false))
}

// 添加签署任务文档
func TestAddSignTaskDoc(t *testing.T) {
	c := client.NewClient(APPID, APPSECRET, SERVERURL)
	accessTokenRes := c.GetAuthToken()
	accessToken := accessTokenRes.Data.AccessToken
	req := &signtaskRequestModel.AddSignTaskDocReq{
		SignTaskId: SignTaskId,
		Docs: []commonModel.Doc{
			commonModel.Doc{
				DocId:         "doc2",
				DocName:       "doc2",
				DocFileId:     DocFileId,
				DocTemplateId: "",
				DocFields: []commonModel.Field{
					commonModel.Field{
						FieldId:   "FieldId3",
						FieldName: "FieldId3",
						Position: &commonModel.FieldPosition{
							PositionMode:   "pixel",
							PositionPageNo: 1,
							PositionX:      "200",
							PositionY:      "200",
						},
						FieldType: "date_sign",
					},
				},
			},
		},
	}
	res := c.AddSignTaskDoc(req, accessToken)
	fmt.Println(ModelToJsonString(res, false))
}

// 移除签署任务文档
func TestDeleteSignTaskDoc(t *testing.T) {
	c := client.NewClient(APPID, APPSECRET, SERVERURL)
	accessTokenRes := c.GetAuthToken()
	accessToken := accessTokenRes.Data.AccessToken
	req := &signtaskRequestModel.DeleteSignTaskDocReq{
		SignTaskId: SignTaskId,
		DocIds: []string{
			"doc2",
		},
	}
	res := c.DeleteSignTaskDoc(req, accessToken)
	fmt.Println(ModelToJsonString(res, false))
}

// 添加签署任务控件
func TestAddSignTaskField(t *testing.T) {
	c := client.NewClient(APPID, APPSECRET, SERVERURL)
	accessTokenRes := c.GetAuthToken()
	accessToken := accessTokenRes.Data.AccessToken
	req := &signtaskRequestModel.AddSignTaskFieldReq{
		SignTaskId: SignTaskId,
		ActorId:    "个人方",
		Fields: []signtaskRequestModel.AddSignTaskField{
			signtaskRequestModel.AddSignTaskField{
				DocId: "doc1",
				DocFields: []commonModel.Field{
					commonModel.Field{
						FieldId:   "FieldId4",
						FieldName: "FieldId4",
						Position: &commonModel.FieldPosition{
							PositionMode:   "pixel",
							PositionPageNo: 1,
							PositionX:      "300",
							PositionY:      "300",
						},
						FieldType: "date_sign",
					},
				},
			},
		},
	}
	res := c.AddSignTaskField(req, accessToken)
	fmt.Println(ModelToJsonString(res, false))
}

// 移除签署任务控件
func TestDeleteSignTaskField(t *testing.T) {
	c := client.NewClient(APPID, APPSECRET, SERVERURL)
	accessTokenRes := c.GetAuthToken()
	accessToken := accessTokenRes.Data.AccessToken
	req := &signtaskRequestModel.DeleteSignTaskFieldReq{
		SignTaskId: SignTaskId,
		Fields: []signtaskRequestModel.DeleteSignTaskField{
			signtaskRequestModel.DeleteSignTaskField{
				DocId: "doc2",
				DocFields: []string{
					"FieldId3", "FieldId4",
				},
			},
		},
	}
	res := c.DeleteSignTaskField(req, accessToken)
	fmt.Println(ModelToJsonString(res, false))
}

// 添加签署任务附件
func TestAddSignTaskAttach(t *testing.T) {
	c := client.NewClient(APPID, APPSECRET, SERVERURL)
	accessTokenRes := c.GetAuthToken()
	accessToken := accessTokenRes.Data.AccessToken
	req := &signtaskRequestModel.AddSignTaskAttachReq{
		SignTaskId: SignTaskId,
		Attachs: []commonModel.Attach{
			commonModel.Attach{
				AttachId:     "附件1",
				AttachName:   "附件1",
				AttachFileId: AttachFileId,
			},
			commonModel.Attach{
				AttachId:     "附件2",
				AttachName:   "附件2",
				AttachFileId: AttachFileId,
			},
		},
	}
	res := c.AddSignTaskAttach(req, accessToken)
	fmt.Println(ModelToJsonString(res, false))
}

// 移除签署任务附件
func TestDeleteSignTaskAttach(t *testing.T) {
	c := client.NewClient(APPID, APPSECRET, SERVERURL)
	accessTokenRes := c.GetAuthToken()
	accessToken := accessTokenRes.Data.AccessToken
	req := &signtaskRequestModel.DeleteSignTaskAttachReq{
		SignTaskId: SignTaskId,
		AttachIds: []string{
			"附件1", "附件2",
		},
	}
	res := c.DeleteSignTaskAttach(req, accessToken)
	fmt.Println(ModelToJsonString(res, false))
}

// 添加签署任务参与方
func TestAddSignTaskActor(t *testing.T) {
	c := client.NewClient(APPID, APPSECRET, SERVERURL)
	accessTokenRes := c.GetAuthToken()
	accessToken := accessTokenRes.Data.AccessToken
	req := &signtaskRequestModel.AddSignTaskActorReq{
		SignTaskId: SignTaskId,
		Actors: []commonModel.SignTaskActor{
			commonModel.SignTaskActor{
				Actor: &commonModel.Actor{
					ActorType:         "corp",
					ActorId:           "ActorId3",
					ActorName:         "ActorId3",
					Permissions:       []string{"sign"},
					ActorOpenId:       "",
					ActorFDDId:        "",
					IdentNameForMatch: "",
					CertNoForMatch:    "",
					ActorCorpMember:   nil,
					Notification:      nil,
				},
				FillFields:     nil,
				SignFields:     nil,
				SignConfigInfo: nil,
			},
		},
	}
	res := c.AddSignTaskActor(req, accessToken)
	fmt.Println(ModelToJsonString(res, false))
}

// 移除签署任务参与方
func TestDeleteSignTaskActor(t *testing.T) {
	c := client.NewClient(APPID, APPSECRET, SERVERURL)
	accessTokenRes := c.GetAuthToken()
	accessToken := accessTokenRes.Data.AccessToken
	req := &signtaskRequestModel.DeleteSignTaskActorReq{
		SignTaskId: SignTaskId,
		ActorIds: []string{
			"ActorId3", "ActorId4",
		},
	}
	res := c.DeleteSignTaskActor(req, accessToken)
	fmt.Println(ModelToJsonString(res, false))
}

// 填写签署任务控件内容
func TestFillSignTaskFieldValues(t *testing.T) {
	c := client.NewClient(APPID, APPSECRET, SERVERURL)
	accessTokenRes := c.GetAuthToken()
	accessToken := accessTokenRes.Data.AccessToken
	req := &signtaskRequestModel.FillFieldValuesReq{
		SignTaskId: SignTaskId,
		DocFieldValues: []signtaskRequestModel.DocFieldValue{
			signtaskRequestModel.DocFieldValue{
				DocId:      "doc1",
				FieldId:    "5635387223",
				FieldValue: "填写内容啊啊",
			},
		},
	}
	res := c.FillSignTaskFieldValues(req, accessToken)
	fmt.Println(ModelToJsonString(res, false))
}

// 提交签署任务
func TestStartSignTask(t *testing.T) {
	c := client.NewClient(APPID, APPSECRET, SERVERURL)
	accessTokenRes := c.GetAuthToken()
	accessToken := accessTokenRes.Data.AccessToken
	req := &signtaskRequestModel.StartSignTaskReq{SignTaskId: SignTaskId}
	res := c.StartSignTask(req, accessToken)
	fmt.Println(ModelToJsonString(res, false))
}

// 获取签署任务参与方专属链接
func TestGetSignTaskActorUrl(t *testing.T) {
	c := client.NewClient(APPID, APPSECRET, SERVERURL)
	accessTokenRes := c.GetAuthToken()
	accessToken := accessTokenRes.Data.AccessToken
	req := &signtaskRequestModel.SignTaskActorGetUrlReq{
		SignTaskId:   SignTaskId,
		ActorId:      "个人方",
		ClientUserId: "",
		RedirectUrl:  "",
	}
	res := c.GetSignTaskActorUrl(req, accessToken)
	fmt.Println(ModelToJsonString(res, false))
}

// 定稿签署任务文档
func TestFinalizeSignTaskDoc(t *testing.T) {
	c := client.NewClient(APPID, APPSECRET, SERVERURL)
	accessTokenRes := c.GetAuthToken()
	accessToken := accessTokenRes.Data.AccessToken
	req := &signtaskRequestModel.FinalizeSignTaskReq{
		SignTaskId: SignTaskId,
	}
	res := c.FinalizeSignTaskDoc(req, accessToken)
	fmt.Println(ModelToJsonString(res, false))
}

// 删除签署任务
func TestDeleteSignTask(t *testing.T) {
	c := client.NewClient(APPID, APPSECRET, SERVERURL)
	accessTokenRes := c.GetAuthToken()
	accessToken := accessTokenRes.Data.AccessToken
	req := &signtaskRequestModel.DeleteSignTaskReq{
		SignTaskId: SignTaskId,
	}
	res := c.DeleteSignTask(req, accessToken)
	fmt.Println(ModelToJsonString(res, false))
}

// 作废签署任务
func TestAbolishSignTask(t *testing.T) {
	c := client.NewClient(APPID, APPSECRET, SERVERURL)
	accessTokenRes := c.GetAuthToken()
	accessToken := accessTokenRes.Data.AccessToken
	req := &signtaskRequestModel.AbolishSignTaskReq{
		SignTaskId: "1681992137774158285",
		AbolishedInitiator: &signtaskRequestModel.AbolishedInitiator{
			InitiatorId: OpenCorpId,
		},
		DocSource: "platform",
		Reason:    "hhhhh",
	}
	res := c.AbolishSignTask(req, accessToken)
	fmt.Println(ModelToJsonString(res, false))
}

// 结束签署任务
func TestFinishSignTask(t *testing.T) {
	c := client.NewClient(APPID, APPSECRET, SERVERURL)
	accessTokenRes := c.GetAuthToken()
	accessToken := accessTokenRes.Data.AccessToken
	req := &signtaskRequestModel.FinishSignTaskReq{SignTaskId: SignTaskId}
	res := c.FinishSignTask(req, accessToken)
	fmt.Println(ModelToJsonString(res, false))
}

// 催办签署任务
func TestUrgeSignTask(t *testing.T) {
	c := client.NewClient(APPID, APPSECRET, SERVERURL)
	accessTokenRes := c.GetAuthToken()
	accessToken := accessTokenRes.Data.AccessToken
	req := &signtaskRequestModel.SignTaskUrgeReq{
		SignTaskId: SignTaskId,
	}
	res := c.UrgeSignTask(req, accessToken)
	fmt.Println(ModelToJsonString(res, false))
}

// 阻塞签署任务
func TestBlockSignTask(t *testing.T) {
	c := client.NewClient(APPID, APPSECRET, SERVERURL)
	accessTokenRes := c.GetAuthToken()
	accessToken := accessTokenRes.Data.AccessToken
	req := &signtaskRequestModel.BlockSignTaskReq{
		SignTaskId: SignTaskId,
		ActorId:    "企业方",
	}
	res := c.BlockSignTask(req, accessToken)
	fmt.Println(ModelToJsonString(res, false))
}

// 解阻签署任务
func TestUnBlockSignTask(t *testing.T) {
	c := client.NewClient(APPID, APPSECRET, SERVERURL)
	accessTokenRes := c.GetAuthToken()
	accessToken := accessTokenRes.Data.AccessToken
	req := &signtaskRequestModel.UnBlockSignTaskReq{
		SignTaskId: SignTaskId,
		ActorId:    "企业方",
	}
	res := c.UnBlockSignTask(req, accessToken)
	fmt.Println(ModelToJsonString(res, false))
}

// 撤销签署任务
func TestCancelSignTask(t *testing.T) {
	c := client.NewClient(APPID, APPSECRET, SERVERURL)
	accessTokenRes := c.GetAuthToken()
	accessToken := accessTokenRes.Data.AccessToken
	req := &signtaskRequestModel.CancelSignTaskReq{
		SignTaskId: SignTaskId,
	}
	res := c.CancelSignTask(req, accessToken)
	fmt.Println(ModelToJsonString(res, false))
}

// 获取签署任务编辑链接
func TestGetSignTaskEditUrl(t *testing.T) {
	c := client.NewClient(APPID, APPSECRET, SERVERURL)
	accessTokenRes := c.GetAuthToken()
	accessToken := accessTokenRes.Data.AccessToken
	req := &signtaskRequestModel.GetSignTaskUrlReq{
		SignTaskId:  SignTaskId,
		RedirectUrl: "",
	}
	res := c.GetSignTaskEditUrl(req, accessToken)
	fmt.Println(ModelToJsonString(res, false))
}

// 获取签署任务预览链接
func TestGetSignTaskPreviewUrl(t *testing.T) {
	c := client.NewClient(APPID, APPSECRET, SERVERURL)
	accessTokenRes := c.GetAuthToken()
	accessToken := accessTokenRes.Data.AccessToken
	req := &signtaskRequestModel.GetSignTaskUrlReq{
		SignTaskId: SignTaskId,
	}
	res := c.GetSignTaskPreviewUrl(req, accessToken)
	fmt.Println(ModelToJsonString(res, false))
}

// 获取签署任务批量签署链接
func TestGetBatchSignUrl(t *testing.T) {
	c := client.NewClient(APPID, APPSECRET, SERVERURL)
	accessTokenRes := c.GetAuthToken()
	accessToken := accessTokenRes.Data.AccessToken
	memberId1 := int64(1660639745225159620)
	req := &signtaskRequestModel.GetBatchSignUrlReq{
		OwnerId: &commonModel.OpenId{
			IdType: "corp",
			OpenId: OpenCorpId,
		},
		MemberId:     &memberId1,
		ClientUserId: "",
		SignTaskIds: []string{
			"SignTaskId1", "SignTaskId2",
		},
		RedirectUrl: "",
	}
	res := c.GetBatchSignUrl(req, accessToken)
	fmt.Println(ModelToJsonString(res, false))
}

// 获取指定归属方的签署任务列表
func TestGetOwnerSignTaskListResponse(t *testing.T) {
	c := client.NewClient(APPID, APPSECRET, SERVERURL)
	accessTokenRes := c.GetAuthToken()
	accessToken := accessTokenRes.Data.AccessToken
	req := &signtaskRequestModel.GetOwnerSignTaskListReq{
		OwnerId: &commonModel.OpenId{
			IdType: "corp",
			OpenId: OpenCorpId,
		},
		OwnerRole: "actor",
		ListFilter: &signtaskRequestModel.SignTaskListFilter{
			SignTaskSubject: "签署",
			SignTaskStatus: []string{
				"sign_progress",
			},
		},
		ListPageNo:   1,
		ListPageSize: 100,
	}
	res := c.GetOwnerSignTaskListResponse(req, accessToken)
	fmt.Println(ModelToJsonString(res, false))
}

// 获取签署任务详情
func TestGetSignTaskDetailResponse(t *testing.T) {
	c := client.NewClient(APPID, APPSECRET, SERVERURL)
	accessTokenRes := c.GetAuthToken()
	accessToken := accessTokenRes.Data.AccessToken
	req := &signtaskRequestModel.GetSignTaskDetailReq{SignTaskId: SignTaskId}
	res := c.GetSignTaskDetailResponse(req, accessToken)
	fmt.Println(ModelToJsonString(res, false))
}

// 获取指定归属方的签署任务文档下载地址
func TestGetOwnerSignTaskUrlResponse(t *testing.T) {
	c := client.NewClient(APPID, APPSECRET, SERVERURL)
	accessTokenRes := c.GetAuthToken()
	accessToken := accessTokenRes.Data.AccessToken
	req := &signtaskRequestModel.GetOwnerSignTaskUrlReq{
		OwnerId: &commonModel.OpenId{
			IdType: "corp",
			OpenId: OpenCorpId,
		},
		SignTaskId: SignTaskId,
		FileType:   "doc",
		Id:         "",
	}
	res := c.GetOwnerSignTaskUrlResponse(req, accessToken)
	fmt.Println(ModelToJsonString(res, false))
}

// 查询企业签署任务文件夹
func TestGetSignTaskCatalogList(t *testing.T) {
	c := client.NewClient(APPID, APPSECRET, SERVERURL)
	accessTokenRes := c.GetAuthToken()
	accessToken := accessTokenRes.Data.AccessToken
	req := &signtaskRequestModel.SignTaskCatalogListReq{
		OwnerId: &commonModel.OpenId{
			IdType: "corp",
			OpenId: OpenCorpId,
		},
	}
	res := c.GetSignTaskCatalogList(req, accessToken)
	fmt.Println(ModelToJsonString(res, false))
}

// 查询签署任务控件信息
func TestGetSignTaskFieldList(t *testing.T) {
	c := client.NewClient(APPID, APPSECRET, SERVERURL)
	accessTokenRes := c.GetAuthToken()
	accessToken := accessTokenRes.Data.AccessToken
	req := &signtaskRequestModel.ListSignTaskFieldReq{
		SignTaskId: SignTaskId,
	}
	res := c.GetSignTaskFieldList(req, accessToken)
	fmt.Println(ModelToJsonString(res, false))
}

// 查询签署任务参与方信息
func TestGetSignTaskActorList(t *testing.T) {
	c := client.NewClient(APPID, APPSECRET, SERVERURL)
	accessTokenRes := c.GetAuthToken()
	accessToken := accessTokenRes.Data.AccessToken
	req := &signtaskRequestModel.ListSignTaskActorReq{
		SignTaskId: SignTaskId,
	}
	res := c.GetSignTaskActorList(req, accessToken)
	fmt.Println(ModelToJsonString(res, false))
}

// 获取刷脸地图
func TestGetSignTaskActorFacePicture(t *testing.T) {
	c := client.NewClient(APPID, APPSECRET, SERVERURL)
	accessTokenRes := c.GetAuthToken()
	accessToken := accessTokenRes.Data.AccessToken
	req := &signtaskRequestModel.GetActorFacePictureReq{
		SignTaskId: SignTaskId,
		ActorId:    "企业方",
	}
	res := c.GetSignTaskActorFacePicture(req, accessToken)
	fmt.Println(ModelToJsonString(res, false))
}

// 签署文档切图
func TestGetSignTaskSlicingTicketId(t *testing.T) {
	c := client.NewClient(APPID, APPSECRET, SERVERURL)
	accessTokenRes := c.GetAuthToken()
	accessToken := accessTokenRes.Data.AccessToken
	req := &signtaskRequestModel.GetSignTaskSlicingTicketIdReq{
		OwnerId: &commonModel.OpenId{
			IdType: "corp",
			OpenId: OpenCorpId,
		},
		SignTaskId: "",
	}
	res := c.GetSignTaskSlicingTicketId(req, accessToken)
	fmt.Println(ModelToJsonString(res, false))
}

// 获取图片版签署文档下载地址
func TestGetSignTaskPicDownloadUrl(t *testing.T) {
	c := client.NewClient(APPID, APPSECRET, SERVERURL)
	accessTokenRes := c.GetAuthToken()
	accessToken := accessTokenRes.Data.AccessToken
	req := &signtaskRequestModel.GetSignTaskPicDownloadUrlReq{SlicingTicketId: ""}
	res := c.GetSignTaskPicDownloadUrl(req, accessToken)
	fmt.Println(ModelToJsonString(res, false))
}

// 查询签署任务审批信息
func TestGetApprovalInfo(t *testing.T) {
	c := client.NewClient(APPID, APPSECRET, SERVERURL)
	accessTokenRes := c.GetAuthToken()
	accessToken := accessTokenRes.Data.AccessToken
	req := &signtaskRequestModel.GetApprovalInfoReq{
		SignTaskId: SignTaskId,
	}
	res := c.GetApprovalInfo(req, accessToken)
	fmt.Println(ModelToJsonString(res, false))
}

// 获取签署任务公证处保全报告下载地址
func TestGetSignTaskDownLoadEvidenceReport(t *testing.T) {
	c := client.NewClient(APPID, APPSECRET, SERVERURL)
	accessTokenRes := c.GetAuthToken()
	accessToken := accessTokenRes.Data.AccessToken
	req := &signtaskRequestModel.GetEvidenceReportReq{SignTaskId: SignTaskId}
	res := c.GetSignTaskDownLoadEvidenceReport(req, accessToken)
	fmt.Println(ModelToJsonString(res, false))
}

// 查询签署业务类型列表
func TestGetSignTaskBusinessTypeList(t *testing.T) {
	c := client.NewClient(APPID, APPSECRET, SERVERURL)
	accessTokenRes := c.GetAuthToken()
	accessToken := accessTokenRes.Data.AccessToken
	req := &signtaskRequestModel.GetSignTaskBusinessTypeListReq{OpenCorpId: OpenCorpId}
	res := c.GetSignTaskBusinessTypeList(req, accessToken)
	fmt.Println(ModelToJsonString(res, false))
}
