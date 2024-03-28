package commonModel

// Field 文档控件
type Field struct {
	FieldId             string               `json:"fieldId,omitempty"`
	FieldName           string               `json:"fieldName,omitempty"`
	FieldKey            string               `json:"fieldKey,omitempty"`
	Position            *FieldPosition       `json:"position,omitempty"`
	Moveable            bool                 `json:"moveable"`
	FieldType           string               `json:"fieldType,omitempty"`
	FieldPersonSign     *FieldPersonSign     `json:"fieldPersonSign,omitempty"`
	FieldCorpSeal       *FieldCorpSeal       `json:"fieldCorpSeal,omitempty"`
	FieldRemarkSign     *FieldRemarkSign     `json:"fieldRemarkSign,omitempty"`
	FieldTextSingleLine *FieldTextSingleLine `json:"fieldTextSingleLine,omitempty"`
	FieldTextMultiLine  *FieldTextMultiLine  `json:"fieldTextMultiLine,omitempty"`
	FieldCheckBox       *FieldCheckBox       `json:"fieldCheckBox,omitempty"`
	FieldNumber         *FieldNumber         `json:"fieldNumber,omitempty"`
	FieldIdCard         *FieldIdCard         `json:"fieldIdCard,omitempty"`
	FieldFillDate       *FieldFillDate       `json:"fieldFillDate,omitempty"`
	FieldMultiRadio     *FieldMultiRadio     `json:"fieldMultiRadio,omitempty"`
	FieldMultiCheckbox  *FieldMultiCheckbox  `json:"fieldMultiCheckbox,omitempty"`
	FieldPicture        *FieldPicture        `json:"fieldPicture,omitempty"`
	FieldSelectBox      *FieldSelectBox      `json:"fieldSelectBox,omitempty"`
	FieldTable          *FieldTable          `json:"fieldTable,omitempty"`
}
