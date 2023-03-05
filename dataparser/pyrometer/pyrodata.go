package pyrometer

type EyeData struct {
	Sph string `json:"sph"`
	Cyl string `json:"cyl"`
	Axs string `json:"axs"`
	Pxi string `json:"pxi"`
	Pxo string `json:"pxo"`
	Pyu string `json:"pyu"`
	Pyd string `json:"pyd"`
	Ad1 string `json:"ad1"`
	Ad2 string `json:"ad2"`
	Uv  string `json:"uv"`
	Pd  string `json:"pd"`
}
type PyroData struct {
	Od EyeData `json:"od"`
	Os EyeData `json:"os"`
	R  EyeData `json:"r"`
	L  EyeData `json:"l"`
}
