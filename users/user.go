package users

import (
	"fmt"
	"time"

	"github.com/diomedd/godesde0/modelos"
)

func AltaUsuario() {

	u := new(modelos.User)
	u.AddUser(10, "pablo", time.Now(), true)
	fmt.Println(u)
}
