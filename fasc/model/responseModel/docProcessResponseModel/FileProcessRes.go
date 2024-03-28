package docProcessResponseModel

// FileProcessRes 文件处理

type FileProcessRes struct {
	RequestId string             `json:"requestId"`
	Code      string             `json:"code"`
	Msg       string             `json:"msg"`
	Data      FileProcessResData `json:"data"`
}

type FileProcessResData struct {
	FileIdList []FileProcessArr `json:"fileIdList"`
}

type FileProcessArr struct {
	FileId string `json:"fileId"`
}
