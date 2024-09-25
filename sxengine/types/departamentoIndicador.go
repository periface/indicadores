package types

import "time"

type DepartamentoIndicador struct {
	IdDepartamento int          `json:"IdDepartamento"`
	Departamento   Departamento `json:"Departamento"`
	IdIndicador    int          `json:"IdIndicador"`
	Indicador      Indicador    `json:"Indicador"`
	FechaInicio    time.Time    `json:"FechaInicio"`
}
