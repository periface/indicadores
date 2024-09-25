package types

type Indicador struct {
	IdIndicador     int    `json:"IdIndicador"`
	Nombre          string `json:"Nombre"`
	Descripcion     string `json:"descripcion"`
	MetodoDeCalculo string `json:"MetodoDeCalculo"`
	UnidadDeMedida  string `json:"UnidadDeMedida"`

	IdDimension int        `json:"IdDimension"`
	Dimension   Dimension  `json:"Dimension"`
	Variables   []Variable `json:"Variables"`
	Codigo      string     `json:"Codigo"`
}

func (i *Indicador) SetMetodoCalculo(formula string) {
	i.MetodoDeCalculo = formula
}
func (i *Indicador) AddVariable(idVariable int, nombre string, codigo string) Variable {
	variable := Variable{
		IdVariable:  idVariable,
		Nombre:      nombre,
		Codigo:      codigo,
		IdIndicador: i.IdIndicador,
		Indicador:   *i,
	}
	i.Variables = append(i.Variables, variable)
	return variable
}
