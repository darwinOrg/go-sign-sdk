package dgsign

import (
	dgctx "github.com/darwinOrg/go-common/context"
	dglogger "github.com/darwinOrg/go-logger"
	"os"
	"testing"
	"time"
)

const (
	templateId = "1683707663168191410"
	taskId     = "1683774564368192515"
)

func TestGetSignTemplateList(t *testing.T) {
	ctx := &dgctx.DgContext{TraceId: "123"}
	c := DefaultDgSignClient()
	resp, err := c.GetSignTemplateList(ctx, &GetSignTemplateListRequest{
		TemplateName: "测试",
		PageNo:       1,
		PageSize:     10,
	})
	if err != nil {
		dglogger.Error(ctx, err)
	} else {
		dglogger.Info(ctx, resp)
	}
}

func TestCreateSignTaskWithTemplate(t *testing.T) {
	ctx := &dgctx.DgContext{TraceId: "123"}
	c := DefaultDgSignClient()
	resp, err := c.CreateSignTaskWithTemplate(ctx, &CreateSignTaskWithTemplateRequest{
		Subject:    "测试签约",
		TemplateId: templateId,
		//TransReferenceId: "456",
		ExpiresTime: time.Now().Add(time.Hour * 72),
		AutoStart:   false,
		Actors: []*SignTaskActor{
			{
				Actor: ACTOR_PLATFORM,
				FillParams: map[string]string{
					"input1":  "text1",
					"input2":  "text2",
					"input3":  "123",
					"radio1":  "[true,false]",
					"check1":  "[true]",
					"select1": "选项1",
					"date1":   "2023-05-20",
				},
				Notification: &Notification{
					NotifyWay:     NOTIFY_WAY_MOBILE,
					NotifyAddress: "1590143xxxx",
				},
			},
			{
				Actor: ACTOR_EMPLOYEE,
				Notification: &Notification{
					NotifyWay:     NOTIFY_WAY_EMAIL,
					NotifyAddress: "1590143xxxx@163.com",
				},
			},
		},
	})
	if err != nil {
		dglogger.Error(ctx, err)
	} else {
		dglogger.Info(ctx, resp)
	}
}

func TestGetSignTaskEditUrl(t *testing.T) {
	ctx := &dgctx.DgContext{TraceId: "123"}
	c := DefaultDgSignClient()
	resp, err := c.GetSignTaskEditUrl(ctx, &GetSignTaskEditUrlRequest{
		TaskId:      taskId,
		RedirectUrl: "https://www.baidu.com",
	})
	if err != nil {
		dglogger.Error(ctx, err)
	} else {
		dglogger.Info(ctx, resp)
	}
}

func TestGetSignTaskPreviewUrl(t *testing.T) {
	ctx := &dgctx.DgContext{TraceId: "123"}
	c := DefaultDgSignClient()
	resp, err := c.GetSignTaskPreviewUrl(ctx, &GetSignTaskPreviewUrlRequest{
		TaskId:      taskId,
		RedirectUrl: "https://www.baidu.com",
	})
	if err != nil {
		dglogger.Error(ctx, err)
	} else {
		dglogger.Info(ctx, resp)
	}
}

func TestStartSignTask(t *testing.T) {
	ctx := &dgctx.DgContext{TraceId: "123"}
	c := DefaultDgSignClient()
	err := c.StartSignTask(ctx, &StartSignTaskRequest{
		TaskId: taskId,
	})
	if err != nil {
		dglogger.Error(ctx, err)
	}
}

func TestGetActorSignTaskUrl4Platform(t *testing.T) {
	ctx := &dgctx.DgContext{TraceId: "123"}
	c := DefaultDgSignClient()
	resp, err := c.GetActorSignTaskUrl(ctx, &GetActorSignTaskUrlRequest{
		TaskId:      taskId,
		RedirectUrl: "https://www.baidu.com",
		Actor:       ACTOR_PLATFORM,
	})
	if err != nil {
		dglogger.Error(ctx, err)
	} else {
		dglogger.Info(ctx, resp)
	}
}

func TestGetActorSignTaskUrl4Employee(t *testing.T) {
	ctx := &dgctx.DgContext{TraceId: "123"}
	c := DefaultDgSignClient()
	resp, err := c.GetActorSignTaskUrl(ctx, &GetActorSignTaskUrlRequest{
		TaskId:       taskId,
		RedirectUrl:  "https://www.baidu.com",
		Actor:        ACTOR_EMPLOYEE,
		ClientUserId: "789",
	})
	if err != nil {
		dglogger.Error(ctx, err)
	} else {
		dglogger.Info(ctx, resp)
	}
}

func TestUploadLocalFile(t *testing.T) {
	ctx := &dgctx.DgContext{TraceId: "123"}
	c := DefaultDgSignClient()
	resp, err := c.UploadLocalFile(ctx, &UploadLocalFileRequest{
		FileType: FILE_TYPE_ATTACH,
		FilePath: "/Users/mac/Downloads/4.jpeg",
		FileName: "4.jpeg",
	})
	if err != nil {
		dglogger.Error(ctx, err)
	} else {
		dglogger.Info(ctx, resp)
	}
}

func TestAddAttachments(t *testing.T) {
	ctx := &dgctx.DgContext{TraceId: "123"}
	c := DefaultDgSignClient()
	err := c.AddAttachments(ctx, &AddAttachmentsRequest{
		TaskId: taskId,
		Attachments: []*Attachment{
			{
				AttachId:     "",
				AttachName:   "",
				AttachFileId: "",
			},
		},
	})
	if err != nil {
		dglogger.Error(ctx, err)
	}
}

func TestGetSignTaskPackageFileContent(t *testing.T) {
	ctx := &dgctx.DgContext{TraceId: "123"}
	c := DefaultDgSignClient()
	resp, err := c.GetSignTaskPackageFileContent(ctx, &GetSignTaskPackageFileContentRequest{
		TaskId: taskId,
	})
	if err != nil {
		dglogger.Error(ctx, err)
	} else {
		dglogger.Info(ctx, len(resp.Content))
		os.WriteFile("/Users/mac/Downloads/"+taskId+".zip", resp.Content, os.ModePerm)
	}
}
