package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func insert(original []int, position int, value int) []int {

	//Creamos un slice nuevo con 1 posición más que el arreglo original.
	newArray := make([]int, len(original)+1)

	//Copiamos todo lo que esta antes de la posición, es decir, todo lo que está a la izquierda de la posición.
	copy(newArray, original[:position])

	//Insertarmos el valor en la posición deseada usando un puntero
	ptr := &newArray[position]
	*ptr = value

	//Copiamos despues de la posicion deseada en el nuevo arreglo,
	//todo lo que esta desde la posicion hasta el final del arreglo original,
	//es decir, todo lo que está desde la posicion hacia la derecha.
	copy(newArray[position+1:], original[position:])

	return newArray
}

func main() {

	// Sustituir elemento en una posición definida.

	array := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Printf("Array original: ")
	for i := 0; i < len(array); i++ {
		fmt.Printf("%d ", array[i])
	}
	fmt.Println()

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Posicion a modificar: ")
	scanner.Scan()
	pos, _ := strconv.ParseInt(scanner.Text(), 10, 64)

	ptr := &array[pos]

	fmt.Printf("Numero nuevo: ")
	scanner.Scan()
	newNumber, _ := strconv.ParseInt(scanner.Text(), 10, 32)
	*ptr = int(newNumber)

	fmt.Printf("Array modificado: ")
	for i := 0; i < len(array); i++ {
		fmt.Printf("%d ", array[i])
	}
	fmt.Println()

	//--------------------------------------------------------------------
	// Agregar elemento en una posición definida y extender el array.

	//Para usar memoria dinamica se deben de usar slices,
	//ya que los arreglos se crean con memoria definida.

	array2 := array[0:] // Creamos un slice del arreglo original desde la posición 0,
	// incluyendola, hasta el final del array.

	array3 := insert(array2, 3, 199)

	fmt.Printf("Array con elemento agregado: ")
	for i := 0; i < len(array3); i++ {
		fmt.Printf("%d ", array3[i])
	}

	//Otros hallazgos:
	//La aritmetica de punteros en Golang se considera unsafe, por lo cual no es recomendable usarla.
	//Además, en caso de querer usarla, se debe de usar el import unsafe, junto con ciertas funciones,
	//a diferencia de C que solamente hacemos PUNTERO++.

}
