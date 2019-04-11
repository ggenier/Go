package main

import (
	"fmt"
)

type (
	globalTypePerson struct{
		name string
		firstName string
		age int
	}
)
func main(){
	type localTypePerson struct {
		name string
		firstName string
		age int
	}

	me := localTypePerson{"GG", "Gregoire", 38}
	fmt.Println(me)

	fmt.Printf("My name is %v, my first name is %v and I'my %v old\n", me.name, me.firstName, me.age)

	//Age modification
	me.age = 39
	fmt.Printf("After one year, I'my %v old\n", me.age)

	me2 := new(globalTypePerson)
	me2.name = "TEST"
	me2.firstName = "TEST2"
	me2.age = 60
	fmt.Printf("(2)My name is %v, my first name is %v and I'my %v old\n", me2.name, me2.firstName, me2.age)

}
