package zeiss

import (
	"qpdatagather/dataparser/biometer"
	"strings"
)

func IolMaster500DataParse(byteSlice []byte) biometer.BioData {
	result := biometer.BioData{}

	if len(byteSlice) == 0 {
		return result
	}

	eyeDataRight := biometer.EyeData{}
	eyeDataLeft := biometer.EyeData{}

	var obj = string(byteSlice)

	items := strings.Split(obj, ";")

	eyeDataRight.Al = extractData(items[5])
	eyeDataLeft.Al = extractData(items[6])

	eyeDataRight.K1 = extractData(items[7]) + "D@" + extractData(items[9]) + "°"
	eyeDataRight.K2 = extractData(items[8]) + "D@" + extractData(items[9]) + "°"

	eyeDataLeft.K1 = extractData(items[10]) + "D@" + extractData(items[12]) + "°"
	eyeDataLeft.K2 = extractData(items[11]) + "D@" + extractData(items[12]) + "°"

	if len(items) > 13 {
		eyeDataRight.Ad = extractData(items[13])
	}
	if len(items) > 14 {
		eyeDataLeft.Ad = extractData(items[14])
	}

	if len(items) > 15 {
		eyeDataRight.Wtw = extractData(items[15])
	}
	if len(items) > 16 {
		eyeDataLeft.Wtw = extractData(items[16]) //越界处理
	}

	result.Od = eyeDataRight
	result.Os = eyeDataLeft

	return result
}

func extractData(input string) string {
	if len(input) == 0 {
		return ""
	}
	return strings.Replace(input, ",", ".", 1)
}
