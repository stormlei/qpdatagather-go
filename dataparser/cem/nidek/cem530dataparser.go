package nidek

import (
	"encoding/json"
	"qpdatagather/dataparser/cem"
)

func Cem530DataParse(byteSlice []byte) any {

	if len(byteSlice) == 0 {
		return nil
	}

	//od
	eyeDataRight := cem.EyeData{}
	//os
	eyeDataLeft := cem.EyeData{}

	var obj = string(byteSlice)

	var jsonObj = make(map[string]string)
	var _ = json.Unmarshal([]byte(obj), &jsonObj)

	//var xmlStr = strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(jsonObj["Opt"], "\\r", ""), "\\n", ""), "\\t", "")

	result := cem.CemData{}
	result.Od = eyeDataRight
	result.Os = eyeDataLeft

	return result
}
