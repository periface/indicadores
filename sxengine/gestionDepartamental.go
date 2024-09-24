package sxengine

import "errors"

func (i *GestionDepartamental) CreateDependencia(idDependencia int, nombre string) {
	dependencia := Dependencia{
		IdDependencia: idDependencia,
		Nombre:        nombre,
		Areas:         []Area{},
	}
	i.dependencias = append(i.dependencias, dependencia)
}
func (i *GestionDepartamental) AddDependencia(dependencia Dependencia) Dependencia {
	i.dependencias = append(i.dependencias, dependencia)
	return dependencia
}

func (i *GestionDepartamental) AddArea(dependencia *Dependencia, area Area) Area {
	dependencia.Areas = append(dependencia.Areas, area)
	for index, dep := range i.dependencias {
		if dep.IdDependencia == dependencia.IdDependencia {
			i.dependencias[index] = *dependencia
		}
	}
	return area
}
func (i *GestionDepartamental) AddDepartamento(area *Area, departamento Departamento) Departamento {
	area.Departamentos = append(area.Departamentos, departamento)
	for index, dep := range i.dependencias {
		for indexArea, ar := range dep.Areas {
			if ar.IdArea == area.IdArea {
				i.dependencias[index].Areas[indexArea] = *area
			}
		}
	}
	return departamento
}

func (i *GestionDepartamental) GetAreas() []Area {
	Areas := []Area{}
	for _, dependencia := range i.dependencias {
		for _, area := range dependencia.Areas {
			Areas = append(Areas, area)
		}
	}
	return Areas
}
func (i *GestionDepartamental) GetArea(idArea int) (Area, error) {
	for _, dependencia := range i.dependencias {
		for _, area := range dependencia.Areas {
			if area.IdArea == idArea {
				return area, nil
			}
		}
	}
	return Area{}, errors.New("Area no encontrada")
}

func (i *GestionDepartamental) GetDependencias() []Dependencia {
	return i.dependencias
}

func (i *GestionDepartamental) GetDependencia(idDependencia int) (Dependencia, error) {
	for _, dependencia := range i.dependencias {
		if dependencia.IdDependencia == idDependencia {
			return dependencia, nil
		}
	}
	return Dependencia{}, errors.New("Dependencia no encontrada")
}

func (i *GestionDepartamental) GetDepartamentos() []Departamento {
	Departamentos := []Departamento{}
	for _, dependencia := range i.dependencias {
		for _, area := range dependencia.Areas {
			for _, departamento := range area.Departamentos {
				Departamentos = append(Departamentos, departamento)
			}
		}
	}
	return Departamentos
}
func (i *GestionDepartamental) GetDepartamento(idDepartamento int) (Departamento, error) {
	for _, dependencia := range i.dependencias {
		for _, area := range dependencia.Areas {
			for _, departamento := range area.Departamentos {
				if departamento.IdDepartamento == idDepartamento {
					return departamento, nil
				}
			}
		}
	}
	return Departamento{}, errors.New("Departamento no encontrado")
}
