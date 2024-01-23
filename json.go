package json

import (
	"encoding/json"
)

// Marshal returns the JSON encoding of a map
func Marshal[JsonType map[string]string | []map[string]string](content JsonType) (string, error) {
	contentByte, err := json.Marshal(content)
	if err != nil {
		return "", err
	}
	return string(contentByte), nil
}

// Unmarshal returns the map representation of the JSON
func Unmarshal(content string) (any, error) {
	if content[0] == '[' {
		var v []map[string]string
		err := json.Unmarshal([]byte(content), &v)
		if err != nil {
			return v, err
		}
		return v, nil
	}
	var v map[string]string
	err := json.Unmarshal([]byte(content), &v)
	if err != nil {
		return v, err
	}
	return v, nil
}
