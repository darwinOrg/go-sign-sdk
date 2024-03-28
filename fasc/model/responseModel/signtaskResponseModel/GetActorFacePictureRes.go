package signtaskResponseModel

type GetActorFacePictureRes struct {
	RequestId string                  `json:"requestId"`
	Code      string                  `json:"code"`
	Msg       string                  `json:"msg"`
	Data      GetActorFacePictureData `json:"data"`
}

type GetActorFacePictureData struct {
	Picture string `json:"picture,omitempty"`
}
