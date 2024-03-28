package orgRequestModel

type GetCorpMemberDetailReq struct {
	OpenCorpId string `json:"openCorpId,omitempty"`
	MemberId   int64  `json:"memberId"`
}
