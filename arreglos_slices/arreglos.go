package arreglos_slices

import "fmt"

var tabla [10]int = [10]int{1, 3, 4, 69}
var matriz [20][30]int

func MuestroArreglos() {
	tabla[7] = 33
	tabla[2] = 55
	tabla2 := [10]string{"diego", "gaston", "diomede"}
	fmt.Println(tabla)
	fmt.Println("vector de string", tabla2)

	for i := 0; i < len(tabla); i++ {

		fmt.Println(tabla[i])
	}

	matriz[7][24] = 15

	fmt.Println(matriz)
}
