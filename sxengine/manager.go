package sxengine

import "database/sql"

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
type GestorDeIndicadores struct {
	GestionDepartamental          GestionDepartamental
	GestionDeIndicadores          GestionDeIndicadores
	_db                           *sql.DB
	RelacionDepartamentoIndicador []DepartamentoIndicador
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
func (manager *GestorDeIndicadores) AsignarIndicador(idDepartamento int, idIndicador int) {

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
	})
}
func (manager *GestorDeIndicadores) GetIndicadores(idDepartamento int) []Indicador {
	indicadores := []Indicador{}
	for _, relacion := range manager.RelacionDepartamentoIndicador {
		if relacion.IdDepartamento == idDepartamento {
			indicador, err := manager.GestionDeIndicadores.GetIndicadorById(relacion.IdIndicador)
			if err == nil {
				indicadores = append(indicadores, indicador)
			}
		}
	}
	return indicadores
}

func (gdi *GestorDeIndicadores) ShowReport() {
	dependencias := gdi.GestionDepartamental.GetDependencias()
	for _, dependencia := range dependencias {
		println("====================================")
		println("Dependencia: ", dependencia.Nombre)
		println("====================================")
		areas := dependencia.Areas
		for _, area := range areas {
			println("Area: ", area.Nombre)
			println("====================================")

			departamentos := area.Departamentos
			for _, departamento := range departamentos {
				indicadores := gdi.GetIndicadores(departamento.IdDepartamento)
				if len(indicadores) == 0 {
					println(departamento.Nombre + ": No hay indicadores en este departamento")
					continue
				}
				println("Departamento: ", departamento.Nombre)
				for _, indicador := range indicadores {

					variables := indicador.Variables
					if len(variables) > 0 {

						println("Indicador: ")
						println(">>>>>> ", indicador.Nombre)
						println("Variables: ")
						for _, variable := range variables {
							println("<<<<<< ", variable.Nombre)
						}
					} else {
                        println("Indicador: ")
                        println(">>>>>> ", indicador.Nombre)
                        println("Variables: ")
                        println("<<<<<< No hay variables en este indicador")
                    }

				}
				println("====================================")
			}
			println("====================================")
		}
	}
}
