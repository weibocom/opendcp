package utils

import (
	"encoding/json"
)

var (
	Json = &Helper{}
)

type Helper struct {
}

func (h *Helper) ToArray(text string) ([]interface{}, error) {
	arr := make([]interface{}, 0)
	err := json.Unmarshal([]byte(text), &arr)
	if err != nil {
		return nil, err
	}
	return arr, nil
}

func (h *Helper) ToMap(text string) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	err := json.Unmarshal([]byte(text), &m)
	if err != nil {
		return nil, err
	}
	return m, nil
}
