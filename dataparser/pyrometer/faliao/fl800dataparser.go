package faliao

import (
	"qpdatagather/dataparser/pyrometer"
	"strings"
)

func FL800DataParse(byteSlice []byte) pyrometer.PyroData {
	result := pyrometer.PyroData{}
	eyeDataRight := pyrometer.EyeData{}
	eyeDataLeft := pyrometer.EyeData{}

	if len(byteSlice) == 0 {
		return result
	}

	var obj = string(byteSlice)
	if strings.Contains(obj, "R") && strings.Contains(obj, "L") {
		var rlData = strings.Split(obj, ",L,")
		//右眼数据
		var rData = rlData[0]
		var (
			rs = ""
			rc = ""
			ra = ""
		)
		if strings.Contains(rData, "SPH") {
			var pos = strings.Index(rData, "SPH")
			rs = rData[pos+4 : pos+10]
			if strings.Index(rs, "0") == 1 {
				rs = strings.Replace(rs, "0", "", 1)
			}
		}
		if strings.Contains(rData, "CYL") {
			var pos = strings.Index(rData, "CYL")
			rc = rData[pos+4 : pos+10]
			if strings.Index(rc, "0") == 1 {
				rc = strings.Replace(rc, "0", "", 1)
			}
		}
		if strings.Contains(rData, "AXS") {
			var pos = strings.Index(rData, "AXS")
			rc = rData[pos+4 : pos+10]
			if strings.Index(rc, "0") == 1 {
				rc = strings.Replace(rc, "0", "", 1)
			}
		}
		eyeDataRight.Sph = rs
		eyeDataRight.Cyl = rc
		eyeDataRight.Axs = ra

		//左眼数据
		var lData = rlData[1]
		var (
			ls = ""
			lc = ""
			la = ""
		)
		if strings.Contains(lData, "SPH") {
			var pos = strings.Index(lData, "SPH")
			ls = lData[pos+4 : pos+10]
			if strings.Index(ls, "0") == 1 {
				ls = strings.Replace(ls, "0", "", 1)
			}
		}
		if strings.Contains(lData, "CYL") {
			var pos = strings.Index(lData, "CYL")
			lc = lData[pos+4 : pos+10]
			if strings.Index(lc, "0") == 1 {
				lc = strings.Replace(lc, "0", "", 1)
			}
		}
		if strings.Contains(lData, "AXS") {
			var pos = strings.Index(lData, "AXS")
			lc = lData[pos+4 : pos+10]
			if strings.Index(lc, "0") == 1 {
				lc = strings.Replace(lc, "0", "", 1)
			}
		}
		eyeDataLeft.Sph = ls
		eyeDataLeft.Cyl = lc
		eyeDataLeft.Axs = la

	}

	result.Od = eyeDataRight
	result.Os = eyeDataLeft
	result.R = eyeDataRight
	result.L = eyeDataLeft

	return result
}
