package dgsign

import (
	"encoding/json"
	"io"
	"net/http"
)

func mustJsonMarshal(req any) string {
	bytes, _ := json.Marshal(req)
	return string(bytes)
}

func readUrlBody(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}
