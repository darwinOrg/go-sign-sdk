package commonModel

type FieldValueInfo struct {
	FieldDocId int    `json:"fieldDocId,omitempty"`
	FieldId    string `json:"fieldId,omitempty"`
	FieldName  string `json:"fieldName,omitempty"`
	FieldValue string `json:"fieldValue,omitempty"`
}
