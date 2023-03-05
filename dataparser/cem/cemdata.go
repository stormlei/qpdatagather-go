package cem

type EyeData struct {
	Num string `json:"num"`
	Cd  string `json:"cd"`
	Avg string `json:"avg"`
	Sd  string `json:"sd"`
	Cv  string `json:"cv"`
	Max string `json:"max"`
	Min string `json:"min"`
	Hex string `json:"hex"`
	Ct  string `json:"ct"`
}
type CemData struct {
	Od EyeData `json:"od"`
	Os EyeData `json:"os"`
}
