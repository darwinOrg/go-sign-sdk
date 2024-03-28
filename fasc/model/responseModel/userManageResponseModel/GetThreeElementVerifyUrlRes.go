package userManageResponseModel

type GetThreeElementVerifyUrlRes struct {
	RequestId string                       `json:"requestId"`
	Code      string                       `json:"code"`
	Msg       string                       `json:"msg"`
	Data      GetThreeElementVerifyUrlData `json:"data"`
}

type GetThreeElementVerifyUrlData struct {
	VerifyUrl string `json:"verifyUrl,omitempty"`
}
