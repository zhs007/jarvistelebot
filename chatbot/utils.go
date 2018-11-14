package chatbot

import "encoding/json"

// FormatJSON - format JSON string
func FormatJSON(str string) (string, error) {
	var mapResult map[string]interface{}
	err := json.Unmarshal([]byte(str), &mapResult)
	if err != nil {
		return "", err
	}

	jsonStr, err := json.MarshalIndent(mapResult, "", "  ")
	if err != nil {
		return "", err
	}

	return string(jsonStr), nil
}

// FormatJSONObj - format JSON string
func FormatJSONObj(obj interface{}) (string, error) {
	jsonStr, err := json.MarshalIndent(obj, "", "  ")
	if err != nil {
		return "", err
	}

	return string(jsonStr), nil
}
