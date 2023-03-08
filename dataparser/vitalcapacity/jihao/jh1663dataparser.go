package jihao

import (
	"qpdatagather/dataparser/vitalcapacity"
	"strconv"
	"strings"
)

func JH1663DataParse(byteSlice []byte) vitalcapacity.VCData {
	vcData := vitalcapacity.VCData{}

	if len(byteSlice) == 0 {
		return vcData
	}

	result := string(byteSlice)

	var vc = ""

	if strings.Contains(result, "$FHJ:") {
		beginIndex := strings.LastIndex(result, "$FHJ:") + 5
		endIndex := beginIndex + 4
		vc = result[beginIndex:endIndex]
		vtI, _ := strconv.Atoi(vc)
		vc = strconv.Itoa(vtI)
	}

	vcData.VC = vc

	return vcData
}
