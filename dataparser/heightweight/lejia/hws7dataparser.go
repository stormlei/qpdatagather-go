package lejia

import (
	"qpdatagather/dataparser/heightweight"
	"strings"
)

func HWS7DataParse(byteSlice []byte) any {
	if len(byteSlice) == 0 {
		return nil
	}

	orderValue := strings.Replace(string(byteSlice), "OK$", "", 1)

	var (
		growthkit_height = ""
		growthkit_weight = ""
	)

	if len(orderValue) >= 14 {
		growthkit_weight = orderValue[1:4] + "." + orderValue[4:6]
		if strings.Contains(growthkit_weight, " ") {
			growthkit_weight = strings.Trim(growthkit_weight, " ")
		}
		if strings.Index(growthkit_weight, "0") == 0 {
			growthkit_weight = strings.Replace(growthkit_weight, "0", "", 1)
		}

		growthkit_height = orderValue[8:11] + "." + orderValue[11:12]
		if strings.Index(growthkit_height, "0") == 0 {
			growthkit_height = strings.Replace(growthkit_height, "0", "", 1)
		}
	}

	hwData := heightweight.HWData{}
	hwData.W = growthkit_weight
	hwData.H = growthkit_height

	return hwData
}
