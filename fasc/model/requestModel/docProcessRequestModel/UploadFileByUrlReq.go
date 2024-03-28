package docProcessRequestModel

// UploadFileByUrlReq 通过网络文件地址上传
type UploadFileByUrlReq struct {
	FileType string `json:"fileType,omitempty"`
	FileUrl  string `json:"fileUrl,omitempty"`
}
