package docProcessResponseModel

type VerifyFileSignRes struct {
	RequestId string             `json:"requestId"`
	Code      string             `json:"code"`
	Msg       string             `json:"msg"`
	Data      VerifyFileSignData `json:"data"`
}

type VerifyFileSignData struct {
	VerifyResult  bool            `json:"verifyResult,omitempty"`
	Reason        string          `json:"reason,omitempty"`
	SignatureInfo []SignatureInfo `json:"signatureInfo,omitempty"`
}

type SignatureInfo struct {
	Signer              string `json:"signer,omitempty"`
	SignedonTime        string `json:"signedonTime,omitempty"`
	Authority           string `json:"authority,omitempty"`
	TimestampFlag       bool   `json:"timestampFlag,omitempty"`
	TimestampTime       string `json:"timestampTime,omitempty"`
	TimestampVerifyFlag bool   `json:"timestampVerifyFlag,omitempty"`
	IntegrityFlag       bool   `json:"integrityFlag,omitempty"`
}
