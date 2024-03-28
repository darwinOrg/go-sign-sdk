package docProcessRequestModel

type VerifyFileSignReq struct {
	FileId   string `json:"fileId,omitempty"`
	FileHash string `json:"fileHash,omitempty"`
}
