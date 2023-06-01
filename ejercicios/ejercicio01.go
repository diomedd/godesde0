package ejerciciosi

import "strconv"

func ConvaNumerico(texto string) (int, string) {

	//num, _ := strconv.Atoi(texto) //con el guion _ no asigno el argumento a ninguna variable

	num, err := strconv.Atoi(texto)
	if err != nil {
		return 0, "hubo un error " + err.Error()
	}
	if num > 100 {

		return num, "es mayor a 100 "
	} else {

		return num, "es menor 100 "
	}
}
