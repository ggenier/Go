package main

import (
	"fmt"
	"reflect"
	"os"
)

//Global package constant
const (
	aConstantPi = 3.14
)

//Global package variables
//initialize variables without default value => manual type
var (
	firstName string
	lastName  string
	weight 	 float64
)

//initialize variables with default value => autmatic type
var (
	sex, maried, height = "M", true, 170.0
)



func main(){
	//Locals variables
	var (
		phoneNumber string
	)

	path := os.Getenv("PATH")
	fmt.Println("PATH environement variable : ", path)
	pid := os.Getpid()
	fmt.Println("Program PID : ", pid)

	city := "tours"
	//adress := "somewhere" => variable must be used, if not compilation error
	phoneNumber = "0000000000"
	fmt.Println("Phone number : ", phoneNumber, "is type ", reflect.TypeOf(phoneNumber))

	fmt.Println("First name : ", firstName, "is type ", reflect.TypeOf(firstName))
	fmt.Println("Last name : ", lastName, "is type ", reflect.TypeOf(lastName))
	fmt.Println("Weight : ", weight, "is type ", reflect.TypeOf(weight))

	fmt.Println("Sex : ", sex, "is type ", reflect.TypeOf(sex))
	fmt.Println("Maried ? : ", maried, "is type ", reflect.TypeOf(maried))
	fmt.Println("Height : ", height, "is type ", reflect.TypeOf(height))

	fmt.Println("City : ", city, "is type ", reflect.TypeOf(city))

	fmt.Println("Memory adress for city variable  : ", &city)

	pointer := &city
	//& => give memory adresse, * => give value of a memory aadress
	fmt.Println("Memory adress for pointer variable  : ", pointer, " and value of ointer is ", *pointer)

	//Passsing by value
	fmt.Println("=== Passing variable by value")
	fmt.Println("City variable before changeCityByValue function : ", city)
	changeCityByValue(city)
	fmt.Println("City variable after changeCityByValue function : ", city)

	//Passsing by reference
	fmt.Println("=== Passing variable by reference")
	fmt.Println("City variable before changeCityByReference function : ", city)
	//We need to pass amemory adress of variable
	changeCityByReference(&city)
	fmt.Println("City variable after changeCityByReference function : ", city)

	//Function with return value
	fmt.Println("=== Function with return value")
	fmt.Println("Maried before function changeMaried : ", maried)
	maried = changeMaried(maried)
	fmt.Println("Maried after function changeMaried : ", maried)

	//It is not possible to change a constant
	fmt.Println("Pi value : ", aConstantPi)
	//aConstantPi = 2.5
}

func changeCityByValue(city string){
	fmt.Println("   In function changeCityByValue before changes : ", city)
	city="Paris"
	fmt.Println("   In function changeCityByValue after changes", city)
}

//For reference we need to put "*" for in every place where we use variable, if we want to 
//use the value
func changeCityByReference(city *string){
	fmt.Println("   In function changeCityByReference before changes : ", *city)
	*city="Paris"
	fmt.Println("   In function changeCityByReference after changes", *city)
}

func changeMaried(maried bool) bool{
	return !maried
}
