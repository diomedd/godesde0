package mapas

import "fmt"

func MostrarMapas() {

	paises := make(map[string]string, 5)

	fmt.Println(paises)
	paises["Argentina"] = "BA"
	paises["Mexico"] = "DF"

	//fmt.Println(paises)
	//fmt.Println(paises["Argentina"])

	campeonato := map[string]int{

		"barcelona":   39,
		"real madrid": 38,
		"river plate": 30,
		"velez":       25,
	}

	fmt.Println(campeonato)

	for equipo, puntaje := range campeonato {

		fmt.Printf("Equipo %s, tiene un puntaje de %d \n", equipo, puntaje)
	}

	delete(campeonato, "real madrid")
	fmt.Println(campeonato)

	puntaje, existe := campeonato["river plate"]
	fmt.Printf("el puntaje caputrado es %d, y el euqipo es %t \n", puntaje, existe)
}
