package userManageResponseModel

type GetFourElementVerifyUrlRes struct {
	RequestId string                      `json:"requestId"`
	Code      string                      `json:"code"`
	Msg       string                      `json:"msg"`
	Data      GetFourElementVerifyUrlData `json:"data"`
}

type GetFourElementVerifyUrlData struct {
	VerifyUrl string `json:"verifyUrl,omitempty"`
}
