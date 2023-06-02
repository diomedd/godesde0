package middleware

import "fmt"

func sumar(a, b int) int {

	return a + b
}

func restar(a, b int) int {

	return a - b
}

func multiplicar(a, b int) int {

	return a * b
}

func MiMiddleware() {

	fmt.Println("Inicio")

	result := OperacionesMid(sumar)(5, 3)
	fmt.Println(result)

	result = OperacionesMid(restar)(10, 3)
	fmt.Println(result)

	result = OperacionesMid(multiplicar)(10, 5)
	fmt.Println(result)

}

func OperacionesMid(f func(int, int) int) func(int, int) int {

	return func(a, b int) int {

		fmt.Println("Inicio de operacion")
		return f(a, b)
	}
}
