package types

type Departamento struct {
	IdDepartamento int    `json:"IdDepartamento"`
	Nombre         string `json:"Nombre"`
	IdArea         int    `json:"IdArea"`
	Area           Area   `json:"Area"`
}
