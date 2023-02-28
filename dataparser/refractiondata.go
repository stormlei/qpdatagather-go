package dataparser

type EyeData struct {
	S string `json:"s"`
	C string `json:"c"`
	A string `json:"a"`
}
type RefractionData struct {
	Od EyeData `json:"od"`
	Os EyeData `json:"os"`
	R  EyeData `json:"r"`
	L  EyeData `json:"l"`
	Pd string  `json:"pd"`
}
