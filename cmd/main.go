package main

import (
	"encoding/json"

	"secretaria.admin/indicadores/models"
)

func printAsJson(i interface{}) {
	json, err := json.Marshal(i)
	if err != nil {
		println(err)
	}
	println(string(json))
}
func main() {
	manager := models.GestorDeIndicadores{}
	//Creamos la dependencia de la secretaria de administracion
	print(manager)
}
