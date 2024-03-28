package templateResponseModel

type AppSignTemplateDetailRes struct {
	RequestId string                    `json:"requestId,omitempty"`
	Code      string                    `json:"code,omitempty"`
	Msg       string                    `json:"msg,omitempty"`
	Data      AppSignTemplateDetailData `json:"data,omitempty"`
}

type AppSignTemplateDetailData struct {
	AppSignTemplateId     string             `json:"appSignTemplateId,omitempty"`
	AppSignTemplateName   string             `json:"appSignTemplateName,omitempty"`
	AppSignTemplateStatus string             `json:"appSignTemplateStatus,omitempty"`
	CertCAOrg             string             `json:"certCAOrg,omitempty"`
	SignInOrder           bool               `json:"signInOrder,omitempty"`
	Docs                  []Doc              `json:"docs,omitempty"`
	Attachs               []Attach           `json:"attachs,omitempty"`
	Actors                []SignTaskTemActor `json:"actors,omitempty"`
}
