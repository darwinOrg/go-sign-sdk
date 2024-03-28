package signtaskRequestModel

import "github.com/darwinOrg/go-sign-sdk/fasc/model/commonModel"

type SignTaskCatalogListReq struct {
	OwnerId *commonModel.OpenId `json:"ownerId,omitempty"`
}
