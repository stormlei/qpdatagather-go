package tianle

import (
	"qpdatagather/dataparser"
	"regexp"
	"strconv"
	"strings"
)

func Kr9800DataParse(byteSlice []byte) dataparser.RefractionData {
	var result dataparser.RefractionData
	var eyeDataRight dataparser.EyeData
	var eyeDataLeft dataparser.EyeData

	var obj = string(byteSlice)
	var rArray = strings.Split(obj, "E")
	for i := 0; i < len(rArray); i++ {
		if strings.HasPrefix(rArray[i], "S-R-R") {
			reg, _ := regexp.Compile("[-+]?\\d+[\\.]?\\d*")
			rList := reg.FindAllString(rArray[i], -1)
			var s = rList[0]
			if strings.Index(s, "0") == 1 {
				s = strings.Replace(s, "0", "", 1)
			}
			var c = rList[1]
			if strings.Index(c, "0") == 1 {
				c = strings.Replace(c, "0", "", 1)
			}
			var a = rList[2]
			if strings.Index(a, "0") == 1 {
				a = strings.Replace(a, "0", "", 1)
			}
			eyeDataRight = dataparser.EyeData{}
			eyeDataRight.S = s
			eyeDataRight.C = c
			aInt, _ := strconv.Atoi(a)
			eyeDataRight.A = strconv.Itoa(aInt)
		}
		if strings.HasPrefix(rArray[i], "S-R-L") {
			reg, _ := regexp.Compile("[-+]?\\d+[\\.]?\\d*")
			lList := reg.FindAllString(rArray[i], -1)
			var s = lList[0]
			if strings.Index(s, "0") == 1 {
				s = strings.Replace(s, "0", "", 1)
			}
			var c = lList[1]
			if strings.Index(c, "0") == 1 {
				c = strings.Replace(c, "0", "", 1)
			}
			var a = lList[2]
			if strings.Index(a, "0") == 1 {
				a = strings.Replace(a, "0", "", 1)
			}
			eyeDataLeft = dataparser.EyeData{}
			eyeDataLeft.S = s
			eyeDataLeft.C = c
			aInt, _ := strconv.Atoi(a)
			eyeDataLeft.A = strconv.Itoa(aInt)
		}
	}

	result = dataparser.RefractionData{}
	result.Od = eyeDataRight
	result.Os = eyeDataLeft
	result.R = eyeDataRight
	result.L = eyeDataLeft
	return result
}
