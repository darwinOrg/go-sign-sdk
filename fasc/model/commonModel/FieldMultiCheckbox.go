package commonModel

type FieldMultiCheckbox struct {
	Required     bool   `json:"required"`
	DefaultValue []bool `json:"defaultValue,omitempty"`
}
