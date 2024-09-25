package sxengine

import (
	"errors"
	"strconv"
)

func (gi *GestionDeIndicadores) GetDimensiones() []Dimension {
	return gi.dimensiones
}

func (gi *GestionDeIndicadores) AddDimension(idDimension int, nombre string, frecuencia Frecuencia, sentido string) Dimension {
	dimension := Dimension{
		IdDimension: idDimension,
		Nombre:      nombre,
		Frecuencia:  frecuencia.String(),
		Sentido:     sentido,
	}
	gi.dimensiones = append(gi.dimensiones, dimension)
	return dimension
}

func (gi *GestionDeIndicadores) GetDimensionById(idDimension int) (Dimension, error) {
	for i := range gi.dimensiones {
		if gi.dimensiones[i].IdDimension == idDimension {
			return gi.dimensiones[i], nil
		}
	}
    return Dimension{},
        errors.New("Dimension not found with id: " + strconv.Itoa(idDimension))
}

func (gi *GestionDeIndicadores) AddIndicador(dimension *Dimension, indicador Indicador) Indicador {
	dimension.Indicadores = append(dimension.Indicadores, indicador)
	for i := range gi.dimensiones {
		if gi.dimensiones[i].IdDimension == dimension.IdDimension {
			gi.dimensiones[i] = *dimension
		}
	}
	return indicador
}
func (gi *GestionDeIndicadores) GetIndicadores() []Indicador {
	indicadores := []Indicador{}
	for i := range gi.dimensiones {
		for j := range gi.dimensiones[i].Indicadores {
			indicadores = append(indicadores, gi.dimensiones[i].Indicadores[j])
		}
	}
	return indicadores
}
func (gi *GestionDeIndicadores) GetIndicadorById(idIndicador int) (Indicador, error) {
	for i := range gi.dimensiones {
		for j := range gi.dimensiones[i].Indicadores {
			if gi.dimensiones[i].Indicadores[j].IdIndicador == idIndicador {
				return gi.dimensiones[i].Indicadores[j], nil
			}
		}
	}
	return Indicador{}, errors.New("Indicador not found")
}

func (gi *GestionDeIndicadores) AddVariable(indicador *Indicador, variable Variable) Variable {
	indicador.Variables = append(indicador.Variables, variable)
	for i := range gi.dimensiones {
		for j := range gi.dimensiones[i].Indicadores {
			if gi.dimensiones[i].Indicadores[j].IdIndicador == indicador.IdIndicador {
				gi.dimensiones[i].Indicadores[j] = *indicador
			}
		}
	}
	return variable
}

func (gi *GestionDeIndicadores) GetVariables() []Variable {
	variables := []Variable{}
	for i := range gi.dimensiones {
		for j := range gi.dimensiones[i].Indicadores {
			for k := range gi.dimensiones[i].Indicadores[j].Variables {
				variables = append(variables, gi.dimensiones[i].Indicadores[j].Variables[k])
			}
		}
	}
	return variables
}
func (gi *GestionDeIndicadores) GetVariableById(idVariable int) (Variable, error) {
	for i := range gi.dimensiones {
		for j := range gi.dimensiones[i].Indicadores {
			for k := range gi.dimensiones[i].Indicadores[j].Variables {
				if gi.dimensiones[i].Indicadores[j].Variables[k].IdVariable == idVariable {
					return gi.dimensiones[i].Indicadores[j].Variables[k], nil
				}
			}
		}
	}
	return Variable{}, errors.New("Variable not found")
}
