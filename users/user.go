package users

import (
	"fmt"
	"time"

	"github.com/marcegabal/godesde0/modelos"
)

func AltaUsuario() {
	u := new(modelos.User)
	u.AddUser(1, "Pepe", time.Now() ,true)
	fmt.Println(u)
}