package orgRequestModel

type GetCorpOrgManageUrlReq struct {
	OpenCorpId   string   `json:"openCorpId,omitempty"`
	ClientUserId string   `json:"clientUserId,omitempty"`
	Modules      []string `json:"modules,omitempty"`
	RedirectUrl  string   `json:"redirectUrl,omitempty"`
}
