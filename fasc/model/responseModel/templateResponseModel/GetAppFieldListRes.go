package templateResponseModel

type GetAppFieldListRes struct {
	RequestId string                `json:"requestId,omitempty"`
	Code      string                `json:"code,omitempty"`
	Msg       string                `json:"msg,omitempty"`
	Data      []GetAppFieldListData `json:"data,omitempty"`
}

type GetAppFieldListData struct {
	FieldKey    string `json:"fieldKey"`
	FieldName   string `json:"fieldName"`
	FieldType   string `json:"fieldType"`
	FieldStatus string `json:"fieldStatus"`
}
