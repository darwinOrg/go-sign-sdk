package dgsign

type CreateSignTaskResponse struct {
	TaskId string `desc:"任务id" json:"taskId"`
}

type GetActorSignTaskUrlResponse struct {
	EmbedUrl string `desc:"嵌入url" json:"embedUrl"`
}

type GetSignTaskEditUrlResponse struct {
	SignTaskEditUrl string `desc:"签署任务编辑链接" json:"signTaskEditUrl"`
}

type GetSignTaskPreviewUrlResponse struct {
	SignTaskPreviewUrl string `desc:"签署任务预览链接" json:"signTaskPreviewUrl"`
}

type GetSignTaskPackageFileContentResponse struct {
	Content []byte `desc:"内容" json:"content"`
}

type UploadResponse struct {
	FileId string `desc:"文件id" json:"fileId"`
}

type SignTemplate struct {
	TemplateId     string `desc:"模板id" json:"templateId"`
	TemplateName   string `desc:"模板名称" json:"templateName"`
	TemplateStatus string `desc:"模板状态" json:"templateStatus"`
}
