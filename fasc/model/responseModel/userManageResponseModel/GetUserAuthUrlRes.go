package userManageResponseModel

type GetUserAuthUrlRes struct {
	RequestId string             `json:"requestId"`
	Code      string             `json:"code"`
	Msg       string             `json:"msg"`
	Data      UserAuthUrlResData `json:"data"`
}
type UserAuthUrlResData struct {
	AuthUrl string `json:"authUrl,omitempty"`
	EUrl    string `json:"eUrl,omitempty"`
}
