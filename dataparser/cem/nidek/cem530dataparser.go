package nidek

import (
	"encoding/json"
	"encoding/xml"
	"golang.org/x/net/html/charset"
	"io"
	"qpdatagather/dataparser/cem"
	"strings"
)

type xmlData struct {
}

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

	var xmlStr = jsonObj["Opt"]

	decoder := xml.NewDecoder(strings.NewReader(xmlStr))
	decoder.CharsetReader = func(charsetT string, input io.Reader) (io.Reader, error) {
		return charset.NewReader(input, charsetT)
	}
	var data xmlData
	var err = decoder.Decode(&data)
	if err != nil {
		return nil
	}

	result := cem.CemData{}
	result.Od = eyeDataRight
	result.Os = eyeDataLeft

	return result
}
