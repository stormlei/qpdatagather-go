package topcon

import (
	"qpdatagather/dataparser/tonometer"
	"strings"
)

var (
	byteStart = 0x40
	byteR     = 0x52
	byteL     = 0x4c
	byteCR    = 0x0d
	byteEnd   = 0x04
)

func CT1DataParse(byteSlice []byte) any {

	if len(byteSlice) == 0 {
		return nil
	}

	if int(byteSlice[0]) != byteStart || int(byteSlice[len(byteSlice)-1]) != byteEnd {
		return nil
	}

	result := tonometer.IopData{}
	eyeDataRight := tonometer.EyeData{}
	eyeDataLeft := tonometer.EyeData{}

	var byte22 = readUntilCR(byteSlice)
	for i := 0; i < len(byte22); i++ {
		var lineBytes = byte22[i]

		var readSize = len(lineBytes)
		var parsedData = parseEyeData(lineBytes)
		switch readSize {
		case 9:
			if strings.Contains(parsedData, "R") {
				eyeDataRight.Iop = strings.Replace(parsedData, "R", "", 1)
			} else {
				eyeDataLeft.Iop = strings.Replace(parsedData, "L", "", 1)
			}
		}

	}

	result.Od = eyeDataRight
	result.Os = eyeDataLeft
	result.R = eyeDataRight
	result.L = eyeDataLeft

	return result
}

func parseEyeData(lineBytes []byte) string {
	var length = len(lineBytes)
	var sData = make([]byte, length)
	for i := 0; i < length; i++ {
		sData[i] = lineBytes[i]
	}

	return strings.ReplaceAll(string(sData), "\\s+", "")
}

func readUntilCR(byteSlice []byte) [][]byte {
	var resultT = make([][]byte, 0)
	var byteList = make([]byte, 0)
	for _, b := range byteSlice {
		if int(b) == byteCR {
			resultT = append(resultT, byteList)
			byteList = make([]byte, 0)
		} else {
			byteList = append(byteList, b)
		}
	}
	return resultT
}
