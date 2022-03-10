package main

import (
	"fmt"
	"strings"
)

//Funcion creada por mi
func buscarSubcadena(subString string, originalString string) bool {
	i := 0
	validar := 1
	for i < len(originalString) {
		if strings.ToLower(string(originalString[i])) == strings.ToLower(string(subString[0])) {
			i++
			for validar < len(subString) {
				if strings.ToLower(string(originalString[i])) == strings.ToLower(string(subString[validar])) {
					validar++
					i++
				} else {
					validar = 0
					break
				}
			}
			if validar == len(subString) {
				return true
			}
		}
		i++
	}
	return false
}

func main() {

	fmt.Println(
		buscarSubcadena("Programando en Go", "Estoy programando en go, y es divertido"))

	fmt.Println(
		buscarSubcadena("Voy manejando", "Estoy programando en go, y es divertido"))

	//Funciones ya predefinidas/existentes.
	fmt.Println(strings.Contains("Ayer Julián se cayó de la bici", "Julián se cayó"))

	fmt.Println(strings.Contains("Ayer Julián se cayó de la bici", "No hay verduras"))

}
