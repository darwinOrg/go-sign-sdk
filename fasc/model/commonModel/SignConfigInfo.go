package commonModel

type SignConfigInfo struct {
	OrderNo           int      `json:"orderNo,omitempty"`
	RequestMemberSign bool     `json:"requestMemberSign"`
	VerifyMethods     []string `json:"verifyMethods,omitempty"`
	ReadingToEnd      bool     `json:"readingToEnd"`
	ReadingTime       string   `json:"readingTime,omitempty"`
	BlockHere         bool     `json:"blockHere"`
	RequestVerifyFree bool     `json:"requestVerifyFree"`
	SignerSignMethod  string   `json:"signerSignMethod,omitempty"`
	JoinByLink        bool     `json:"joinByLink"`
	IdentifiedView    bool     `json:"identifiedView"`
	ResizeSeal        bool     `json:"resizeSeal"`
}
