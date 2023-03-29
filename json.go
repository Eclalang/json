package json

import (
	"encoding/json"
)

// Marshal returns the JSON encoding of a map
func Marshal(content any) string {
	contentByte, _ := json.Marshal(content)
	return string(contentByte)
}

// Unmarshal returns the map representation of the JSON
func Unmarshal(content string) any {
	if content[0] == '[' {
		var v []map[string]interface{}
		json.Unmarshal([]byte(content), &v)
		return v
	}
	var v map[string]interface{}
	json.Unmarshal([]byte(content), &v)
	return v
}
