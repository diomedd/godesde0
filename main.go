package main

import (
	//ejerinterfaces "github.com/diomedd/godesde0/ejer_interfaces"
	//"github.com/diomedd/godesde0/modelos"

	//dp "github.com/diomedd/godesde0/defer_panic"
	//"github.com/diomedd/godesde0/files"
	//"github.com/diomedd/godesde0/iteracciones"
	//"fmt"

	//"github.com/diomedd/godesde0/gorutines"
	"github.com/diomedd/godesde0/webserver"
)

func main() {

	//fmt.Println("Hola Mundo")
	//variables.MuestroEnteros()
	//variables.RestoVariables()
	/*estado, texto := variables.ConviertoaTexto(1588)

	fmt.Println(estado)
	fmt.Println(texto)

	if os := runtime.GOOS; os == "linux" || os == "OS X." {

		fmt.Println("esto no es windows, es ", os)

	} else {

		fmt.Println("esto es windows", os)

	}

	switch os := runtime.GOOS; os {
	case "linux":
		fmt.Println("esto es linux")
	case "darwin":
		fmt.Println("esto es darwin")
	default:
		fmt.Printf("%s \n", os)



	numero, texto := ejerciciosi.ConvaNumerico("500")
	fmt.Println(numero)
	fmt.Println(texto)	}*/

	//teclado.IngresoNumero()

	//iteracciones.Iterar()

	//fmt.Println(ejercicios.TabladeMultiplicar())
	//files.GrabaTabla()
	//files.SumaTabla()
	//files.LeoArchivo()

	//funciones.Calculos()
	//funciones.LlamarClouser()
	//funciones.Exponencia(10)
	//	arreglos_slices.Capacidad()
	//mapas.MostrarMapas()
	//users.AltaUsuario()

	//Pedro := new(modelos.Hombre)
	//ejerinterfaces.HumanosRespirando(Pedro)

	//dp.VemosDefer()
	//dp.EjemploPanic()

	/*canal1 := make(chan bool)
	go gorutines.MiNombreLento("diego diomede", canal1)
	defer func() {
		<-canal1
	}()
	fmt.Println("estoy aqui")
	//_ = <-canal1

	//var x string
	//fmt.Scanln(&x)*/

	webserver.MiwebServer()

}
