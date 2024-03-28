package commonModel

type FieldTable struct {
	Required       bool       `json:"required"`
	Header         []string   `json:"header,omitempty"`
	RequiredCount  int        `json:"requiredCount,omitempty"`
	FontType       string     `json:"fontType,omitempty"`
	FontSize       int        `json:"fontSize,omitempty"`
	Alignment      string     `json:"alignment,omitempty"`
	HeaderPosition string     `json:"headerPosition,omitempty"`
	Rows           int        `json:"rows,omitempty"`
	Cols           int        `json:"cols,omitempty"`
	RowHeight      int        `json:"rowHeight,omitempty"`
	Widths         []int      `json:"widths,omitempty"`
	DynamicFilling bool       `json:"dynamicFilling"`
	DefaultValue   [][]string `json:"defaultValue,omitempty"`
	HideHeader     bool       `json:"hideHeader"`
}
