package sxengine

import (
	. "secretaria.admin/indicadores/sxengine/types"
)

func (i *AlmacenDeRegistros) AddRegistro(registro Registro) Registro {
	i.registros = append(i.registros, registro)
	for index, reg := range i.registros {
		if reg.IdRegistro == registro.IdRegistro {
			i.registros[index] = registro
		}
	}
	return registro
}

func (i *AlmacenDeRegistros) GetRegistros() []Registro {
	return i.registros
}
