package enum

type Type string

const (
	// Optometry 验光仪
	Optometry Type = "optometry"

	// Biometer 生物测量仪
	Biometer Type = "eyeBiometrics"

	// Tonometer 眼压计
	Tonometer Type = "tonometer"
)
