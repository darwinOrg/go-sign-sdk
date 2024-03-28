package signtaskRequestModel

import "github.com/darwinOrg/go-sign-sdk/fasc/model/commonModel"

type GetBatchSignUrlReq struct {
	OwnerId      *commonModel.OpenId `json:"ownerId,omitempty"`
	MemberId     *int64              `json:"memberId,omitempty"`
	ClientUserId string              `json:"clientUserId,omitempty"`
	SignTaskIds  []string            `json:"signTaskIds,omitempty"`
	RedirectUrl  string              `json:"redirectUrl,omitempty"`
}
