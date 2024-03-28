package commonModel

type FieldRemarkSign struct {
	DefaultValue string `json:"defaultValue,omitempty"`
	Tips         string `json:"tips,omitempty"`
	Editable     bool   `json:"editable"`
	FontType     string `json:"fontType,omitempty"`
	FontSize     int    `json:"fontSize,omitempty"`
	Width        int    `json:"width,omitempty"`
	Height       int    `json:"height,omitempty"`
}
