package files

import (
	"fmt"
	//"io/ioutil"
	"bufio"
	"os"

	"github.com/diomedd/godesde0/ejercicios"
)

var filename string = "./files/txt/tabla.txt"

func GrabaTabla() {

	var texto string = ejercicios.TabladeMultiplicar()
	archivo, err := os.Create("./files/txt/tabla.txt")

	if err != nil {

		fmt.Println("error al mostrar el archivo " + err.Error())
		return
	}
	fmt.Fprintln(archivo, texto)
	archivo.Close()
}

func SumaTabla() {

	var texto string = ejercicios.TabladeMultiplicar()

	//if Append(filename, texto) == false { esto es una sugerencia de go
	if !Append(filename, texto) {
		fmt.Println("error al concatener contenido")
	}
}

func Append(filen string, texto string) bool {

	arch, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {

		fmt.Println("error durante el append " + err.Error())
		return false
	}
	_, err = arch.WriteString(texto)
	if err != nil {

		fmt.Println("error durante el writestring " + err.Error())
		return false
	}
	arch.Close()
	return true
}

/*func LeoArchivo() { esta deprecado io

	archivo, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("error al leer el archivo" + err.Error())
		return

	}

	fmt.Println(string(archivo))
}*/

func LeoArchivo() {

	archivo, err := os.Open(filename)
	if err != nil {

		fmt.Println("error leyendo archivo " + err.Error())
	}

	scanner := bufio.NewScanner(archivo)
	for scanner.Scan() {
		registro := scanner.Text()
		fmt.Println("> " + registro)
	}
	archivo.Close()
}
