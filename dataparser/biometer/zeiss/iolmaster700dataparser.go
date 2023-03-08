package zeiss

import (
	"encoding/json"
	"qpdatagather/dataparser/biometer"
	"regexp"
	"strings"
)

func IolMaster700DataParse(byteSlice []byte) biometer.BioData {
	result := biometer.BioData{}

	if len(byteSlice) == 0 {
		return result
	}

	eyeDataRight := biometer.EyeData{}
	eyeDataLeft := biometer.EyeData{}

	var obj = string(byteSlice)
	reg, _ := regexp.Compile("<p>.*?</p>")

	//od
	jsonObj := make(map[string]string)
	_ = json.Unmarshal([]byte(obj), &jsonObj)
	odStr := jsonObj["Od"]
	rList := make([]string, 0)
	tSlice := reg.FindAllString(odStr, -1)
	for _, item := range tSlice {
		if item != "<p>(!)</p>" {
			rList = append(rList, item)
		}
	}
	for i := 0; i < len(rList); i++ {
		if "<p>AL:</p>" == rList[i] {
			eyeDataRight.Al = extractData700(rList[i+4])
		}
		if "<p>中央角膜厚度:</p>" == rList[i] {
			eyeDataRight.Cct = extractData700(rList[i+4])
		}
		if "<p>ACD:</p>" == rList[i] {
			eyeDataRight.Ad = extractData700(rList[i+4])
		}
		if "<p>LT:</p>" == rList[i] {
			eyeDataRight.Lt = extractData700(rList[i+4])
		}
		if "<p>WTW:</p>" == rList[i] {
			eyeDataRight.Wtw = extractData700(rList[i+1])
		}
		if "<p>P:</p>" == rList[i] {
			eyeDataRight.Pd = extractData700(rList[i+1])
		}
		if "<p>SE:</p>" == rList[i] {
			eyeDataRight.Se = extractData700(rList[i+4])
		}
		if "<p>K1:</p>" == rList[i] {
			if extractData700(rList[i+4]) == "" {
				eyeDataRight.K1 = ""
			} else {
				eyeDataRight.K1 = extractData700(rList[i+4]) + "D@" + extractData700(rList[i+14]) + "°"
			}
		}
		if "<p>K2:</p>" == rList[i] {
			if extractData700(rList[i+4]) == "" {
				eyeDataRight.K2 = ""
			} else {
				eyeDataRight.K2 = extractData700(rList[i+4]) + "D@" + extractData700(rList[i+14]) + "°"
			}
		}
		if "<p>ΔK:</p>" == rList[i] {
			if extractData700(rList[i+4]) == "" {
				eyeDataRight.Ast = ""
			} else {
				eyeDataRight.Ast = extractData700(rList[i+4]) + "D@" + extractData700(rList[i+14]) + "°"
			}
		}
	}

	//os
	osStr := jsonObj["Os"]
	lList := make([]string, 0)
	ttSlice := reg.FindAllString(osStr, -1)
	for _, item := range ttSlice {
		if item != "<p>(!)</p>" {
			lList = append(lList, item)
		}
	}
	for i := 0; i < len(lList); i++ {
		if "<p>AL:</p>" == lList[i] {
			eyeDataLeft.Al = extractData700(lList[i+4])
		}
		if "<p>中央角膜厚度:</p>" == lList[i] {
			eyeDataLeft.Cct = extractData700(lList[i+4])
		}
		if "<p>ACD:</p>" == lList[i] {
			eyeDataLeft.Ad = extractData700(lList[i+4])
		}
		if "<p>LT:</p>" == lList[i] {
			eyeDataLeft.Lt = extractData700(lList[i+4])
		}
		if "<p>WTW:</p>" == lList[i] {
			eyeDataLeft.Wtw = extractData700(lList[i+1])
		}
		if "<p>P:</p>" == lList[i] {
			eyeDataLeft.Pd = extractData700(lList[i+1])
		}
		if "<p>SE:</p>" == lList[i] {
			eyeDataLeft.Se = extractData700(lList[i+4])
		}
		if "<p>K1:</p>" == lList[i] {
			if extractData700(lList[i+4]) == "" {
				eyeDataLeft.K1 = ""
			} else {
				eyeDataLeft.K1 = extractData700(lList[i+4]) + "D@" + extractData700(lList[i+14]) + "°"
			}
		}
		if "<p>K2:</p>" == lList[i] {
			if extractData700(lList[i+4]) == "" {
				eyeDataLeft.K2 = ""
			} else {
				eyeDataLeft.K2 = extractData700(lList[i+4]) + "D@" + extractData700(lList[i+14]) + "°"
			}
		}
		if "<p>ΔK:</p>" == lList[i] {
			if extractData700(lList[i+4]) == "" {
				eyeDataLeft.Ast = ""
			} else {
				eyeDataLeft.Ast = extractData700(lList[i+4]) + "D@" + extractData700(lList[i+14]) + "°"
			}
		}
	}

	result.Od = eyeDataRight
	result.Os = eyeDataLeft

	return result
}

func extractData700(input string) string {
	if len(input) == 0 {
		return ""
	}
	reg, _ := regexp.Compile("[-+]?\\d+[\\.]?\\d*")
	list := reg.FindAllString(input, -1)
	if len(list) > 0 {
		return list[0]
	} else {
		return ""
	}

	return strings.Replace(input, ",", ".", 1)
}
