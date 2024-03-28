package commonModel

type FieldMultiRadio struct {
	Required     bool   `json:"required"`
	DefaultValue []bool `json:"defaultValue,omitempty"`
}
