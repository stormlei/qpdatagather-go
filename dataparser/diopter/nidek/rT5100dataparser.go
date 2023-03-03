package nidek

import (
	"qpdatagather/dataparser/diopter"
	"strings"
)

func RT5100DataParse(byteSlice []byte) any {
	result := diopter.RefractionData{}
	eyeDataRight := diopter.EyeData{}
	eyeDataLeft := diopter.EyeData{}

	if len(byteSlice) == 0 {
		return nil
	}

	var str = string(byteSlice)

	var rPos = strings.Index(str, "R")
	var rStr = str[rPos+1 : rPos+17]

	var rSph = strings.Replace(rStr[rPos:rPos+6], " ", "", 1)
	var rCyl = strings.Replace(rStr[rPos+6:rPos+12], " ", "", 1)
	var rAx = strings.Replace(rStr[rPos+12:rPos+15], " ", "", 1)

	eyeDataRight.S = rSph
	eyeDataRight.C = rCyl
	eyeDataRight.A = rAx

	var lPos = strings.Index(str, "L")
	var lStr = str[lPos+1 : lPos+17]
	var lSph = strings.Replace(lStr[lPos:lPos+6], " ", "", 1)
	var lCyl = strings.Replace(lStr[lPos+6:lPos+12], " ", "", 1)
	var lAx = strings.Replace(lStr[lPos+12:lPos+15], " ", "", 1)

	eyeDataLeft.S = lSph
	eyeDataLeft.C = lCyl
	eyeDataLeft.A = lAx

	result.Od = eyeDataRight
	result.Os = eyeDataLeft

	return result
}
