package tonometer

type EyeData struct {
	Iop string `json:"iop"`
}
type IopData struct {
	Od EyeData `json:"od"`
	Os EyeData `json:"os"`
	R  EyeData `json:"r"`
	L  EyeData `json:"l"`
}
