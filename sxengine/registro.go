package sxengine

type Registro struct {
    IdRegistro int
    IdIndicador int
    IdDepartamento int
    IdVariable int
    idDimension int
    Valor float64

    Indicador Indicador
    Departamento Departamento
    Variable Variable
    Dimension Dimension
}
