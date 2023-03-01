package meiwo

import (
	"encoding/json"
	"fmt"
	"math"
	"net/url"
	"qpdatagather/dataparser"
	"strconv"
	"strings"
)

func V100DataParse(byteSlice []byte) any {
	result := dataparser.RefractionData{}
	eyeDataRight := dataparser.EyeData{}
	eyeDataLeft := dataparser.EyeData{}

	if len(byteSlice) == 0 {
		return nil
	}

	var str = string(byteSlice)
	var strSlice = strings.Split(str, "\\r\\n")
	var content = strSlice[len(strSlice)-1]

	var startPos = strings.Index(content, "{")
	if startPos == -1 {
		return nil
	}
	var data = content[startPos:]
	tempMap := make(map[string]string)
	_ = json.Unmarshal([]byte(data), &tempMap)

	var rSphT, _ = url.QueryUnescape(tempMap["SphereRight"])
	var rSph = transferData(rSphT)
	var rCyl = transferData(tempMap["CylinderRight"])
	var rAx = tempMap["AxisRight"]

	eyeDataRight.S = rSph
	eyeDataRight.C = rCyl
	eyeDataRight.A = rAx

	var lSphT, _ = url.QueryUnescape(tempMap["SphereLeft"])
	var lSph = transferData(lSphT)
	var lCyl = transferData(tempMap["CylinderLeft"])
	var lAx = tempMap["AxisLeft"]

	eyeDataLeft.S = lSph
	eyeDataLeft.C = lCyl
	eyeDataLeft.A = lAx

	result.Od = eyeDataRight
	result.Os = eyeDataLeft

	return result
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
	var a = s[1:2]
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
	if tmp > 0 {
		return "+" + fmt.Sprintf("%.2f", oo)
	} else {
		return "-" + fmt.Sprintf("%.2f", oo)
	}

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
