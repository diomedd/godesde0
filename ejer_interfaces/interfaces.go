package ejer_interfaces

import (
	"fmt"

	"github.com/diomedd/godesde0/interfaces"
)

func HumanosRespirando(hu interfaces.Humano) {

	hu.Respirar()
	//hu.Pensar()
	fmt.Printf("Soy un/a %s, y estoy pensando \n", hu.Sexo())
}

//func EsunSerVivo(hu interfaces.Humano) {

//	hu.EstaVivo()
//hu.Pensar()
//	fmt.Printf("Esta vivo/a %b, y estoy pensando \n", hu.EstaVivo())
//}
