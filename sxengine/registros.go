package sxengine

import (
	"errors"
	. "secretaria.admin/indicadores/sxengine/types"
	"time"
)

func isSameDayMonthYear(date1 time.Time, date2 time.Time) bool {
	return date1.Day() == date2.Day() && date1.Month() == date2.Month() && date1.Year() == date2.Year()
}
func (i *AlmacenDeRegistros) AddRegistro(registro Registro) (Registro, error) {

	isFirstRegistro := len(i.registros) == 0

	idDepartamento := registro.IdDepartamento
	idIndicador := registro.IdIndicador

	manager := GestorDeIndicadores{}

	indicador, err := manager.GestionDeIndicadores.GetIndicadorById(idIndicador)

	if err != nil {
		return Registro{}, err
	}

	dimension, err := manager.GestionDeIndicadores.GetDimensionById(indicador.IdDimension)

	if err != nil {
		return Registro{}, err
	}

	estaRelacion := DepartamentoIndicador{}
	for _, relacion := range manager.RelacionDepartamentoIndicador {
		if relacion.IdDepartamento == idDepartamento && relacion.IdIndicador == idIndicador {
			estaRelacion = relacion
			break
		}
	}
	if estaRelacion.IdDepartamento == 0 {
		return Registro{}, errors.New("No se encontro la relacion entre el departamento y el indicador")

	}
	isFechaRegistroIgualaFechaInicio := isSameDayMonthYear(registro.Fecha, estaRelacion.FechaInicio)
	if isFirstRegistro && isFechaRegistroIgualaFechaInicio {
		registro.IsOpen = false
		i.registros = append(i.registros, registro)
		for index, reg := range i.registros {
			if reg.IdRegistro == registro.IdRegistro {
				i.registros[index] = registro
			}
		}

		// crear nuevo registro vacio para el siguiente registro
		espacioRegistro := Registro{
			IsOpen: true,
		}
		espacioRegistro.IdRegistro = registro.IdRegistro + 1
		espacioRegistro.Fecha = obtenerSiguienteFecha(registro.Fecha, dimension.Frecuencia)
		espacioRegistro.Valor = 0
		i.registros = append(i.registros, espacioRegistro)

		return registro, nil
	}
	// si no es el primer registro buscar si ya hay un registro con espacio para el siguiente
	registrosAbieros := []Registro{}
	for _, reg := range i.registros {
		if reg.IsOpen {
			registrosAbieros = append(registrosAbieros, reg)
		}
	}
	if len(registrosAbieros) == 0 {
		return Registro{}, errors.New("No hay espacio para el siguiente registro")
	}
	for _, reg := range registrosAbieros {
		if reg.IdRegistro == registro.IdRegistro && reg.Valor == 0 {
			// comprobar si la fecha de este registro concuerda con el espacio disponible
			if isSameDayMonthYear(reg.Fecha, registro.Fecha) {
				registro.IsOpen = false
				i.registros = append(i.registros, registro)
				for index, reg := range i.registros {
					if reg.IdRegistro == registro.IdRegistro {
						i.registros[index] = registro
					}
				}

				// crear nuevo registro vacio para el siguiente registro
				espacioRegistro := Registro{
					IsOpen: true,
				}
				espacioRegistro.IdRegistro = registro.IdRegistro + 1
				espacioRegistro.Fecha = obtenerSiguienteFecha(registro.Fecha, dimension.Frecuencia)
				espacioRegistro.Valor = 0
				i.registros = append(i.registros, espacioRegistro)

				return registro, nil
			}
			return Registro{}, errors.New("No estas en la fecha correcta para este registro" + reg.Fecha.String())
		}
		return Registro{}, errors.New("No estas en la fecha correcta para este registro" + reg.Fecha.String())
	}
	return registro, nil
}
func obtenerSiguienteFecha(fechaRegistro time.Time, frecuencia string) time.Time {
	switch frecuencia {
	case SEMANAL.String():
		return fechaRegistro.AddDate(0, 0, 7)
	case MENSUAL.String():
		return fechaRegistro.AddDate(0, 1, 0)
	case TRIMESTRAL.String():
		return fechaRegistro.AddDate(0, 3, 0)
	case SEMESTRAL.String():
		return fechaRegistro.AddDate(0, 6, 0)
	case ANUAL.String():
		return fechaRegistro.AddDate(1, 0, 0)
	}
	return fechaRegistro
}
func (i *AlmacenDeRegistros) GetRegistros() []Registro {
	return i.registros
}
