package utils

import (
	"encoding/json"
	"io"

	"golang.org/x/crypto/bcrypt"
)

func ConvertJsonToData(r io.Reader, target interface{}) error {
	return json.NewDecoder(r).Decode(target)
}

func ConvertDataToJson(data interface{}) []byte {
	json_data, _ := json.Marshal(data)
	return json_data
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
