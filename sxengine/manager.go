package sxengine

import (
	"database/sql"
	"time"

	"secretaria.admin/indicadores/helpers"
	. "secretaria.admin/indicadores/sxengine/types"
)

type Frecuencia int

const (
	SEMANAL Frecuencia = iota
	MENSUAL
	TRIMESTRAL
	SEMESTRAL
	ANUAL
)

func (f Frecuencia) String() string {
	return [...]string{"SEMANAL", "MENSUAL", "TRIMESTRAL", "SEMESTRAL", "ANUAL"}[f]
}

type GestionDepartamental struct {
	_db          *sql.DB
	dependencias []Dependencia
}
type GestionDeIndicadores struct {
	_db         *sql.DB
	dimensiones []Dimension
}
type AlmacenDeRegistros struct {
	_db       *sql.DB
	registros []Registro
}
type GestorDeIndicadores struct {
	GestionDepartamental          GestionDepartamental
	GestionDeIndicadores          GestionDeIndicadores
	RelacionDepartamentoIndicador []DepartamentoIndicador
	_db                           *sql.DB
}

func NewGestionDepartamental(db *sql.DB) GestionDepartamental {
	return GestionDepartamental{
		_db:          db,
		dependencias: []Dependencia{},
	}
}
func NewGestionDeIndicadores(db *sql.DB) GestionDeIndicadores {
	return GestionDeIndicadores{
		_db:         db,
		dimensiones: []Dimension{},
	}
}
func NewGestorDeIndicadores(db *sql.DB) GestorDeIndicadores {
	return GestorDeIndicadores{
		GestionDepartamental: NewGestionDepartamental(db),
		GestionDeIndicadores: NewGestionDeIndicadores(db),
		_db:                  db,
	}
}
func (manager *GestorDeIndicadores) AsignarIndicador(idDepartamento int, idIndicador int, fechaInicio time.Time) {

	departamento, err := manager.GestionDepartamental.GetDepartamento(idDepartamento)
	if err != nil {
		panic(err)
	}
	indicador, err := manager.GestionDeIndicadores.GetIndicadorById(idIndicador)
	if err != nil {
		panic(err)
	}
	manager.RelacionDepartamentoIndicador = append(manager.RelacionDepartamentoIndicador, DepartamentoIndicador{
		IdDepartamento: idDepartamento,
		IdIndicador:    idIndicador,
		Indicador:      indicador,
		Departamento:   departamento,
		FechaInicio:    fechaInicio,
	})

}
func (manager *GestorDeIndicadores) GetIndicadoresDeDepartamento(idDepartamento int) []Indicador {
	indicadores := []Indicador{}
	for _, relacion := range manager.RelacionDepartamentoIndicador {
		if relacion.IdDepartamento == idDepartamento {
			println("Relacion: ")
			helpers.PrintAsJson(relacion)
			println("============")
			indicador, err := manager.GestionDeIndicadores.GetIndicadorById(relacion.Indicador.IdIndicador)
			indicador.Dimension, err = manager.GestionDeIndicadores.GetDimensionById(indicador.IdDimension)
			if err != nil {
				panic(err)
			}
			indicadores = append(indicadores, indicador)

		}
	}
	return indicadores
}

func (gdi *GestorDeIndicadores) ShowReport() {
	dependencias := gdi.GestionDepartamental.GetDependencias()
	for _, dependencia := range dependencias {
		println("Dependencia: ", dependencia.Nombre)
		println("====================================")
		areas := dependencia.Areas
		for _, area := range areas {
			println("Area: ", area.Nombre)
			println("====================================")
			departamentos := area.Departamentos
			for _, departamento := range departamentos {
				indicadores := gdi.GetIndicadoresDeDepartamento(departamento.IdDepartamento)
				if len(indicadores) == 0 {
					println(departamento.Nombre + ": No hay indicadores en este departamento")
					continue
				}
				println("Departamento: ", departamento.Nombre)
				for _, indicador := range indicadores {
					println("Dimension: ", indicador.Dimension.Nombre)
					println("====================================")
					variables := indicador.Variables
					if len(variables) == 0 {
						println("Indicador: ")
						println(">>>>>> ", indicador.Nombre)
						println("Variables: ")
						println("<<<<<< No hay variables")
						println("====================================")
					}
					if len(variables) > 0 {
						println("Indicador: ")
						println(">>>>>> ", indicador.Nombre)
						println("Variables: ")
						for _, variable := range variables {
							println("<<<<<< ", variable.Nombre)
						}
						println("====================================")
					}
				}
				println("====================================")
			}
			println("====================================")
		}
	}
}
