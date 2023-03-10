package enum

type Type string

const (
	// Optometry 验光仪
	Optometry Type = "optometry"

	// Biometer 生物测量仪
	Biometer Type = "eyeBiometrics"

	// Tonometer 眼压计
	Tonometer Type = "tonometer"

	// Cem 角膜内皮镜
	Cem Type = "cem"

	// Pyrometer 焦度计
	Pyrometer Type = "pyrometer"

	// HeightWeight 身高体重称
	HeightWeight Type = "heightWeight"

	// BloodPressure 血压
	BloodPressure Type = "bloodpressure"

	// VitalCapacity 肺活量
	VitalCapacity Type = "vitalcapacity"
)
