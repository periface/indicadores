package sxengine


type Dimension struct {
	Nombre      string      `json:"Nombre"`
	Sentido     string      `json:"Sentido"`
	IdDimension int         `json:"IdDimension"`
	Frecuencia  string      `json:"Frecuencia"`
	Indicadores []Indicador `json:"indicadores"`
}

func (d *Dimension) AddIndicador(idIndicador int, nombre string, codigo string, unidadMedida string, metodoCalculo string) Indicador {
	indicador := Indicador{
		IdIndicador: idIndicador,
		Nombre:      nombre,
		Codigo:      codigo,
        UnidadDeMedida: unidadMedida,
        MetodoDeCalculo: metodoCalculo,
        Dimension: *d,
        IdDimension: d.IdDimension,
        Variables: []Variable{},
	}
	d.Indicadores = append(d.Indicadores, indicador)
	return indicador
}
