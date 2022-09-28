package convertions

import (
	"encoding/json"
	"fmt"
)

func MapToJson(m map[string]int)string {
	json, error := json.Marshal(m)
	if error !=nil {
		fmt.Printf("Error: %s", error.Error())
		return error.Error()
	} else {
		return string(json)
	}
}
func JsonToMap(jsonString []byte )map[string]int {
	var newMap map[string]int
	 json.Unmarshal(jsonString,&newMap)
	return newMap
	
}

func StructToJson(structure interface{})string{
	json, error := json.MarshalIndent(structure,"", "  ") //pretty print the json
	if error !=nil {
		fmt.Printf("Error: %s", error.Error())
		return error.Error()
	} else {
		return string(json)
	}
}
