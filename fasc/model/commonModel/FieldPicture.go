package commonModel

type FieldPicture struct {
	Required     bool   `json:"required"`
	DefaultValue string `json:"defaultValue,omitempty"`
	Width        int    `json:"width,omitempty"`
	Height       int    `json:"height,omitempty"`
}
