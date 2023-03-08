package lejia

import (
	"qpdatagather/dataparser/heightweight"
	"strings"
)

func LJ700DataParse(byteSlice []byte) any {
	if len(byteSlice) == 0 {
		return nil
	}

	orderValue := strings.Replace(strings.Replace(string(byteSlice), "OK$", "", 1), "�", "", 1)

	var (
		growthkit_height = ""
		growthkit_weight = ""
	)

	if len(orderValue) >= 14 {
		growthkit_weight = orderValue[1:4] + "." + orderValue[4:5]
		if strings.Contains(growthkit_weight, " ") {
			growthkit_weight = strings.Trim(growthkit_weight, " ")
		}
		if strings.Index(growthkit_weight, "0") == 0 {
			growthkit_weight = strings.Replace(growthkit_weight, "0", "", 1)
		}

		growthkit_height = orderValue[7:10] + "." + orderValue[10:11]
		if strings.Index(growthkit_height, "0") == 0 {
			growthkit_height = strings.Replace(growthkit_height, "0", "", 1)
		}
	}

	hwData := heightweight.HWData{}
	hwData.W = growthkit_weight
	hwData.H = growthkit_height

	return hwData
}
