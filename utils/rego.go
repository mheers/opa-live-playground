package utils

import "encoding/json"

func BeautifyJson(in string) (string, error) {
	inMap := map[string]any{}
	if err := json.Unmarshal([]byte(in), &inMap); err != nil {
		return "", err
	}

	return MarshalInBeautified(inMap)
}

func MarshalInBeautified(in interface{}) (string, error) {
	inBeautified, err := json.MarshalIndent(in, "", "  ")
	if err != nil {
		return "", err
	}

	return string(inBeautified), nil
}
