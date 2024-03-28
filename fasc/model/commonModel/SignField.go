package commonModel

type SignField struct {
	SignFieldStatus string         `json:"signFieldStatus,omitempty"`
	FieldDocId      string         `json:"fieldDocId,omitempty"`
	FieldId         string         `json:"fieldId,omitempty"`
	FieldName       string         `json:"fieldName,omitempty"`
	SealId          *int64         `json:"sealId,omitempty"`
	Moveable        bool           `json:"moveable"`
	Position        *FieldPosition `json:"position,omitempty"`
}
