package commonModel

type FieldIdCard struct {
	Required     bool   `json:"required"`
	DefaultValue string `json:"defaultValue,omitempty"`
	Tips         string `json:"tips,omitempty"`
	Width        int    `json:"width,omitempty"`
	Height       int    `json:"height,omitempty"`
	FontType     string `json:"fontType,omitempty"`
	FontSize     int    `json:"fontSize,omitempty"`
	Alignment    string `json:"alignment,omitempty"`
}
