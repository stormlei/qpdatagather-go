package nidek

import (
	"encoding/json"
	"qpdatagather/dataparser/biometer"
)

func ALScanV2DataParse(byteSlice []byte) any {

	if len(byteSlice) == 0 {
		return nil
	}

	//od
	eyeDataRight := biometer.EyeData{}
	//os
	eyeDataLeft := biometer.EyeData{}

	var obj = string(byteSlice)

	var jsonObj = make(map[string]string)
	var _ = json.Unmarshal([]byte(obj), &jsonObj)

	//var xmlStr = strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(jsonObj["Opt"], "\\r", ""), "\\n", ""), "\\t", "")

	result := biometer.BioData{}
	result.Od = eyeDataRight
	result.Os = eyeDataLeft

	return result
}
