package defer_panic

import (
	"fmt"
	"log"
)

func VemosDefer() {

	fmt.Println("este es un primer mensaje")

	defer fmt.Println("este es el mensaje final")

	fmt.Println("esta es el segundo mensaje")

}

func EjemploPanic() {
	defer func() {

		reco := recover()
		if reco != nil {
			log.Fatalf("ocurrio un error q genero panic \n %v", reco)

		}
	}()
	a := 1
	if a == 1 {
		panic("se encontro el valor 1")
	}
}
