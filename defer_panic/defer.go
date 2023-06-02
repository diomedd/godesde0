package defer_panic

import (
	"fmt"
)

func VemosDefer() {

	fmt.Println("este es un primer mensaje")
	defer fmt.Println("este es el mensaje final")

	fmt.Println("esta es el segundo mensaje")

}
