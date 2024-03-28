package templateRequestModel

type ModifyAppFieldReq struct {
	FieldKey  string `json:"fieldKey,omitempty"`
	FieldName string `json:"fieldName,omitempty"`
}
