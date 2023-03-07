package nidek

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"golang.org/x/net/html/charset"
	"io"
	"qpdatagather/dataparser/cem"
	"qpdatagather/util"
)

type DataFormat struct {
	Measure struct {
		Type string `xml:"type,attr"`
		SM   struct {
			R struct {
				List struct {
					Num string `xml:"NUM"`
					Cd  string `xml:"CD"`
					Avg string `xml:"AVG"`
					Sd  string `xml:"SD"`
					Cv  string `xml:"CV"`
					Max string `xml:"MAX"`
					Min string `xml:"MIN"`
					Hex string `xml:"HEX"`
					Ct  string `xml:"CT"`
				} `xml:"List"`
			} `xml:"R"`
			L struct {
				List struct {
					Num string `xml:"NUM"`
					Cd  string `xml:"CD"`
					Avg string `xml:"AVG"`
					Sd  string `xml:"SD"`
					Cv  string `xml:"CV"`
					Max string `xml:"MAX"`
					Min string `xml:"MIN"`
					Hex string `xml:"HEX"`
					Ct  string `xml:"CT"`
				} `xml:"List"`
			} `xml:"L"`
		} `xml:"SM"`
	} `xml:"Measure"`
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

	dataSlice, err := util.ReadUTF16([]byte(xmlStr))
	if err != nil {
		fmt.Printf("%s", err)
		return nil
	}

	var data DataFormat
	decoder := xml.NewDecoder(bytes.NewBuffer(dataSlice))
	decoder.CharsetReader = func(charsetT string, input io.Reader) (io.Reader, error) {
		return charset.NewReader(input, charsetT)
	}
	err = decoder.Decode(&data)
	if err != nil {
		fmt.Println("Error unmarshalling from XML", err)
		return nil
	}

	eyeDataRight.Num = data.Measure.SM.R.List.Num
	eyeDataRight.Cd = data.Measure.SM.R.List.Cd
	eyeDataRight.Avg = data.Measure.SM.R.List.Avg
	eyeDataRight.Sd = data.Measure.SM.R.List.Sd
	eyeDataRight.Cv = data.Measure.SM.R.List.Cv
	eyeDataRight.Max = data.Measure.SM.R.List.Max
	eyeDataRight.Min = data.Measure.SM.R.List.Min
	eyeDataRight.Hex = data.Measure.SM.R.List.Hex
	eyeDataRight.Ct = data.Measure.SM.R.List.Ct

	eyeDataLeft.Num = data.Measure.SM.L.List.Num
	eyeDataLeft.Cd = data.Measure.SM.L.List.Cd
	eyeDataLeft.Avg = data.Measure.SM.L.List.Avg
	eyeDataLeft.Sd = data.Measure.SM.L.List.Sd
	eyeDataLeft.Cv = data.Measure.SM.L.List.Cv
	eyeDataLeft.Max = data.Measure.SM.L.List.Max
	eyeDataLeft.Min = data.Measure.SM.L.List.Min
	eyeDataLeft.Hex = data.Measure.SM.L.List.Hex
	eyeDataLeft.Ct = data.Measure.SM.L.List.Ct

	result := cem.CemData{}
	result.Od = eyeDataRight
	result.Os = eyeDataLeft

	return result
}
