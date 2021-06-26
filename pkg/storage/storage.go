package storage

import "fmt"

type Data struct {
	Key   string `json:"Key"`
	Value string `json:"Value"`
}

func (d Data) String() string {
	return fmt.Sprintf("Key: %v, Value: %v", d.Key, d.Value)
}

var DataMap map[string]*Data
