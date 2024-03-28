package euiResponseModel

type GetAppResourceUrlRes struct {
	RequestId string                `json:"requestId"`
	Code      string                `json:"code"`
	Msg       string                `json:"msg"`
	Data      AppResourceUrlResData `json:"data"`
}

type AppResourceUrlResData struct {
	ResourceUrl string `json:"resourceUrl"`
}
