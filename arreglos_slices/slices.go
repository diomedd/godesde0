package arreglos_slices

import "fmt"

var tablaS []int = []int{2, 5, 4}
var arreglo [10]int = [10]int{6, 333, 3454, 657, 787, 11, 484, 222}

func MuestroSlices() {
	fmt.Println(tablaS)

	porcion := arreglo[3:]   //slice creado desde un vector desde la pocision 3
	porcion2 := arreglo[:5]  //desde la primera hasta la 5
	porcion3 := arreglo[6:7] //conjunto de elementos

	fmt.Println("porcion 1 ", porcion)
	fmt.Println("porcion 2 ", porcion2)
	fmt.Println("porcion 3 ", porcion3)
}

func Capacidad() {

	elementos := make([]int, 5, 20)
	fmt.Printf("largo %d, capacidad %d", len(elementos), cap((elementos)))
	nums := make([]int, 0, 0)

	for i := 0; i < 100; i++ {

		nums = append(nums, i)

	}
	fmt.Printf("\n largo %d, capacidad %d", len(nums), cap((nums)))
}
