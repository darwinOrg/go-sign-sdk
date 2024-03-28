package dgsign

import (
	dgcoll "github.com/darwinOrg/go-common/collection"
	dgctx "github.com/darwinOrg/go-common/context"
	dgerr "github.com/darwinOrg/go-common/enums/error"
	"github.com/darwinOrg/go-common/page"
	dgsys "github.com/darwinOrg/go-common/sys"
	dghttp "github.com/darwinOrg/go-httpclient"
	dglogger "github.com/darwinOrg/go-logger"
	"github.com/darwinOrg/go-sign-sdk/fasc/client"
	"github.com/darwinOrg/go-sign-sdk/fasc/common"
	"github.com/darwinOrg/go-sign-sdk/fasc/model/commonModel"
	"github.com/darwinOrg/go-sign-sdk/fasc/model/requestModel/docProcessRequestModel"
	"github.com/darwinOrg/go-sign-sdk/fasc/model/requestModel/signtaskRequestModel"
	"github.com/darwinOrg/go-sign-sdk/fasc/model/requestModel/templateRequestModel"
	"github.com/darwinOrg/go-sign-sdk/fasc/model/responseModel/templateResponseModel"
	ve "github.com/darwinOrg/go-validator-ext"
	"github.com/go-playground/validator/v10"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

type DgSignClient struct {
	appId                string
	appSecret            string
	openCorpId           string
	platformClientUserId string
	verifyFreeBusinessId string `desc:"免验证签场景id"`
	platformOpenId       *commonModel.OpenId
	openApiClient        *client.OpenApiClient
	validator            *validator.Validate
}

func DefaultDgSignClient() *DgSignClient {
	if dgsys.IsProd() {
		return NewDgSignClient("00000325", "PIXGADRS3EZHNH3RKTA1FDWZ6GYM86KC", "e43e2a52e51e43f1be00330047f8be25", "1654428968614871040", "")
	} else {
		return NewDgSignClient("00000938", "OZMTUFSQMXYOVNK7PZEMXWPSLFW0KYWO", "90b00021c3c9439286acdf16b36b29db", "1656209484033757184", "99135b28d3b9291248dbfa5b4696e55c")
	}
}

func NewDgSignClient(appId string, appSecret string, openCorpId string, platformClientUserId string, verifyFreeBusinessId string) *DgSignClient {
	var serverUrl string
	if dgsys.IsProd() {
		serverUrl = "https://api.fadada.com/api/v5"
	} else {
		serverUrl = "https://uat-api.fadada.com/api/v5"
	}

	return &DgSignClient{
		appId:                appId,
		appSecret:            appSecret,
		openCorpId:           openCorpId,
		platformClientUserId: platformClientUserId,
		verifyFreeBusinessId: verifyFreeBusinessId,
		platformOpenId: &commonModel.OpenId{
			IdType: ACTOR_TYPE_CORP.String(),
			OpenId: openCorpId,
		},
		openApiClient: client.NewClient(appId, appSecret, serverUrl),
		validator:     ve.NewCustomValidator(),
	}
}

// GetSignTemplateList 查询签署任务模板列表
func (c *DgSignClient) GetSignTemplateList(ctx *dgctx.DgContext, request *GetSignTemplateListRequest) (*page.PageList[SignTemplate], error) {
	err := c.validate(ctx, request)
	if err != nil {
		return nil, err
	}

	accessToken, err := c.getAccessToken(ctx)
	if err != nil {
		return nil, err
	}

	req := &templateRequestModel.GetSignTemplateListReq{
		ListFilter: &templateRequestModel.SignTemplateListFilterInfo{
			SignTemplateName: request.TemplateName,
		},
	}
	dglogger.Infof(ctx, "get sign template list request: %s", mustJsonMarshal(req))
	resp := c.openApiClient.GetSignTemplateList(req, accessToken)
	if resp.Code != successCode {
		dglogger.Errorf(ctx, "get sign template list fail, templateName: %s, code: %s, msg: %s", request.TemplateName, resp.Code, resp.Msg)
		return nil, dgerr.SimpleDgError(resp.Msg)
	}
	dglogger.Infof(ctx, "get sign template list response: %s", mustJsonMarshal(resp))

	signTemplates := dgcoll.MapToList(resp.Data.SignTemplates, func(t templateResponseModel.SignTemplateListInfo) *SignTemplate {
		return &SignTemplate{
			TemplateId:     t.SignTemplateId,
			TemplateName:   t.SignTemplateName,
			TemplateStatus: t.SignTemplateStatus,
		}
	})

	return &page.PageList[SignTemplate]{
		PageNo:     request.PageNo,
		PageSize:   request.PageSize,
		TotalCount: resp.Data.TotalCount,
		TotalPages: resp.Data.ListPageCount,
		List:       signTemplates,
	}, nil
}

// CreateSignTaskWithTemplate 创建签署任务基于签署模板
func (c *DgSignClient) CreateSignTaskWithTemplate(ctx *dgctx.DgContext, request *CreateSignTaskWithTemplateRequest) (*CreateSignTaskResponse, error) {
	err := c.validate(ctx, request)
	if err != nil {
		return nil, err
	}

	accessToken, err := c.getAccessToken(ctx)
	if err != nil {
		return nil, err
	}

	tdResp := c.openApiClient.GetSignTemplateDetail(&templateRequestModel.GetSignTemplateDetailReq{
		OwnerId:        c.platformOpenId,
		SignTemplateId: request.TemplateId,
	}, accessToken)
	if tdResp.Code != successCode {
		dglogger.Errorf(ctx, "get template detail fail, templateId: %s, code: %s, msg: %s", request.TemplateId, tdResp.Code, tdResp.Msg)
		return nil, dgerr.SimpleDgError(tdResp.Msg)
	}

	req := &signtaskRequestModel.CreateSignTaskWithTemplateReq{
		Initiator:        c.platformOpenId,
		SignTaskSubject:  request.Subject,
		SignTemplateId:   request.TemplateId,
		ExpiresTime:      strconv.FormatInt(request.ExpiresTime.UnixMilli(), 10),
		AutoStart:        request.AutoStart,
		AutoFillFinalize: true,
	}
	if request.PlatformVerifyFree {
		req.BusinessScene = &signtaskRequestModel.BusinessSceneInfo{
			BusinessId:       c.verifyFreeBusinessId,
			TransReferenceId: request.TransReferenceId,
		}
	}

	for i, actor := range request.Actors {
		ar := &commonModel.Actor{
			ActorType:      actor.Actor.ActorType.String(),
			ActorId:        actor.Actor.ActorId,
			ActorName:      actor.Actor.ActorName,
			Permissions:    dgcoll.MapToList(actor.Actor.Permissions, func(p ActorPermission) string { return p.String() }),
			CertNoForMatch: actor.CertNoForMatch,
		}

		if ACTOR_TYPE_CORP == actor.Actor.ActorType {
			ar.ActorOpenId = c.openCorpId
		}

		if actor.Notification != nil {
			ar.Notification = &commonModel.Notification{
				SendNotification: true,
				NotifyWay:        actor.Notification.NotifyWay.String(),
				NotifyAddress:    actor.Notification.NotifyAddress,
			}
		}

		var fillFields []commonModel.FillField
		if len(actor.FillParams) > 0 {
			filteredActors := dgcoll.FilterList(tdResp.Data.Actors, func(ar templateResponseModel.SignTaskTemActor) bool {
				return ar.ActorInfo.ActorId == actor.Actor.ActorId
			})
			if len(filteredActors) == 0 {
				dglogger.Errorf(ctx, "not found actorId: %s, templateId: %s", actor.Actor.ActorId, request.TemplateId)
				return nil, dgerr.SimpleDgError("not found actorId")
			}
			ffs := filteredActors[0].FillFields
			fieldId2FieldDocIdMap := dgcoll.Extract2Map(
				ffs,
				func(ff templateResponseModel.FillField) string { return ff.FieldId },
				func(ff templateResponseModel.FillField) string { return ff.FieldDocId },
			)

			for fieldId, fieldValue := range actor.FillParams {
				if fieldId2FieldDocIdMap[fieldId] == "" {
					dglogger.Errorf(ctx, "not found fieldId: %s, templateId: %s", fieldId, request.TemplateId)
					return nil, dgerr.SimpleDgError("not found fieldId")
				}

				fillFields = append(fillFields, commonModel.FillField{
					FieldId:    fieldId,
					FieldValue: fieldValue,
					FieldDocId: fieldId2FieldDocIdMap[fieldId],
				})
			}
		}

		req.Actors = append(req.Actors, commonModel.SignTaskActor{
			Actor:      ar,
			FillFields: fillFields,
			SignConfigInfo: &commonModel.SignConfigInfo{
				OrderNo:           i,
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
		})
	}
	dglogger.Infof(ctx, "create sign task with template request: %s", mustJsonMarshal(req))
	resp := c.openApiClient.CreateSignTaskWithTemplate(req, accessToken)
	if resp.Code != successCode {
		dglogger.Errorf(ctx, "create sign task fail, requestId: %s, code: %s, msg: %s", resp.RequestId, resp.Code, resp.Msg)
		return nil, dgerr.SimpleDgError(resp.Msg)
	}

	if resp.Data.SignTaskId == "" {
		dglogger.Errorf(ctx, "create sign task fail, taskId is blank, requestId: %s", resp.RequestId)
		return nil, dgerr.SimpleDgError("taskId is blank")
	}

	taskResp := &CreateSignTaskResponse{
		TaskId: resp.Data.SignTaskId,
	}
	dglogger.Infof(ctx, "create sign task with template response: %s", mustJsonMarshal(taskResp))

	return taskResp, nil
}

// GetSignTaskEditUrl 获取签署任务编辑链接
func (c *DgSignClient) GetSignTaskEditUrl(ctx *dgctx.DgContext, request *GetSignTaskEditUrlRequest) (*GetSignTaskEditUrlResponse, error) {
	err := c.validate(ctx, request)
	if err != nil {
		return nil, err
	}

	accessToken, err := c.getAccessToken(ctx)
	if err != nil {
		return nil, err
	}

	req := &signtaskRequestModel.GetSignTaskUrlReq{
		SignTaskId:  request.TaskId,
		RedirectUrl: url.PathEscape(request.RedirectUrl),
	}
	dglogger.Infof(ctx, "get sign task edit url request: %s", mustJsonMarshal(req))
	resp := c.openApiClient.GetSignTaskEditUrl(req, accessToken)
	if resp.Code != successCode {
		dglogger.Errorf(ctx, "get sign task edit url fail, taskId: %s, redirectUrl: %s, code: %s, msg: %s",
			request.TaskId, request.RedirectUrl, resp.Code, resp.Msg)
		return nil, dgerr.SimpleDgError(resp.Msg)
	}
	dglogger.Infof(ctx, "get sign task edit url response: %s", mustJsonMarshal(resp))

	return &GetSignTaskEditUrlResponse{
		SignTaskEditUrl: resp.Data.SignTaskEditUrl,
	}, nil
}

// StartSignTask 提交签署任务
func (c *DgSignClient) StartSignTask(ctx *dgctx.DgContext, request *StartSignTaskRequest) error {
	err := c.validate(ctx, request)
	if err != nil {
		return err
	}

	accessToken, err := c.getAccessToken(ctx)
	if err != nil {
		return err
	}

	req := &signtaskRequestModel.StartSignTaskReq{
		SignTaskId: request.TaskId,
	}
	dglogger.Infof(ctx, "start sign task request: %s", mustJsonMarshal(req))
	resp := c.openApiClient.StartSignTask(req, accessToken)
	if resp.Code != successCode {
		dglogger.Errorf(ctx, "start sign task fail, taskId: %s, code: %s, msg: %s", request.TaskId, resp.Code, resp.Msg)
		return dgerr.SimpleDgError(resp.Msg)
	}
	dglogger.Infof(ctx, "start sign task response: %s", mustJsonMarshal(resp))

	return nil
}

// GetSignTaskPreviewUrl 获取签署任务预览链接
func (c *DgSignClient) GetSignTaskPreviewUrl(ctx *dgctx.DgContext, request *GetSignTaskPreviewUrlRequest) (*GetSignTaskPreviewUrlResponse, error) {
	err := c.validate(ctx, request)
	if err != nil {
		return nil, err
	}

	accessToken, err := c.getAccessToken(ctx)
	if err != nil {
		return nil, err
	}

	req := &signtaskRequestModel.GetSignTaskUrlReq{
		SignTaskId:  request.TaskId,
		RedirectUrl: url.PathEscape(request.RedirectUrl),
	}
	dglogger.Infof(ctx, "get sign task preview url request: %s", mustJsonMarshal(req))
	resp := c.openApiClient.GetSignTaskPreviewUrl(req, accessToken)
	if resp.Code != successCode {
		dglogger.Errorf(ctx, "get sign task preview url fail, taskId: %s, redirectUrl: %s, code: %s, msg: %s",
			request.TaskId, request.RedirectUrl, resp.Code, resp.Msg)
		return nil, dgerr.SimpleDgError(resp.Msg)
	}
	dglogger.Infof(ctx, "get sign task preview url response: %s", mustJsonMarshal(resp))

	return &GetSignTaskPreviewUrlResponse{
		SignTaskPreviewUrl: resp.Data.SignTaskPreviewUrl,
	}, nil
}

// GetActorSignTaskUrl 获取参与方签署链接
func (c *DgSignClient) GetActorSignTaskUrl(ctx *dgctx.DgContext, request *GetActorSignTaskUrlRequest) (*GetActorSignTaskUrlResponse, error) {
	err := c.validate(ctx, request)
	if err != nil {
		return nil, err
	}

	accessToken, err := c.getAccessToken(ctx)
	if err != nil {
		return nil, err
	}

	req := &signtaskRequestModel.SignTaskActorGetUrlReq{
		SignTaskId:   request.TaskId,
		RedirectUrl:  url.PathEscape(request.RedirectUrl),
		ActorId:      request.Actor.ActorId,
		ClientUserId: request.ClientUserId,
	}
	if request.Actor == ACTOR_PLATFORM {
		req.ClientUserId = c.platformClientUserId
	}
	dglogger.Infof(ctx, "get actor sign task url request: %s", mustJsonMarshal(req))
	resp := c.openApiClient.GetSignTaskActorUrl(req, accessToken)
	if resp.Code != successCode {
		dglogger.Errorf(ctx, "get actor sign task url fail, taskId: %s, code: %s, msg: %s", request.TaskId, resp.Code, resp.Msg)
		return nil, dgerr.SimpleDgError(resp.Msg)
	}

	if resp.Data.ActorSignTaskEmbedUrl == "" {
		dglogger.Errorf(ctx, "get actor sign task url fail, embed url is blank, taskId: %s", request.TaskId)
		return nil, dgerr.SimpleDgError("embed url is blank")
	}

	taskResp := &GetActorSignTaskUrlResponse{
		EmbedUrl: resp.Data.ActorSignTaskEmbedUrl,
	}
	dglogger.Infof(ctx, "get actor sign task url response: %s", mustJsonMarshal(taskResp))

	return taskResp, nil
}

// UploadLocalFile 上传本地文件
func (c *DgSignClient) UploadLocalFile(ctx *dgctx.DgContext, request *UploadLocalFileRequest) (*UploadResponse, error) {
	err := c.validate(ctx, request)
	if err != nil {
		return nil, err
	}
	fh, err := os.Open(request.FilePath)
	if err != nil {
		dglogger.Errorf(ctx, "error opening file: %s", request.FilePath)
		return nil, dgerr.SimpleDgError("error opening file")
	}
	defer fh.Close()

	return c.UploadBody(ctx, &UploadBodyRequest{
		FileType: request.FileType,
		FileName: request.FileName,
		Body:     fh,
	})
}

// UploadBody 上传内容体
func (c *DgSignClient) UploadBody(ctx *dgctx.DgContext, request *UploadBodyRequest) (*UploadResponse, error) {
	err := c.validate(ctx, request)
	if err != nil {
		return nil, err
	}

	accessToken, err := c.getAccessToken(ctx)
	if err != nil {
		return nil, err
	}

	uploadReq := &docProcessRequestModel.GetLocalUploadFileUrlReq{
		FileType: request.FileType.String(),
	}
	dglogger.Infof(ctx, "get local upload file url request: %s", mustJsonMarshal(uploadReq))
	uploadResp := c.openApiClient.GetLocalFileUploadUrl(uploadReq, accessToken)
	if uploadResp.Code != successCode {
		dglogger.Errorf(ctx, "get local upload file url fail, fileType: %s, code: %s, msg: %s", request.FileType, uploadResp.Code, uploadResp.Msg)
		return nil, dgerr.SimpleDgError(uploadResp.Msg)
	}
	if uploadResp.Data.UploadUrl == "" {
		dglogger.Errorf(ctx, "get local upload file url fail, upload url is blank, fileType: %s", request.FileType)
		return nil, dgerr.SimpleDgError("upload url is blank")
	}
	dglogger.Infof(ctx, "get local upload file url response: %s", mustJsonMarshal(uploadResp))

	_, err = dghttp.GlobalHttpClient.DoUploadBody(ctx, http.MethodPut, uploadResp.Data.UploadUrl, request.Body)
	if err != nil {
		dglogger.Errorf(ctx, "upload local file fail, upload url: %s", uploadResp.Data.UploadUrl)
		return nil, err
	}

	processReq := &docProcessRequestModel.FileProcessReq{
		FddFileUrlList: []docProcessRequestModel.FddUploadFile{
			{
				FileType:   request.FileType.String(),
				FileName:   request.FileName,
				FddFileUrl: uploadResp.Data.FddFileUrl,
			},
		},
	}
	dglogger.Infof(ctx, "process file request: %s", mustJsonMarshal(processReq))
	processResp := c.openApiClient.FileProcess(processReq, accessToken)
	if processResp.Code != successCode {
		dglogger.Errorf(ctx, "process file fail, fileType: %s, fileName: %s, fddFileUrl: %s, code: %s, msg: %s",
			request.FileType, request.FileName, uploadResp.Data.FddFileUrl, processResp.Code, processResp.Msg)
		return nil, dgerr.SimpleDgError(processResp.Msg)
	}
	if len(processResp.Data.FileIdList) == 0 || processResp.Data.FileIdList[0].FileId == "" {
		dglogger.Errorf(ctx, "process file fail, fileId is blank, fileType: %s, fileName: %s, fddFileUrl: %s",
			request.FileType, request.FileName, uploadResp.Data.FddFileUrl)
		return nil, dgerr.SimpleDgError("fileId is blank")
	}
	dglogger.Infof(ctx, "process file response: %s", mustJsonMarshal(processResp))

	resp := &UploadResponse{
		FileId: processResp.Data.FileIdList[0].FileId,
	}
	dglogger.Infof(ctx, "upload local file response: %s", mustJsonMarshal(resp))

	return resp, nil
}

// AddAttachments 添加附件
func (c *DgSignClient) AddAttachments(ctx *dgctx.DgContext, request *AddAttachmentsRequest) error {
	err := c.validate(ctx, request)
	if err != nil {
		return err
	}

	accessToken, err := c.getAccessToken(ctx)
	if err != nil {
		return err
	}

	attachReq := &signtaskRequestModel.AddSignTaskAttachReq{
		SignTaskId: request.TaskId,
		Attachs: dgcoll.MapToList(request.Attachments, func(attach *Attachment) commonModel.Attach {
			return commonModel.Attach{
				AttachId:     attach.AttachId,
				AttachName:   attach.AttachName,
				AttachFileId: attach.AttachFileId,
			}
		}),
	}
	dglogger.Infof(ctx, "add attachments request: %s", mustJsonMarshal(attachReq))
	uploadResp := c.openApiClient.AddSignTaskAttach(attachReq, accessToken)
	if uploadResp.Code != successCode {
		dglogger.Errorf(ctx, "add attachments fail, taskId: %s, code: %s, msg: %s", request.TaskId, uploadResp.Code, uploadResp.Msg)
		return dgerr.SimpleDgError(uploadResp.Msg)
	}

	return nil
}

// DeleteAttachments 删除附件
func (c *DgSignClient) DeleteAttachments(ctx *dgctx.DgContext, request *DeleteAttachmentsRequest) error {
	err := c.validate(ctx, request)
	if err != nil {
		return err
	}

	accessToken, err := c.getAccessToken(ctx)
	if err != nil {
		return err
	}

	attachReq := &signtaskRequestModel.DeleteSignTaskAttachReq{
		SignTaskId: request.TaskId,
		AttachIds:  request.AttachIds,
	}
	dglogger.Infof(ctx, "delete attachments request: %s", mustJsonMarshal(attachReq))
	uploadResp := c.openApiClient.DeleteSignTaskAttach(attachReq, accessToken)
	if uploadResp.Code != successCode {
		dglogger.Errorf(ctx, "delete attachments fail, taskId: %s, code: %s, msg: %s", request.TaskId, uploadResp.Code, uploadResp.Msg)
		return dgerr.SimpleDgError(uploadResp.Msg)
	}

	return nil
}

// GetSignTaskPackageFileContent 获取签署任务打包文件内容
func (c *DgSignClient) GetSignTaskPackageFileContent(ctx *dgctx.DgContext, request *GetSignTaskPackageFileContentRequest) (*GetSignTaskPackageFileContentResponse, error) {
	err := c.validate(ctx, request)
	if err != nil {
		return nil, err
	}

	accessToken, err := c.getAccessToken(ctx)
	if err != nil {
		return nil, err
	}

	req := &signtaskRequestModel.GetOwnerSignTaskUrlReq{
		SignTaskId: request.TaskId,
		OwnerId:    c.platformOpenId,
	}
	dglogger.Infof(ctx, "get owner sign task url request: %s", mustJsonMarshal(req))
	resp := c.openApiClient.GetOwnerSignTaskUrlResponse(req, accessToken)
	if resp.Code != successCode {
		dglogger.Errorf(ctx, "get owner sign task download url fail, taskId: %s, code: %s, msg: %s", request.TaskId, resp.Code, resp.Msg)
		return nil, dgerr.SimpleDgError(resp.Msg)
	}
	if resp.Data.DownloadUrl == "" {
		dglogger.Errorf(ctx, "get owner sign task download url fail, download url is blank, taskId: %s", request.TaskId)
		return nil, dgerr.SimpleDgError("download url is blank")
	}
	dglogger.Infof(ctx, "get owner sign task url response: %s", mustJsonMarshal(resp))

	bodyBytes, err := readUrlBody(resp.Data.DownloadUrl)
	if err != nil {
		dglogger.Errorf(ctx, "get owner sign task download content fail, taskId: %s, error: %v", request.TaskId, err)
		return nil, err
	}

	if len(bodyBytes) == 0 {
		dglogger.Errorf(ctx, "get owner sign task download content empty, taskId: %s", request.TaskId)
		return nil, dgerr.SimpleDgError("download content is empty")
	}

	taskResp := &GetSignTaskPackageFileContentResponse{
		Content: bodyBytes,
	}

	return taskResp, nil
}

func (c *DgSignClient) validate(ctx *dgctx.DgContext, req any) error {
	err := c.validator.Struct(req)
	if err != nil {
		errMsg := ve.TranslateValidateError(err, ctx.Lang)
		return dgerr.SimpleDgError(errMsg)
	}
	return nil
}

func (c *DgSignClient) GetSign(timestamp string, nonce string, event string, bizContent string) string {
	headMap := map[string]string{
		"X-FASC-App-Id":    c.appId,
		"X-FASC-Sign-Type": client.SignType,
		"X-FASC-Timestamp": timestamp,
		"X-FASC-Nonce":     nonce,
		"X-FASC-Event":     event,
		"bizContent":       bizContent,
	}
	return common.GetSignByMap(headMap, timestamp, c.appSecret)
}

// getAccessToken 获取服务凭证
func (c *DgSignClient) getAccessToken(ctx *dgctx.DgContext) (string, error) {
	if accessToken := ctx.GetExtraValue(accessTokenKey); accessToken != nil {
		return accessToken.(string), nil
	}

	tokenResp := c.openApiClient.GetAuthToken()
	if tokenResp.Code != successCode {
		dglogger.Errorf(ctx, "get auth token fail, code: %s, msg: %s", tokenResp.Code, tokenResp.Msg)
		return "", dgerr.SimpleDgError(tokenResp.Msg)
	}

	if tokenResp.Data.AccessToken == "" {
		dglogger.Errorf(ctx, "auth token")
		return "", dgerr.SimpleDgError("auth token is blank")
	}

	dglogger.Infof(ctx, "get access token: %s", tokenResp.Data.AccessToken)
	ctx.SetExtraKeyValue(accessTokenKey, tokenResp.Data.AccessToken)

	return tokenResp.Data.AccessToken, nil
}
