package orgResponseModel

type CreateCorpMemberRes struct {
	RequestId string               `json:"requestId"`
	Code      string               `json:"code"`
	Msg       string               `json:"msg"`
	Data      []SimpleEmployeeInfo `json:"data"`
}
type SimpleEmployeeInfo struct {
	MemberId             int64  `json:"memberId,string"`
	MemberActiveUrl      string `json:"memberActiveUrl,omitempty"`
	MemberActiveEmbedUrl string `json:"memberActiveEmbedUrl,omitempty"`
}
