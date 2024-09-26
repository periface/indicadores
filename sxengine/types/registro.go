package types

import "time"

type Registro struct {
    IdRegistro int
    IdIndicador int
    IdDepartamento int
    IdVariable int
    idDimension int
    Valor float64
    Fecha time.Time
    IsOpen bool

    Indicador Indicador
    Departamento Departamento
    Variable Variable
    Dimension Dimension
}
