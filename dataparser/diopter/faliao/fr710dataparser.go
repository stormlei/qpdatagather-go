package faliao

import (
	"qpdatagather/dataparser/diopter"
	"strings"
)

func Fr710DataParse(byteSlice []byte) diopter.RefractionData {
	var result diopter.RefractionData
	var eyeDataRight diopter.EyeData
	var eyeDataLeft diopter.EyeData

	if len(byteSlice) == 0 {
		return result
	}

	var str = string(byteSlice)
	var strArr = strings.Split(str, "#!>")
	if strings.Contains(str, "REFR") {
		rStr := strArr[1] + "#!>"
		var (
			s = ""
			c = ""
			a = ""
		)
		if len(rStr) >= 42 {
			s = rStr[26:32]
			c = rStr[32:38]
			a = rStr[38:41]
		}
		if strings.Index(s, "0") == 1 {
			s = strings.Replace(s, "0", "", 1)
		}
		if strings.Index(c, "0") == 1 {
			c = strings.Replace(c, "0", "", 1)
		}
		if strings.Index(a, "0") == 1 {
			a = strings.Replace(a, "0", "", 1)
		}

		eyeDataRight = diopter.EyeData{}
		eyeDataRight.S = s
		eyeDataRight.C = c
		eyeDataRight.A = a
	}
	if strings.Contains(str, "REFL") {
		if len(strArr) > 2 {
			lStr := strArr[2] + "#!>"
			var (
				s = ""
				c = ""
				a = ""
			)
			if len(lStr) >= 42 {
				s = lStr[26:32]
				c = lStr[32:38]
				a = lStr[38:41]
			}
			if strings.Index(s, "0") == 1 {
				s = strings.Replace(s, "0", "", 1)
			}
			if strings.Index(c, "0") == 1 {
				c = strings.Replace(c, "0", "", 1)
			}
			if strings.Index(a, "0") == 1 {
				a = strings.Replace(a, "0", "", 1)
			}

			eyeDataLeft = diopter.EyeData{}
			eyeDataLeft.S = s
			eyeDataLeft.C = c
			eyeDataLeft.A = a
		}
	}

	if strings.Contains(str, "KRTR") {
		if len(strArr) > 3 {
			rStr := strArr[3] + "#!>"
			var (
				r1 = ""
				r2 = ""
				a1 = ""
			)
			if len(rStr) >= 42 {
				r1 = rStr[26:32]
				r2 = rStr[32:38]
				a1 = rStr[38:41]
			}
			if strings.Index(r1, "0") == 1 {
				r1 = strings.Replace(r1, "0", "", 1)
				r1 = strings.Replace(r1, "+", "", 1)
			}
			if strings.Index(r2, "0") == 1 {
				r2 = strings.Replace(r2, "0", "", 1)
				r2 = strings.Replace(r2, "+", "", 1)
			}
			if strings.Index(a1, "0") == 1 {
				a1 = strings.Replace(a1, "0", "", 1)
			}

			eyeDataRight.R1 = r1
			eyeDataRight.R2 = r2
			eyeDataRight.A1 = a1
		}
	}
	if strings.Contains(str, "KRTL") {
		if len(strArr) > 4 {
			rStr := strArr[4] + "#!>"
			var (
				r1 = ""
				r2 = ""
				a1 = ""
			)
			if len(rStr) >= 42 {
				r1 = rStr[26:32]
				r2 = rStr[32:38]
				a1 = rStr[38:41]
			}
			if strings.Index(r1, "0") == 1 {
				r1 = strings.Replace(r1, "0", "", 1)
				r1 = strings.Replace(r1, "+", "", 1)
			}
			if strings.Index(r2, "0") == 1 {
				r2 = strings.Replace(r2, "0", "", 1)
				r2 = strings.Replace(r2, "+", "", 1)
			}
			if strings.Index(a1, "0") == 1 {
				a1 = strings.Replace(a1, "0", "", 1)
			}

			eyeDataLeft.R1 = r1
			eyeDataLeft.R2 = r2
			eyeDataLeft.A1 = a1
		}
	}

	result = diopter.RefractionData{}
	result.Od = eyeDataRight
	result.Os = eyeDataLeft
	result.R = eyeDataRight
	result.L = eyeDataLeft
	return result
}
