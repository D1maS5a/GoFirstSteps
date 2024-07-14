package main

var _ User = &superUser{}

type superUser struct {
	Name      string
	Age       int
	IsBlocked bool
}

// ChangeAddress implements User.
func (s *superUser) Block() {
	s.IsBlocked = true
}

var _ User = &user{}

type user struct {
	FIO, Address, Phone string
	IsBlocked           bool
}

// ChangeAddress implements User.
func (u *user) ChangeAddress(newAddress string) {
	u.Address = newAddress
}

// ChangeFIO implements User.
func (u *user) Block() {
	u.IsBlocked = true
}

type User interface {
	Block()
}

func NewUserMy(fio, addresss, phone string) User {
	// u := user{fio, addresss, phone}
	u := user{}
	return &u
}

// func main(){

// }
