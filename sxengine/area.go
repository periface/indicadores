package sxengine

type Area struct {
	IdArea int    `json:"IdArea"`
	Nombre string `json:"Nombre"`

	DependenciaId int         `json:"DependenciaId"`
	Dependencia   Dependencia `json:"Dependencia"`
	Departamentos []Departamento
}

