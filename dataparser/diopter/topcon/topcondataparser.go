package topcon

import (
	"qpdatagather/dataparser/diopter"
	"strings"
)

var (
	byteStart = 0x41
	byteR     = 0x52
	byteL     = 0x4c
	byteCR    = 0x0d
	byteEnd   = 0x04
)

func TopconDataParse(byteSlice []byte) any {

	if len(byteSlice) == 0 {
		return nil
	}

	if int(byteSlice[0]) != byteStart || int(byteSlice[len(byteSlice)-1]) != byteEnd {
		return nil
	}

	var byte22 = readUntilCR(byteSlice)

	//读出头部两个字节，//头部数量必须为1
	var headBytes = byte22[0]
	if len(headBytes) != 1 {
		return nil
	}
	//当前眼
	var eyeByte = headBytes[0]

	result := diopter.RefractionData{}
	eyeDataRight := diopter.EyeData{}
	eyeDataLeft := diopter.EyeData{}
	pd := ""

	for i := 0; i < len(byte22); i++ {
		var lineBytes = byte22[i]

		var readSize = len(lineBytes)
		var parsedData = parseEyeData(lineBytes)
		switch readSize {
		case 6:
			if int(eyeByte) == byteR {
				eyeDataRight.S = parsedData
			} else {
				eyeDataLeft.S = parsedData
			}
		case 5:
			if int(eyeByte) == byteR {
				eyeDataRight.C = parsedData
			} else {
				eyeDataLeft.C = parsedData
			}
		case 3:
			if int(eyeByte) == byteR {
				eyeDataRight.A = parsedData
			} else {
				eyeDataLeft.A = parsedData
			}
		case 2:
			pd = parsedData
		case 1:
			//只读出来一个byte, 则为R,L,或者空格
			var singleByte = lineBytes[0]
			if int(singleByte) == byteR || int(singleByte) == byteL {
				eyeByte = singleByte
			}
		}

	}

	result.Od = eyeDataRight
	result.Os = eyeDataLeft
	result.R = eyeDataRight
	result.L = eyeDataLeft
	result.Pd = pd

	return result
}

func parseEyeData(lineBytes []byte) string {
	var length = len(lineBytes)
	var sData = make([]byte, length)
	for i := 0; i < length; i++ {
		sData[i] = lineBytes[i]
	}

	return strings.ReplaceAll(string(sData), " ", "")
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
