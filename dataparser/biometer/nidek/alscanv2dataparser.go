package nidek

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"golang.org/x/net/html/charset"
	"io"
	"qpdatagather/dataparser/biometer"
	"qpdatagather/util"
	"strings"
)

type dataFormat struct {
	Measure struct {
		Type string `xml:"type,attr"`
		AL   struct {
			R struct {
				Typical struct {
					AxialLength string `xml:"AxialLength"`
				} `xml:"Typical"`
			} `xml:"R"`
			L struct {
				Typical struct {
					AxialLength string `xml:"AxialLength"`
				} `xml:"Typical"`
			} `xml:"L"`
		} `xml:"AL"`

		KM []struct {
			Condition string `xml:"condition,attr"`
			R         struct {
				Typical struct {
					R1 struct {
						Power string `xml:"Power"`
						Axis  string `xml:"Axis"`
					} `xml:"R1"`
					R2 struct {
						Power string `xml:"Power"`
						Axis  string `xml:"Axis"`
					} `xml:"R2"`
				} `xml:"Typical"`
			} `xml:"R"`
			L struct {
				Typical struct {
					R1 struct {
						Power string `xml:"Power"`
						Axis  string `xml:"Axis"`
					} `xml:"R1"`
					R2 struct {
						Power string `xml:"Power"`
						Axis  string `xml:"Axis"`
					} `xml:"R2"`
				} `xml:"Typical"`
			} `xml:"L"`
		} `xml:"KM"`

		CCT struct {
			R struct {
				Typical struct {
					CCT string `xml:"CCT"`
				} `xml:"Typical"`
			} `xml:"R"`
			L struct {
				Typical struct {
					CCT string `xml:"CCT"`
				} `xml:"Typical"`
			} `xml:"L"`
		} `xml:"CCT"`

		ACD struct {
			R struct {
				Typical struct {
					ACD string `xml:"ACD"`
				} `xml:"Typical"`
			} `xml:"R"`
			L struct {
				Typical struct {
					ACD string `xml:"ACD"`
				} `xml:"Typical"`
			} `xml:"L"`
		} `xml:"ACD"`

		WTW struct {
			R struct {
				Typical struct {
					WhiteToWhite string `xml:"WhiteToWhite"`
				} `xml:"Typical"`
			} `xml:"R"`
			L struct {
				Typical struct {
					WhiteToWhite string `xml:"WhiteToWhite"`
				} `xml:"Typical"`
			} `xml:"L"`
		} `xml:"WTW"`

		PS struct {
			R struct {
				Typical struct {
					PupilSize string `xml:"PupilSize"`
				} `xml:"Typical"`
			} `xml:"R"`
			L struct {
				Typical struct {
					PupilSize string `xml:"PupilSize"`
				} `xml:"Typical"`
			} `xml:"L"`
		} `xml:"PS"`
	} `xml:"Measure"`
}

func ALScanV2DataParse(byteSlice []byte) biometer.BioData {
	result := biometer.BioData{}
	//od
	eyeDataRight := biometer.EyeData{}
	//os
	eyeDataLeft := biometer.EyeData{}

	if len(byteSlice) == 0 {
		return result
	}

	var obj = string(byteSlice)

	var jsonObj = make(map[string]string)
	var _ = json.Unmarshal([]byte(obj), &jsonObj)

	var xmlStr = jsonObj["Opt"]

	dataSlice, err := util.ReadUTF16([]byte(xmlStr))
	if err != nil {
		fmt.Printf("%s", err)
		return result
	}

	var data dataFormat
	decoder := xml.NewDecoder(bytes.NewBuffer(dataSlice))
	decoder.CharsetReader = func(charsetT string, input io.Reader) (io.Reader, error) {
		return charset.NewReader(input, charsetT)
	}
	err = decoder.Decode(&data)
	if err != nil {
		fmt.Println("Error unmarshalling from XML", err)
		return result
	}

	eyeDataRight.Al = data.Measure.AL.R.Typical.AxialLength
	eyeDataRight.Cct = data.Measure.CCT.R.Typical.CCT
	eyeDataRight.Ad = data.Measure.ACD.R.Typical.ACD
	eyeDataRight.Wtw = data.Measure.WTW.R.Typical.WhiteToWhite
	for _, s := range data.Measure.KM {
		if strings.Contains(s.Condition, "2.4mm") {
			eyeDataRight.K1 = s.R.Typical.R1.Power + "D@" + s.R.Typical.R1.Axis + "°"
			eyeDataRight.K2 = s.R.Typical.R2.Power + "D@" + s.R.Typical.R2.Axis + "°"

			eyeDataLeft.K1 = s.L.Typical.R1.Power + "D@" + s.L.Typical.R1.Axis + "°"
			eyeDataLeft.K2 = s.L.Typical.R2.Power + "D@" + s.L.Typical.R2.Axis + "°"
		}
	}

	eyeDataLeft.Al = data.Measure.AL.L.Typical.AxialLength
	eyeDataLeft.Cct = data.Measure.CCT.L.Typical.CCT
	eyeDataLeft.Ad = data.Measure.ACD.L.Typical.ACD
	eyeDataLeft.Wtw = data.Measure.WTW.L.Typical.WhiteToWhite

	result.Od = eyeDataRight
	result.Os = eyeDataLeft

	return result
}
