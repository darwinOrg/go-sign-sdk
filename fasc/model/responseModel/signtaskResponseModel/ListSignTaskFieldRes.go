package signtaskResponseModel

type ListSignTaskFieldRes struct {
	RequestId string                `json:"requestId"`
	Code      string                `json:"code"`
	Msg       string                `json:"msg"`
	Data      ListSignTaskFieldData `json:"data"`
}

type ListSignTaskFieldData struct {
	SignTaskId      string          `json:"signTaskId,omitempty"`
	SignTaskSubject string          `json:"signTaskSubject,omitempty"`
	FillFields      []FillFieldInfo `json:"fillFields,omitempty"`
}

type FillFieldInfo struct {
	FieldId    string `json:"fieldId,omitempty"`
	FieldName  string `json:"fieldName,omitempty"`
	FieldType  string `json:"fieldType,omitempty"`
	FieldValue string `json:"fieldValue,omitempty"`
	ActorId    string `json:"actorId,omitempty"`
	ActorName  string `json:"actorName,omitempty"`
	DocId      string `json:"docId,omitempty"`
	DocName    string `json:"docName,omitempty"`
}
