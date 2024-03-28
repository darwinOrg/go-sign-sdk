package corpManageResponseModel

type AuthUrlResData struct {
	AuthUrl string `json:"authUrl"`
}

type GetCorpAuthUrlRes struct {
	RequestId string         `json:"requestId"`
	Code      string         `json:"code"`
	Msg       string         `json:"msg"`
	Data      AuthUrlResData `json:"data"`
}
