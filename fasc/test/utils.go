package test

import (
	"bytes"
	"encoding/json"
)

// json.Marshal 默认 escapeHtml 为true,会转义 <、>、&
func ModelToJsonString(v interface{}, isEscapeHTML bool) string {
	if isEscapeHTML {
		jsonData, _ := json.Marshal(v)
		return string(jsonData)
	} else {
		bf := bytes.NewBuffer([]byte{})
		jsonEncoder := json.NewEncoder(bf)
		jsonEncoder.SetEscapeHTML(false)
		jsonEncoder.Encode(v)
		return bf.String()
	}
}
