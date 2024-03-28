package signtaskRequestModel

type GetSignTaskUrlReq struct {
	SignTaskId  string `json:"signTaskId,omitempty"`
	RedirectUrl string `json:"redirectUrl,omitempty"`
}
