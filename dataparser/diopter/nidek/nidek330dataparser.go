package nidek

import (
	"qpdatagather/dataparser/diopter"
	"strconv"
	"strings"
)

func Nidek330DataParse(byteSlice []byte) diopter.RefractionData {
	result := diopter.RefractionData{}
	eyeDataRight := diopter.EyeData{}
	eyeDataLeft := diopter.EyeData{}
	pdEyeData := ""

	if len(byteSlice) == 0 {
		return result
	}

	var (
		lEyeData = ""
		rEyeData = ""
	)

	var byte22 = readUntilCR(byteSlice)
	for i := 0; i < len(byte22); i++ {
		var lineBytes = byte22[i]
		if lineBytes[0] == 0x4F && lineBytes[1] == 0x4C {
			if "" == lEyeData {
				lEyeData = parseEyeData(lineBytes)
				var leftArray = strings.Split(lEyeData, ",")
				var s = leftArray[0]
				if strings.Index(s, "0") == 1 {
					s = strings.Replace(s, "0", "", 1)
				}
				eyeDataLeft.S = s

				var c = leftArray[1]
				if strings.Index(c, "0") == 1 {
					c = strings.Replace(c, "0", "", 1)
				}
				eyeDataLeft.C = c

				var a, _ = strconv.Atoi(leftArray[2])
				eyeDataLeft.A = strconv.Itoa(a)
			}
		}
		if lineBytes[0] == 0x4F && lineBytes[1] == 0x52 {
			if "" == rEyeData {
				rEyeData = parseEyeData(lineBytes)
				var rightArray = strings.Split(rEyeData, ",")
				var s = rightArray[0]
				if strings.Index(s, "0") == 1 {
					s = strings.Replace(s, "0", "", 1)
				}
				eyeDataRight.S = s

				var c = rightArray[1]
				if strings.Index(c, "0") == 1 {
					c = strings.Replace(c, "0", "", 1)
				}
				eyeDataRight.C = c

				var a, _ = strconv.Atoi(rightArray[2])
				eyeDataRight.A = strconv.Itoa(a)
			}
		}
		if lineBytes[0] == 0x50 && lineBytes[1] == 0x44 {
			pdEyeData = parsePdData(lineBytes)
		}

	}

	result.Od = eyeDataRight
	result.Os = eyeDataLeft
	result.R = eyeDataRight
	result.L = eyeDataLeft
	result.Pd = pdEyeData

	return result
}
