package diopter

type EyeData struct {
	S  string `json:"s"`
	C  string `json:"c"`
	A  string `json:"a"`
	R1 string `json:"r1"`
	D1 string `json:"d1"`
	A1 string `json:"a1"`
	R2 string `json:"r2"`
	D2 string `json:"d2"`
	A2 string `json:"a2"`
}
type RefractionData struct {
	Od EyeData `json:"od"`
	Os EyeData `json:"os"`
	R  EyeData `json:"r"`
	L  EyeData `json:"l"`
	Pd string  `json:"pd"`
}
