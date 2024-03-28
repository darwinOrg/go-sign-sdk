package signtaskRequestModel

import "github.com/darwinOrg/go-sign-sdk/fasc/model/commonModel"

type GetSignTaskSlicingTicketIdReq struct {
	OwnerId    *commonModel.OpenId `json:"ownerId,omitempty"`
	SignTaskId string              `json:"SignTaskId,omitempty"`
}
