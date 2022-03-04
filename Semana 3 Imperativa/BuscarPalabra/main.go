package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var frase string
var palabra string
var validar int

func buscarPalabra() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Escriba una frase: ")
	scanner.Scan()
	frase = scanner.Text()

	fmt.Printf("Escriba una palabra: ")
	scanner.Scan()
	palabra = scanner.Text()
	i := 0
	validar = 1
	for i < len(frase) {
		if strings.ToLower(string(frase[i])) == strings.ToLower(string(palabra[0])) {
			i++
			for validar < len(palabra) {
				if strings.ToLower(string(frase[i])) == strings.ToLower(string(palabra[validar])) {
					validar++
					i++
				} else {
					validar = 0
					break
				}
			}
			if validar == len(palabra) {
				fmt.Println("La palabra se encuentra en la frase.")
				break
			}
		}
		i++
	}
	if validar != len(palabra) {
		fmt.Println("La palabra no se encuentra en la frase.")
	}

	//fmt.Println(string(frase[0]))
	//fmt.Println(palabra)
}

func main() {

	buscarPalabra()

}
