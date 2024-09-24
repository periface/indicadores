package sxengine

type Dependencia struct {
	IdDependencia int    `json:"IdDependencia"`
	Nombre        string `json:"Nombre"`
	Areas         []Area
}

