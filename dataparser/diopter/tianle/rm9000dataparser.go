package tianle

import (
	"qpdatagather/dataparser/diopter"
	"regexp"
	"strconv"
	"strings"
)

func Rm9000DataParse(byteSlice []byte) diopter.RefractionData {
	var result diopter.RefractionData
	var eyeDataRight diopter.EyeData
	var eyeDataLeft diopter.EyeData

	var obj = string(byteSlice)
	var rArray = strings.Split(obj, "E")
	for i := 0; i < len(rArray); i++ {
		if strings.HasPrefix(rArray[i], "S-R-R") {
			reg, _ := regexp.Compile("[-+]?\\d+[\\.]?\\d+")
			rList := reg.FindAllString(rArray[i], -1)
			var s = rList[0]
			if strings.Index(s, "0") == 1 {
				s = strings.Replace(s, "0", "", 1)
			}
			var c = rList[1]
			if strings.Index(c, "0") == 1 {
				c = strings.Replace(c, "0", "", 1)
			}
			var a = strings.Replace(rList[2], "+", "", 1)

			eyeDataRight = diopter.EyeData{}
			eyeDataRight.S = s
			eyeDataRight.C = c
			aInt, _ := strconv.Atoi(a)
			eyeDataRight.A = strconv.Itoa(aInt)
		}
		if strings.HasPrefix(rArray[i], "S-R-L") {
			reg, _ := regexp.Compile("[-+]?\\d+[\\.]?\\d+")
			lList := reg.FindAllString(rArray[i], -1)
			var s = lList[0]
			if strings.Index(s, "0") == 1 {
				s = strings.Replace(s, "0", "", 1)
			}
			var c = lList[1]
			if strings.Index(c, "0") == 1 {
				c = strings.Replace(c, "0", "", 1)
			}
			var a = strings.Replace(lList[2], "+", "", 1)

			eyeDataLeft = diopter.EyeData{}
			eyeDataLeft.S = s
			eyeDataLeft.C = c
			aInt, _ := strconv.Atoi(a)
			eyeDataLeft.A = strconv.Itoa(aInt)
		}
	}

	result = diopter.RefractionData{}
	result.Od = eyeDataRight
	result.Os = eyeDataLeft
	result.R = eyeDataRight
	result.L = eyeDataLeft
	return result
}
