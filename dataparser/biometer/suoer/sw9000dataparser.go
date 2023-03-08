package suoer

import (
	"encoding/json"
	"qpdatagather/dataparser/biometer"
	"strings"
)

func SW9000DataParse(byteSlice []byte) biometer.BioData {
	result := biometer.BioData{}
	//od
	eyeDataRight := biometer.EyeData{}
	//os
	eyeDataLeft := biometer.EyeData{}

	if len(byteSlice) == 0 {
		return result
	}

	var obj = string(byteSlice)
	if len(obj) >= 6 {
		obj = obj[0 : len(obj)-6]
	} else {
		obj = ""
	}

	if strings.Contains(obj, "Data_OD") && strings.Contains(obj, "Data_OS") {
		jsonObj := make(map[string]string)
		_ = json.Unmarshal([]byte(obj), &jsonObj)
		var (
			odAL  = ""
			odCCT = ""
			odAD  = ""
			odLT  = ""
			odVT  = ""
			odK1  = ""
			odK2  = ""
			odAST = ""
			odPD  = ""
			odWTW = ""
		)
		//右眼
		odObj := make(map[string]string)
		_ = json.Unmarshal([]byte(jsonObj["Data_OD"]), &odObj)
		if odObj != nil {
			odAL = odObj["AL"]
			odCCT = odObj["CCT"]
			odAD = odObj["AD"]
			odLT = odObj["LT"]
			odVT = odObj["VT"]

			odK1 = odObj["K1"] + "D@" + odObj["A1"] + "°"
			odK2 = odObj["K2"] + "D@" + odObj["A2"] + "°"
			odAST = odObj["AST"] + "D@" + odObj["A1"] + "°"
			odPD = odObj["PD"]
			odWTW = odObj["WTW"]
		}

		eyeDataRight.Al = odAL
		eyeDataRight.Cct = odCCT
		eyeDataRight.Ad = odAD
		eyeDataRight.Lt = odLT
		eyeDataRight.Vt = odVT
		eyeDataRight.K1 = odK1
		eyeDataRight.K2 = odK2
		eyeDataRight.Ast = odAST
		eyeDataRight.Pd = odPD
		eyeDataRight.Wtw = odWTW

		var (
			osAL  = ""
			osCCT = ""
			osAD  = ""
			osLT  = ""
			osVT  = ""
			osK1  = ""
			osK2  = ""
			osAST = ""
			osPD  = ""
			osWTW = ""
		)
		//左眼
		osObj := make(map[string]string)
		_ = json.Unmarshal([]byte(jsonObj["Data_OS"]), &osObj)
		if osObj != nil {
			osAL = osObj["AL"]
			osCCT = osObj["CCT"]
			osAD = osObj["AD"]
			osLT = osObj["LT"]
			osVT = osObj["VT"]

			osK1 = osObj["K1"] + "D@" + osObj["A1"] + "°"
			osK2 = osObj["K2"] + "D@" + osObj["A2"] + "°"
			osAST = osObj["AST"] + "D@" + osObj["A1"] + "°"
			osPD = osObj["PD"]
			osWTW = osObj["WTW"]
		}

		eyeDataLeft.Al = osAL
		eyeDataLeft.Cct = osCCT
		eyeDataLeft.Ad = osAD
		eyeDataLeft.Lt = osLT
		eyeDataLeft.Vt = osVT
		eyeDataLeft.K1 = osK1
		eyeDataLeft.K2 = osK2
		eyeDataLeft.Ast = osAST
		eyeDataLeft.Pd = osPD
		eyeDataLeft.Wtw = osWTW

	}

	result.Od = eyeDataRight
	result.Os = eyeDataLeft

	return result
}
