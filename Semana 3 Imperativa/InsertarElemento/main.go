package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

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

	//new_number := &array[pos]
	ptr := &array[pos]

	fmt.Printf("Numero nuevo: ")
	scanner.Scan()
	new_number, _ := strconv.ParseInt(scanner.Text(), 10, 64)
	*ptr = int(new_number)

	fmt.Printf("Array modificado: ")
	for i := 0; i < len(array); i++ {
		fmt.Printf("%d ", array[i])
	}

}
