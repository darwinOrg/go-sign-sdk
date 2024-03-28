package signtaskRequestModel

type DocFieldValue struct {
	DocId      string `json:"docId,omitempty"`
	FieldId    string `json:"fieldId,omitempty"`
	FieldKey   string `json:"fieldKey,omitempty"`
	FieldValue string `json:"fieldValue,omitempty"`
}

// FillFieldValuesReq 填写签署任务控件内容
type FillFieldValuesReq struct {
	SignTaskId     string          `json:"signTaskId,omitempty"`
	DocFieldValues []DocFieldValue `json:"docFieldValues,omitempty"`
}
