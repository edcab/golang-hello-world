package main

import (
	"fmt"
	"time"

	us "./user"
)

type pepe struct {
	us.Usuario
}

func main() {

	user := new(pepe)
	user.AltaUsuario(100, "Edwin", time.Now(), true)
	fmt.Println(user)
}
