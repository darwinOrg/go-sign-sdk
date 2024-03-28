package docProcessRequestModel

// FileProcessReq 文件处理
type FileProcessReq struct {
	FddFileUrlList []FddUploadFile `json:"fddFileUrlList,omitempty"`
}

type FddUploadFile struct {
	FileType   string `json:"fileType,omitempty"`
	FddFileUrl string `json:"fddFileUrl,omitempty"`
	FileName   string `json:"fileName,omitempty"`
}
