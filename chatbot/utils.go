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
