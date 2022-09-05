package main

import (
	"fmt"
	"log"
)

type Database interface {
	GetUser() string
}

type DefaultDatabase struct {}

func (db DefaultDatabase) GetUser() string {
	return "bob"
}

type Greater interface {
	Greet(userName string) 
}

type NiceGreeter struct {}

func (ng NiceGreeter) Greet(userName string) {
	log.Printf( "Hello %s", userName)
}

type Program struct {
	Db Database
	Greeter Greater
}

func (p Program) Execute() {
	user := p.Db.GetUser()
	p.Greeter.Greet(user)
}

func main() {
	fmt.Println("My programm is runs")
	db := DefaultDatabase{}
	greeter := NiceGreeter{}
	p := Program{
		Db: db,
		Greeter: greeter,
	}

	p.Execute()

}