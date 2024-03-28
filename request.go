package dgsign

import (
	"io"
	"time"
)

type CreateSignTaskWithTemplateRequest struct {
	Subject            string           `desc:"主题" json:"subject" binding:"required"`
	TemplateId         string           `desc:"模板id" json:"templateId" binding:"required"`
	ExpiresTime        time.Time        `desc:"过期时间" json:"expiresTime" binding:"required"`
	AutoStart          bool             `desc:"是否自动提交签署任务" json:"autoStart"`
	PlatformVerifyFree bool             `desc:"是否免验签" json:"platformVerifyFree"`
	TransReferenceId   string           `desc:"业务参考id" json:"transReferenceId" binding:"maxLength=128"`
	Actors             []*SignTaskActor `desc:"参与方" json:"actors" binding:"required,minLength=1"`
}

type SignTaskActor struct {
	Actor          *Actor            `desc:"参与方" json:"actors" binding:"required"`
	FillParams     map[string]string `desc:"填充参数" json:"fillParams"`
	CertNoForMatch string            `desc:"参与方证件号码（个人身份证号或企业统信码）" json:"certNoForMatch" binding:"maxLength=32"`
	Notification   *Notification     `desc:"通知信息" json:"notification"`
}

type Notification struct {
	NotifyWay     NotifyWay `desc:"通知方式" json:"notifyWay" binding:"required,mustIn=mobile#email"`
	NotifyAddress string    `desc:"通知地址" json:"notifyAddress" binding:"required"`
}

type StartSignTaskRequest struct {
	TaskId string `desc:"任务id" json:"taskId'" binding:"required"`
}

type GetSignTaskEditUrlRequest struct {
	TaskId      string `desc:"任务id" json:"taskId'" binding:"required"`
	RedirectUrl string `desc:"重定向url" json:"redirectUrl"`
}

type GetSignTaskPreviewUrlRequest struct {
	TaskId      string `desc:"任务id" json:"taskId'" binding:"required"`
	RedirectUrl string `desc:"重定向url" json:"redirectUrl"`
}

type GetActorSignTaskUrlRequest struct {
	TaskId       string `desc:"任务id" json:"taskId'" binding:"required"`
	RedirectUrl  string `desc:"重定向url" json:"redirectUrl"`
	Actor        *Actor `desc:"参与方" json:"actors" binding:"required"`
	ClientUserId string `desc:"应用系统用户标志" json:"platformClientUserId'"`
}

type GetSignTaskPackageFileContentRequest struct {
	TaskId string `desc:"任务id" json:"taskId'" binding:"required"`
}

type UploadLocalFileRequest struct {
	FileType FileType `desc:"文件类型" json:"fileType'" binding:"required,mustIn=doc#attach"`
	FileName string   `desc:"文件名称" json:"fileName'" binding:"required"`
	FilePath string   `desc:"文件路径" json:"filePath'" binding:"required"`
}

type UploadBodyRequest struct {
	FileType FileType  `desc:"文件类型" json:"fileType'" binding:"required,mustIn=doc#attach"`
	FileName string    `desc:"文件名称" json:"fileName'" binding:"required"`
	Body     io.Reader `desc:"内容体" json:"body'" binding:"required,minLength=1"`
}

type AddAttachmentsRequest struct {
	TaskId      string        `desc:"任务id" json:"taskId'" binding:"required"`
	Attachments []*Attachment `desc:"附件" json:"attachments'" binding:"required,minLength=1"`
}

type Attachment struct {
	AttachId     string `desc:"附件id" json:"AttachId'" binding:"required"`
	AttachName   string `desc:"附件名称" json:"attachName'" binding:"required"`
	AttachFileId string `desc:"附件文件id" json:"attachFileId'" binding:"required"`
}

type DeleteAttachmentsRequest struct {
	TaskId    string   `desc:"任务id" json:"taskId'" binding:"required"`
	AttachIds []string `desc:"附件id" json:"attachIds'" binding:"required,minLength=1"`
}

type GetSignTemplateListRequest struct {
	TemplateName string `desc:"模板名称，模糊查询" json:"templateName" binding:"maxLength=100"`
	PageNo       int    `desc:"页码" json:"pageNo" binding:"required,min=1"`
	PageSize     int    `desc:"每页记录数" json:"pageSize" binding:"required,min=1"`
}
