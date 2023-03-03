package nidek

import (
	"fmt"
	"math"
	"qpdatagather/dataparser/diopter"
	"strconv"
	"strings"
)

func NidekDataParse(byteSlice []byte) any {
	result := diopter.RefractionData{}
	eyeDataRight := diopter.EyeData{}
	eyeDataLeft := diopter.EyeData{}
	pdEyeData := ""

	if len(byteSlice) == 0 {
		return nil
	}

	pi := 337.5

	var (
		lEyeData  = ""
		rEyeData  = ""
		lkEyeData = ""
		rkEyeData = ""
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

		if len(lineBytes) >= 7 && lineBytes[5] == 0x20 && lineBytes[6] == 0x4C {
			if "" == lkEyeData {
				lkEyeData = parseLKEyeData(lineBytes)
				var leftArray = strings.Split(lkEyeData, ",")
				var r1 = leftArray[0]
				if strings.Index(r1, "0") == 0 {
					r1 = strings.Replace(r1, "0", "", 1)
				}
				eyeDataLeft.R1 = r1
				temp, _ := strconv.ParseFloat(r1, 64)
				eyeDataLeft.D1 = transferData(fmt.Sprintf("%.2f", pi/temp))
				a1, _ := strconv.Atoi(leftArray[2])
				eyeDataLeft.A1 = strconv.Itoa(a1)

				var r2 = leftArray[1]
				if strings.Index(r2, "0") == 0 {
					r2 = strings.Replace(r2, "0", "", 1)
				}
				eyeDataLeft.R2 = r2
				temp2, _ := strconv.ParseFloat(r2, 64)
				eyeDataLeft.D2 = transferData(fmt.Sprintf("%.2f", pi/temp2))
				if a1 > 90 {
					eyeDataLeft.A2 = strconv.Itoa(a1 - 90)
				} else {
					eyeDataLeft.A2 = strconv.Itoa(a1 + 90)
				}

			}
		}
		if lineBytes[0] == 0x20 && lineBytes[1] == 0x52 {
			if "" == rkEyeData {
				rkEyeData = parseRKEyeData(lineBytes)
				var rightArray = strings.Split(rkEyeData, ",")
				var r1 = rightArray[0]
				if strings.Index(r1, "0") == 0 {
					r1 = strings.Replace(r1, "0", "", 1)
				}
				eyeDataRight.R1 = r1
				temp, _ := strconv.ParseFloat(r1, 64)
				eyeDataRight.D1 = transferData(fmt.Sprintf("%.2f", pi/temp))
				a1, _ := strconv.Atoi(rightArray[2])
				eyeDataRight.A1 = strconv.Itoa(a1)

				var r2 = rightArray[1]
				if strings.Index(r2, "0") == 0 {
					r2 = strings.Replace(r2, "0", "", 1)
				}
				eyeDataRight.R2 = r2
				temp2, _ := strconv.ParseFloat(r2, 64)
				eyeDataRight.D2 = transferData(fmt.Sprintf("%.2f", pi/temp2))
				if a1 > 90 {
					eyeDataRight.A2 = strconv.Itoa(a1 - 90)
				} else {
					eyeDataRight.A2 = strconv.Itoa(a1 + 90)
				}

			}
		}

	}

	result.Od = eyeDataRight
	result.Os = eyeDataLeft
	result.R = eyeDataRight
	result.L = eyeDataLeft
	result.Pd = pdEyeData

	return result
}

func parseLKEyeData(lineBytes []byte) string {
	r1Data := make([]byte, 5)
	for i := 0; i < 5; i++ {
		r1Data[i] = lineBytes[i+7]
	}
	r2Data := make([]byte, 5)
	for i := 0; i < 5; i++ {
		r2Data[i] = lineBytes[i+7+5]
	}
	aData := make([]byte, 3)
	for i := 0; i < 3; i++ {
		aData[i] = lineBytes[i+7+5+5]
	}
	return string(r1Data) + "," + string(r2Data) + "," + string(aData)
}
func parseRKEyeData(lineBytes []byte) string {
	r1Data := make([]byte, 5)
	for i := 0; i < 5; i++ {
		r1Data[i] = lineBytes[i+2]
	}
	r2Data := make([]byte, 5)
	for i := 0; i < 5; i++ {
		r2Data[i] = lineBytes[i+2+5]
	}
	aData := make([]byte, 3)
	for i := 0; i < 3; i++ {
		aData[i] = lineBytes[i+2+5+5]
	}
	return string(r1Data) + "," + string(r2Data) + "," + string(aData)
}

func parsePdData(lineBytes []byte) string {
	var pdData = make([]byte, 2)
	for i := 0; i < 2; i++ {
		pdData[i] = lineBytes[i+2]
	}
	return string(pdData)
}

func parseEyeData(lineBytes []byte) string {
	var sData = make([]byte, 6)
	for i := 0; i < 6; i++ {
		sData[i] = lineBytes[i+2]
	}

	var cData = make([]byte, 6)
	for i := 0; i < 6; i++ {
		cData[i] = lineBytes[i+2+6]
	}

	var aData = make([]byte, 3)
	for i := 0; i < 3; i++ {
		aData[i] = lineBytes[i+2+6+6]
	}

	return string(sData) + "," + string(cData) + "," + string(aData)
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

func transferData(s string) string {
	if "-100.00" == s || len(s) == 0 {
		return ""
	}
	if "0.0" == s || "0.00" == s {
		return "0.00"
	}
	if !isDigit(s) {
		return s
	}
	var a = s[0:strings.Index(s, ".")]
	a1, _ := strconv.ParseFloat(a, 64)
	a1 = a1 * 100
	tmp, _ := strconv.ParseFloat(s, 64)
	val := math.Abs(tmp)
	val = val * 100
	aa := val - a1
	list := []float64{0, 25, 50, 75, 100}
	list2 := make([]float64, len(list))
	for i, item := range list {
		k := math.Abs(item - aa)
		list2[i] = k
	}
	cc := minFloatSlice(list2)
	index := indexOf(cc, list2)
	xx := list[index] / 100
	oo := xx + a1/100
	return fmt.Sprintf("%.2f", oo)

}

func indexOf(element float64, data []float64) int {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1 //not found.
}

func minFloatSlice(values []float64) float64 {
	min := values[0] //assign the first element equal to min
	for _, number := range values {
		if number < min {
			min = number
		}
	}
	return min
}

func isDigit(str string) bool {
	if _, err := strconv.ParseFloat(str, 64); err == nil {
		return true
	} else {
		return false
	}
}
