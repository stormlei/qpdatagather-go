package topcon

import (
	"qpdatagather/dataparser/diopter"
	"strings"
)

func CV5000DataParse(byteSlice []byte) diopter.RefractionData {
	result := diopter.RefractionData{}
	eyeDataRight := diopter.EyeData{}
	eyeDataLeft := diopter.EyeData{}
	pd := ""

	if len(byteSlice) == 0 {
		return result
	}

	var byte22 = readUntilCR(byteSlice)
	for i := 0; i < len(byte22); i++ {
		var lineBytes = byte22[i]
		var parsedData = parseEyeData2(lineBytes)
		if strings.Contains(parsedData, "FR") && len(parsedData) > 2 {
			eyeDataRight.S = strings.Replace(parsedData[2:8], " ", "", 1)
			eyeDataRight.C = strings.Replace(parsedData[8:14], " ", "", 1)
			eyeDataRight.A = parsedData[14:]
		}
		if strings.Contains(parsedData, "FL") && len(parsedData) > 2 {
			eyeDataLeft.S = strings.Replace(parsedData[2:8], " ", "", 1)
			eyeDataLeft.C = strings.Replace(parsedData[8:14], " ", "", 1)
			eyeDataLeft.A = parsedData[14:]
		}
		if strings.Contains(parsedData, "PD") && len(parsedData) > 2 {
			pd = parsedData[2:]
		}

	}

	result.Od = eyeDataRight
	result.Os = eyeDataLeft
	result.R = eyeDataRight
	result.L = eyeDataLeft
	result.Pd = pd

	return result
}

func parseEyeData2(lineBytes []byte) string {
	var length = len(lineBytes)
	var sData = make([]byte, length)
	for i := 0; i < length; i++ {
		sData[i] = lineBytes[i]
	}

	return string(sData)
}
