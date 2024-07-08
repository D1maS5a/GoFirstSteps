package main

import "fmt"

type Whisperer interface {
	Whisper() string
}

type Yeller interface {
	Yell() string
}

type Talker interface {
	Whisperer
	Yeller
}

type Account struct {
	accountId int
	balance   int
	name      string
}

type ManagerAccount struct {
	Account
}

func (a *Account) GetBalance() int {
	return a.balance
}

func (a *Account) String() string {
	return fmt.Sprintf("Standart (%v) $%v \"%v\"", a.accountId, a.balance, a.name)
}

func (m *ManagerAccount) String() string {
	return fmt.Sprintf("Manager (%v) $%v \"%v\"", m.accountId, m.balance, m.name)
}

func main() {
	mgrAccount := ManagerAccount{Account{2, 30, "Cassandra"}}
	fmt.Println(mgrAccount)

	mgrAcc := ManagerAccount{Account{2, 30, "Cassandra"}}
	fmt.Printf("%v \n", mgrAcc)
	fmt.Printf("%v \n", mgrAcc.GetBalance())
	fmt.Printf("%v \n", mgrAcc.accountId)
	fmt.Println(mgrAccount.String())
}
