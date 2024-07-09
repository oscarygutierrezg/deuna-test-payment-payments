package umdw

import (
	"encoding/json"
	"github.com/jmoiron/jsonq"
	"strings"
)

func GetMapFromStruct(in interface{}) (map[string]interface{}, error) {
	var result map[string]interface{}

	inBytes, err := json.Marshal(&in)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(inBytes, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func GetPathFromMap(in map[string]interface{}, path string) (interface{}, error) {
	jq := jsonq.NewQuery(in)
	return jq.Interface(strings.Split(path, ".")...)
}
