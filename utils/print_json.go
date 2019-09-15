package utils

import (
	"bytes"
	"encoding/json"
)

// PrintJSON is a utility function prividing pretty printing any struct data
// that can be represented as JSON
func PrintJSON(data interface{}) (string, error) {
	jsonResult, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	var out bytes.Buffer
	err = json.Indent(&out, jsonResult, "", "    ")
	if err != nil {
		return "", err
	}
	return out.String(), nil
}
