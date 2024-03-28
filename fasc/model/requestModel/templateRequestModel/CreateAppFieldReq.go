package templateRequestModel

type CreateAppFieldReq struct {
	FieldKey  string `json:"fieldKey,omitempty"`
	FieldName string `json:"fieldName,omitempty"`
	FieldType string `json:"fieldType,omitempty"`
}
