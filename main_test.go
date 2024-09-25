package main

import (
	"secretaria.admin/indicadores/sxengine"
	. "secretaria.admin/indicadores/sxengine/types"
	"testing"
)

var gestor = sxengine.GestorDeIndicadores{}

func TestGestorDepartamental(t *testing.T) {
	ConstruirDependencias(t, &gestor)
}
func TestGestorIndicadores(t *testing.T) {
	ConstruirDimensiones(t, &gestor)
}
func TestGestorRelaciones(t *testing.T) {
	ContruirRelaciones(t, &gestor)
}
func TestExpressions(t *testing.T) {
	RunExpressions(t, &gestor)
	gestor.ShowReport()
}
func RunExpressions(t *testing.T, gestor *sxengine.GestorDeIndicadores) {
	println("Running expressions")
}
func ContruirRelaciones(t *testing.T, gestor *sxengine.GestorDeIndicadores) {
	departamentoCompras, _ := gestor.GestionDepartamental.GetDepartamento(1)
	indicadorExpedienteCompras, _ := gestor.GestionDeIndicadores.GetIndicadorById(1)
	indicadorTiempoPromedio, _ := gestor.GestionDeIndicadores.GetIndicadorById(2)
	indicadorContratosPymes, _ := gestor.GestionDeIndicadores.GetIndicadorById(3)

	gestor.AsignarIndicador(departamentoCompras.IdDepartamento, indicadorExpedienteCompras.IdIndicador)
	gestor.AsignarIndicador(departamentoCompras.IdDepartamento, indicadorTiempoPromedio.IdIndicador)
	gestor.AsignarIndicador(departamentoCompras.IdDepartamento, indicadorContratosPymes.IdIndicador)

	if len(gestor.RelacionDepartamentoIndicador) != 3 {
		t.Errorf("Expected 1 relacion, got %d", len(gestor.RelacionDepartamentoIndicador))
	}
}
func ConstruirDimensiones(t *testing.T, gestor *sxengine.GestorDeIndicadores) {

	if len(gestor.GestionDepartamental.GetAreas()) != 2 {
		t.Errorf("Expected 2 areas, got %d", len(gestor.GestionDepartamental.GetAreas()))
	}
	gestionDeIndicadores := &gestor.GestionDeIndicadores
	gestionDeIndicadores.AddDimension(1, "Eficiencia", sxengine.SEMANAL, "Incremento")
	gestionDeIndicadores.AddDimension(2, "Calidad", sxengine.MENSUAL, "Decremento")
	gestionDeIndicadores.AddDimension(3, "Economia", sxengine.TRIMESTRAL, "Constante")

	if len(gestor.GestionDeIndicadores.GetDimensiones()) != 3 {
		t.Errorf("Expected 3 dimensiones, got %d", len(gestor.GestionDeIndicadores.GetDimensiones()))
	}

	eficiencia, _ := gestionDeIndicadores.GetDimensionById(1)
	calidad, _ := gestionDeIndicadores.GetDimensionById(2)
	economia, _ := gestionDeIndicadores.GetDimensionById(3)

	gestionDeIndicadores.AddIndicador(&eficiencia, Indicador{
		IdIndicador:     1,
		Nombre:          "Integración de Expedientes de Compras",
		Codigo:          "IEC",
		UnidadDeMedida:  "Porcentaje",
		Descripcion:     "Mide el porcentaje de expedientes de compras devueltos por la DGCyOP a la Dirección Administrativa por motivo de observaciones",
		MetodoDeCalculo: "(var1 / var2) * 100",
	})

	gestionDeIndicadores.AddIndicador(&calidad, Indicador{
		IdIndicador:     2,
		Nombre:          "Tiempo promedio de atención a observaciones de Expedientes de Compras ante la DGCyOP",
		Codigo:          "TPO",
		UnidadDeMedida:  "Días",
		Descripcion:     "Mide el tiempo promedio que le toma a la dirección administrativa atender observaciones hechas por la DGCyOP a sus expedientes de compras.",
		MetodoDeCalculo: "var1 / var2",
	})
	gestionDeIndicadores.AddIndicador(&economia, Indicador{
		IdIndicador:     3,
		Nombre:          "Contratos con MiPyMES",
		Codigo:          "PME",
		UnidadDeMedida:  "Porcentaje",
		Descripcion:     "Mide la proporción de contratos llevados a cabo con Pequeñas y Medianas Empresas.",
		MetodoDeCalculo: "(var1 / var2) * 100",
	})

	if len(gestionDeIndicadores.GetIndicadores()) != 3 {
		t.Errorf("Expected 3 indicadores, got %d", len(gestionDeIndicadores.GetIndicadores()))
	}

	indicadorExpedienteCompras, _ := gestionDeIndicadores.GetIndicadorById(1)

	if indicadorExpedienteCompras.Nombre != "Integración de Expedientes de Compras" {
		t.Errorf("Expected 'Integración de Expedientes de Compras', got %s", indicadorExpedienteCompras.Nombre)
	}

	gestionDeIndicadores.AddVariable(&indicadorExpedienteCompras, Variable{
		IdVariable: 1,
		Nombre:     "Expedientes devueltos",
		Codigo:     "var1",
	})
	gestionDeIndicadores.AddVariable(&indicadorExpedienteCompras, Variable{
		IdVariable: 2,
		Nombre:     "Expedientes totales",
		Codigo:     "var2",
	})

	if len(indicadorExpedienteCompras.Variables) != 2 {
		t.Errorf("Expected 2 variables, got %d", len(indicadorExpedienteCompras.Variables))
	}

	indicadorTiempoPromedio, _ := gestionDeIndicadores.GetIndicadorById(2)

	gestionDeIndicadores.AddVariable(&indicadorTiempoPromedio, Variable{
		IdVariable: 3,
		Nombre:     "Dias de atencion",
		Codigo:     "var1",
	})
	gestionDeIndicadores.AddVariable(&indicadorTiempoPromedio, Variable{
		IdVariable: 4,
		Nombre:     "Expedientes atendidos",
		Codigo:     "var2",
	})

    if len(indicadorTiempoPromedio.Variables) != 2 {
        t.Errorf("Expected 2 variables, got %d", len(indicadorTiempoPromedio.Variables))
    }

}
func ConstruirDependencias(t *testing.T, gestor *sxengine.GestorDeIndicadores) {

	gestionDepartamental := &gestor.GestionDepartamental
	gestionDepartamental.AddDependencia(Dependencia{IdDependencia: 1, Nombre: "Secretaria de Administracion"})

	if len(gestionDepartamental.GetDependencias()) != 1 {
		t.Errorf("Expected 1 dependencias, got %d", len(gestionDepartamental.GetDependencias()))
	}
	secretariaDeAdministracion, _ := gestionDepartamental.GetDependencia(1)
	gestionDepartamental.AddArea(&secretariaDeAdministracion, Area{IdArea: 1, Nombre: "Direccion General de Compras"})
	gestionDepartamental.AddArea(&secretariaDeAdministracion, Area{IdArea: 2, Nombre: "Direccion General de Contratos"})

	if len(gestionDepartamental.GetAreas()) != 2 {
		t.Errorf("Expected 2 areas, got %d", len(gestionDepartamental.GetAreas()))
	}

	direccionDeCompras, _ := gestionDepartamental.GetArea(1)
	if direccionDeCompras.Nombre != "Direccion General de Compras" {
		t.Errorf("Expected 'Direccion General de Compras', got %s", direccionDeCompras.Nombre)
	}

	direccionDeContratos, _ := gestionDepartamental.GetArea(2)
	if direccionDeContratos.Nombre != "Direccion General de Contratos" {
		t.Errorf("Expected 'Direccion General de Contratos', got %s", direccionDeContratos.Nombre)
	}

	gestionDepartamental.AddDepartamento(&direccionDeCompras, Departamento{IdDepartamento: 1, Nombre: "Departamento de Compras"})
	gestionDepartamental.AddDepartamento(&direccionDeCompras, Departamento{IdDepartamento: 2, Nombre: "Departamento de Licitaciones"})
	gestionDepartamental.AddDepartamento(&direccionDeContratos, Departamento{IdDepartamento: 3, Nombre: "Departamento de Contratos"})
	gestionDepartamental.AddDepartamento(&direccionDeContratos, Departamento{IdDepartamento: 4, Nombre: "Departamento de Adjudicaciones"})

	if len(direccionDeCompras.Departamentos) != 2 {
		t.Errorf("Expected 2 departamentos, got %d", len(direccionDeCompras.Departamentos))
	}
	if len(direccionDeContratos.Departamentos) != 2 {
		t.Errorf("Expected 2 departamentos, got %d", len(direccionDeContratos.Departamentos))
	}
}
