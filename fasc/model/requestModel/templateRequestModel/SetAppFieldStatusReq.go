package templateRequestModel

type SetAppFieldStatusReq struct {
	FieldKey    string `json:"fieldKey,omitempty"`
	FieldStatus string `json:"fieldStatus,omitempty"`
}
