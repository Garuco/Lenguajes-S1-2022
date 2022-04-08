package main

import "fmt"

type Shoe struct {
	brand string
	price int
	size  int
}

func getSizeInRange(array []Shoe, min int, max int) (newArray []Shoe) {
	for i := 0; i < len(array); i++ {
		if array[i].size >= min && array[i].size <= max {
			newArray = append(newArray, array[i])
		}
	}
	return newArray
}

func main() {

	z1 := Shoe{"Nike Metcon 7", 102000, 40}
	z2 := Shoe{"Reebok Nano X1", 120000, 42}
	z3 := Shoe{"Nike AF1", 90000, 41}
	z4 := Shoe{"Nobull Trainer+", 115000, 34}
	z5 := Shoe{"Reebok Club C 85 Vintage", 85000, 36}
	z6 := Shoe{"Reebok Club C Revenge Vintage", 80000, 44}
	z7 := Shoe{"Nike Air Jordan 1 Mid Paris", 150000, 35}
	z8 := Shoe{"Nike React Gato", 75000, 37}
	z9 := Shoe{"Nike React Tiempo", 82000, 39}
	z10 := Shoe{"Nike Mercurial Superfly 8", 98000, 40}

	arrayShoes := []Shoe{z1, z2, z3, z4, z5, z6, z7, z8, z9, z10}

	arrayShoesSize := getSizeInRange(arrayShoes, 37, 42)

	fmt.Println(arrayShoesSize)

	// La estrategia utilizada para crear un array dinamico en ejecucion fue el uso de slices, ya que estos
	// son dinamicos.

	//--------------------------------------------------------------------------------------------
	//Si se quiere leer desde consola la informacion descomentar este codigo y comentar el de arriba.

	/*arrayShoe := []Shoe{}

	scanner := bufio.NewScanner(os.Stdin)

	for i := 0; i < 10; i++ {
		fmt.Printf("Digite la marca y modelo del zapato: ")
		scanner.Scan()
		brand := scanner.Text()

		fmt.Printf("Digite el precio: ")
		scanner.Scan()
		price, _ := strconv.ParseInt(scanner.Text(), 10, 64)

		fmt.Printf("La talla: ")
		scanner.Scan()
		size, _ := strconv.ParseInt(scanner.Text(), 10, 64)

		arrayShoe[i] = Shoe{brand, int(price), int(size)}
		fmt.Println()

	}

	arrayShoesSize := getSizeInRange(arrayShoe, 37, 42)

	fmt.Println(arrayShoesSize)*/

}
