package euiResponseModel

// GetUserResourceUrlRes 获取用户级资源访问链接
type GetUserResourceUrlRes struct {
	RequestId string                 `json:"requestId"`
	Code      string                 `json:"code"`
	Msg       string                 `json:"msg"`
	Data      UserResourceUrlResData `json:"data"`
}

type UserResourceUrlResData struct {
	ResourceUrl string `json:"resourceUrl"`
}
