package util

import (
	"bytes"
	"encoding/json"
)

func ToByteWithJSON(data interface{}) ([]byte, error) {
	buffer := &bytes.Buffer{}
	err := json.NewEncoder(buffer).Encode(data)
	return buffer.Bytes(), err
}
