package nidek

import (
	"qpdatagather/dataparser/diopter"
	"strings"
)

func RT5100DataParse(byteSlice []byte) diopter.RefractionData {
	result := diopter.RefractionData{}
	eyeDataRight := diopter.EyeData{}
	eyeDataLeft := diopter.EyeData{}

	if len(byteSlice) == 0 {
		return result
	}

	var str = strings.Replace(strings.Replace(strings.Replace(string(byteSlice), "RT", "", 1), "RM", "", 1), "LM", "", 1)

	var rPos = strings.Index(str, "R")
	var rStr = str[rPos+1 : rPos+16]

	var rSph = strings.Replace(rStr[0:6], " ", "", 1)
	var rCyl = strings.Replace(rStr[6:12], " ", "", 1)
	var rAx = strings.Replace(rStr[12:15], " ", "", 1)

	eyeDataRight.S = rSph
	eyeDataRight.C = rCyl
	eyeDataRight.A = rAx

	var lPos = strings.Index(str, "L")
	var lStr = str[lPos+1 : lPos+16]
	var lSph = strings.Replace(lStr[0:6], " ", "", 1)
	var lCyl = strings.Replace(lStr[6:12], " ", "", 1)
	var lAx = strings.Replace(lStr[12:15], " ", "", 1)

	eyeDataLeft.S = lSph
	eyeDataLeft.C = lCyl
	eyeDataLeft.A = lAx

	result.Od = eyeDataRight
	result.Os = eyeDataLeft

	return result
}
