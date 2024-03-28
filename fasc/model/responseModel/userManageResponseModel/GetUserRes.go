package userManageResponseModel

type GetUserRes struct {
	RequestId string      `json:"requestId"`
	Code      string      `json:"code"`
	Msg       string      `json:"msg"`
	Data      UserResData `json:"data"`
}

type UserResData struct {
	ClientUserId    string   `json:"clientUserId"`
	OpenUserId      string   `json:"openUserId"`
	BindingStatus   string   `json:"bindingStatus"`
	AuthScope       []string `json:"authScope"`
	IdentStatus     string   `json:"identStatus"`
	AvailableStatus string   `json:"availableStatus"`
}
