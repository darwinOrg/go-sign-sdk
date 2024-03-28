package euiRequestModel

// GetUserResourceUrlReq 获取用户级资源访问链接
type GetUserResourceUrlReq struct {
	OpenCorpId   string        `json:"openCorpId,omitempty"`
	ClientUserId string        `json:"clientUserId,omitempty"`
	Resource     *UserResource `json:"resource,omitempty"`
}

type UserResource struct {
	ResourceId string `json:"resourceId,omitempty"`
	Action     string `json:"action,omitempty"`
	Params     string `json:"params,omitempty"`
}
