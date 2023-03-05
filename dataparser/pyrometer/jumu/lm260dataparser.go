package jumu

import (
	"qpdatagather/dataparser/pyrometer"
	"strings"
)

// 巨目焦度计，波特率机子默认115200
func LM260DataParse(byteSlice []byte) any {

	if len(byteSlice) == 0 {
		return nil
	}

	result := pyrometer.PyroData{}
	eyeDataRight := pyrometer.EyeData{}
	eyeDataLeft := pyrometer.EyeData{}

	var byte22 = readUntilCR(byteSlice)
	for i := 0; i < len(byte22); i++ {
		var lineBytes = byte22[i]
		var parsedData = parseEyeData(lineBytes)
		if strings.Contains(parsedData, "SPH") {
			var rlData = strings.Split(parsedData, "SPH")
			rS := strings.Replace(rlData[0], " ", "", 1)
			lS := strings.Replace(rlData[1], " ", "", 1)

			eyeDataRight.Sph = rS
			eyeDataLeft.Sph = lS
		}
		if strings.Contains(parsedData, "CYL") {
			var rlData = strings.Split(parsedData, "CYL")
			rC := strings.Replace(rlData[0], " ", "", 1)
			lC := strings.Replace(rlData[1], " ", "", 1)

			eyeDataRight.Sph = rC
			eyeDataLeft.Sph = lC
		}
		if strings.Contains(parsedData, "AXS") {
			var rlData = strings.Split(parsedData, "AXS")
			rA := strings.Replace(rlData[0], " ", "", 1)
			lA := strings.Replace(rlData[1], " ", "", 1)

			eyeDataRight.Sph = rA
			eyeDataLeft.Sph = lA
		}
		if strings.Contains(parsedData, "P-X") {
			var rlData = strings.Split(parsedData, "P-X")
			rPxi := strings.Replace(strings.Replace(rlData[0], "l", "", 1), " ", "", 1)
			lPxi := strings.Replace(strings.Replace(rlData[1], "l", "", 1), " ", "", 1)

			eyeDataRight.Pxi = rPxi
			eyeDataLeft.Pxi = lPxi
		}
		if strings.Contains(parsedData, "P-Y") {
			var rlData = strings.Split(parsedData, "P-Y")
			rPyd := strings.Replace(strings.Replace(rlData[0], "D", "", 1), " ", "", 1)
			lPyd := strings.Replace(strings.Replace(rlData[1], "D", "", 1), " ", "", 1)

			eyeDataRight.Pyd = rPyd
			eyeDataLeft.Pyd = lPyd
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

	return string(sData)
}

func readUntilCR(byteSlice []byte) [][]byte {
	var resultT = make([][]byte, 0)
	var byteList = make([]byte, 0)
	for _, b := range byteSlice {
		if b == 0x0A {
			resultT = append(resultT, byteList)
			byteList = make([]byte, 0)
		} else {
			byteList = append(byteList, b)
		}
	}
	return resultT
}
