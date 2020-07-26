package user

import "time"

type Usuario struct {
	ID        int
	Nombre    string
	FechaAlta time.Time
	Status    bool
}

func (user *Usuario) AltaUsuario(id int, nombre string, fechaAlta time.Time, status bool) {
	user.ID = id
	user.Nombre = nombre
	user.Status = status
	user.FechaAlta = fechaAlta
}
