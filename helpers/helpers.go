package helpers

import "encoding/json"

func PrintAsJson(i interface{}) {
	json, err := json.Marshal(i)
	if err != nil {
		println(err)
	}
	println(string(json))
}
