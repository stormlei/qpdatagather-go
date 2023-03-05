package topcon

import (
	"qpdatagather/dataparser/tonometer"
	"regexp"
	"strings"
)

var (
	byteStart800 = 0x01
)

func CT800DataParse(byteSlice []byte) any {

	if len(byteSlice) == 0 {
		return nil
	}

	if int(byteSlice[0]) != byteStart800 || int(byteSlice[len(byteSlice)-1]) != byteEnd {
		return nil
	}

	result := tonometer.IopData{}
	eyeDataRight := tonometer.EyeData{}
	eyeDataLeft := tonometer.EyeData{}

	var byte22 = readUntilCR(byteSlice)
	for i := 0; i < len(byte22); i++ {
		var lineBytes = byte22[i]

		var parsedData = parseEyeData(lineBytes)
		reg, _ := regexp.Compile("[^0-9]")
		rList := reg.FindAllString(parsedData, -1)
		for _, item := range rList {
			if item == "R" {
				eyeDataRight.Iop = strings.ReplaceAll(item, " ", "")
			} else {
				eyeDataLeft.Iop = strings.ReplaceAll(item, " ", "")
			}
		}
	}

	result.Od = eyeDataRight
	result.Os = eyeDataLeft
	result.R = eyeDataRight
	result.L = eyeDataLeft

	return result
}
