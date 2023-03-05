package nidek

import (
	"qpdatagather/dataparser/tonometer"
	"strings"
)

func NT5100DataParse(byteSlice []byte) any {
	result := tonometer.IopData{}
	eyeDataRight := tonometer.EyeData{}
	eyeDataLeft := tonometer.EyeData{}

	if len(byteSlice) == 0 {
		return nil
	}

	var (
		lEyeData = ""
		rEyeData = ""
	)

	var byte22 = readUntilCR(byteSlice)
	for i := 0; i < len(byte22); i++ {
		var lineBytes = byte22[i]
		if lineBytes[0] == 0x20 && lineBytes[1] == 0x4C {
			if "" == lEyeData {
				lEyeData = parseEyeData(lineBytes)
				var leftArray = strings.Split(lEyeData, ",")
				var iop = leftArray[0]
				eyeDataLeft.Iop = iop
			}
		}
		if lineBytes[0] == 0x20 && lineBytes[1] == 0x52 {
			if "" == rEyeData {
				rEyeData = parseEyeData(lineBytes)
				var rightArray = strings.Split(rEyeData, ",")
				var iop = rightArray[0]
				eyeDataRight.Iop = iop
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
	var length1 = len(lineBytes) - 9
	var iopData = make([]byte, 4)
	for i := 0; i < 4; i++ {
		iopData[i] = lineBytes[i+length1]
	}

	var length2 = len(lineBytes) - 4
	var kpaData = make([]byte, 4)
	for i := 0; i < 4; i++ {
		kpaData[i] = lineBytes[i+length2]
	}

	return string(iopData) + "," + string(kpaData)
}

func readUntilCR(byteSlice []byte) [][]byte {
	var resultT = make([][]byte, 0)
	var byteList = make([]byte, 0)
	for _, b := range byteSlice {
		if b == 0x17 {
			resultT = append(resultT, byteList)
			byteList = make([]byte, 0)
		} else {
			byteList = append(byteList, b)
		}
	}
	return resultT
}
