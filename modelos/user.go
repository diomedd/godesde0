package modelos

import "time"

type User struct {
	Id        int
	Name      string
	CreatedAT time.Time
	Status    bool
}

func (usuario *User) AddUser(id int, name string, createdAT time.Time, status bool) {

	usuario.Id = id
	usuario.Name = name
	usuario.CreatedAT = createdAT
	usuario.Status = status
}
