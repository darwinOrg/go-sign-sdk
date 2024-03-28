package orgResponseModel

type GetCorpOrgManageUrlRes struct {
	RequestId string           `json:"requestId"`
	Code      string           `json:"code"`
	Msg       string           `json:"msg"`
	Data      CorpOrgManageUrl `json:"data"`
}

type CorpOrgManageUrl struct {
	ResourceUrl string `json:"resourceUrl,omitempty"`
}
