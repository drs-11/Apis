package utils

import (
	"encoding/json"
	"io"
)

func ConvertJsonToData(r io.Reader, target interface{}) error {
	return json.NewDecoder(r).Decode(target)
}

func ConvertDataToJson(data interface{}) []byte {
	json_data, _ := json.Marshal(data)
	return json_data
}
