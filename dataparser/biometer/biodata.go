package biometer

type EyeData struct {
	Al  string `json:"al"`
	Cct string `json:"cct"`
	Ad  string `json:"ad"`
	Lt  string `json:"lt"`
	Vt  string `json:"vt"`
	K1  string `json:"k1"`
	K2  string `json:"k2"`
	Ast string `json:"ast"`
	Pd  string `json:"pd"`
	Wtw string `json:"wtw"`
	Se  string `json:"se"`
}
type BioData struct {
	Od EyeData `json:"od"`
	Os EyeData `json:"os"`
}
