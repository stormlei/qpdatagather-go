package tianle

import "qpdatagather/dataparser"

func Kr9800Parse(byteSince []byte) dataparser.RefractionData {
	result := dataparser.RefractionData{}
	eyeDataRight := dataparser.EyeData{}
	eyeDataLeft := dataparser.EyeData{}
	eyeDataRight.S = "-0.25"
	eyeDataRight.C = "-0.55"
	eyeDataRight.A = "-0.65"
	eyeDataLeft.S = "-0.25"
	eyeDataLeft.C = "-0.55"
	eyeDataLeft.A = "-0.65"
	result.Od = eyeDataRight
	result.Os = eyeDataLeft
	result.R = eyeDataRight
	result.L = eyeDataLeft
	return result
}
