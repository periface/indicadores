package types
type Variable struct {
	IdVariable  int       `json:"IdVariable"`
	Nombre      string    `json:"Nombre"`
	Codigo      string    `json:"Codigo"`
	IdIndicador int       `json:"IdIndicador"`
	Indicador   Indicador `json:"indicador"`
}
