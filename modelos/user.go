package modelos

import "time"

type User struct {
	Id      int
	Name    string
	Created time.Time
	Status  bool
}

func (us *User) AddUser(id int, name string, created time.Time, status bool) {
	us.Id = id
	us.Name = name
	us.Created = created
	us.Status = status
}
