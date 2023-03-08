package yuwell

import (
	"fmt"
	"math"
	"qpdatagather/dataparser/bloodpressure"
)

func YE900DataParse(byteSlice []byte) bloodpressure.BPData {
	bpData := bloodpressure.BPData{}

	if len(byteSlice) == 0 {
		return bpData
	}

	var (
		sys = ""
		dia = ""
		pl  = ""
	)

	mValue = byteSlice

	sysF := getFloatValue(0x32, 1)
	diaF := getFloatValue(0x32, 3)
	plF := getFloatValue(0x32, 14)

	sys = fmt.Sprintf("%f", sysF)
	dia = fmt.Sprintf("%f", diaF)
	pl = fmt.Sprintf("%f", plF)

	bpData.Sys = sys
	bpData.Dia = dia
	bpData.Pl = pl

	return bpData
}

var mValue []byte

func getFloatValue(formatType int, offset int) float64 {
	if (offset + getTypeLen(formatType)) > size() {
		return 0
	}
	if formatType == 0x32 {
		if mValue[offset+1] == 0x07 && mValue[offset] == 0xFE {
			return 0
		}
		if mValue[offset+1] == 0x07 && mValue[offset] == 0xFF ||
			mValue[offset+1] == 0x08 && mValue[offset] == 0x00 ||
			mValue[offset+1] == 0x08 && mValue[offset] == 0x01 {
			return 0
		}
		if mValue[offset+1] == 0x08 && mValue[offset] == 0x02 {
			return 0
		}

		return bytesToFloat(mValue[offset], mValue[offset+1])
	}
	return 0
}

func bytesToFloat(b0 byte, b1 byte) float64 {
	var mantissa = unsignedToSigned(unsignedByteToInt(b0)+((unsignedByteToInt(b1)&0x0F)<<8), 12)
	var exponent = unsignedToSigned(unsignedByteToInt(b1)>>4, 4)
	return float64(mantissa) * math.Pow10(exponent)
}

func unsignedToSigned(unsigned int, size int) int {
	if (unsigned & (1<<size - 1)) != 0 {
		unsigned = -1 * ((1<<size - 1) - (unsigned & ((1<<size - 1) - 1)))
	}
	return unsigned
}

func unsignedByteToInt(b byte) int {
	return int(b)
}

func size() int {
	if mValue == nil {
		return 0
	}
	return len(mValue)
}

func getTypeLen(formatType int) int {
	return formatType & 0xF
}
