package main

import (
	"fmt"
)

func main(){
	//Map creation with make
	fmt.Println("Create a map with make")
	myMap := make(map[string]string)
	for key, value := range myMap {
		fmt.Printf("Key %v / value %v", key, value)
	}

	//Add a value to the map
	fmt.Println("\nAdd a value to the map")
	myMap["GG"] = "gregoire"
	for key, value := range myMap {
		fmt.Printf("\tKey %v / value %v", key, value)
	}
	fmt.Println("")

	//Map creation without make
	fmt.Println("Create a map without make")
	myMap2 := map[string]string{"JSE" : "Jeremie", "MF" : "Mathilde"}
	for key, value := range myMap2 {
		fmt.Printf("\tKey %v / value %v\n", key, value)
	}
	fmt.Println("")

	//Modify a value
	fmt.Println("Modify a value gregoire => gregory")
	myMap["GG"] = "gregory"
	for key, value := range myMap {
		fmt.Printf("\tKey %v / value %v\n", key, value)
	}
	fmt.Println("")

	//Delete a value
	fmt.Println("Delete myMap2[JSE]")
	delete(myMap2, "JSE")
	for key, value := range myMap2 {
		fmt.Printf("\tKey %v / value %v\n", key, value)
	}
	fmt.Println("")

	//Map in function parameter
	fmt.Println("\nCall function with myMap in parameter")
	modifyMap(myMap2, "GG", "gregoire")
	fmt.Println("\nAfter return of function with myMap in parameter")
	for key, value := range myMap2 {
		fmt.Printf("\tKey %v / value %v\n", key, value)
	}
	fmt.Println("")

}

func modifyMap(myMap map[string]string, key string, value string){
	fmt.Println("\n\tMap before modification")
	for key, value := range myMap {
		fmt.Printf("\tKey %v / value %v\n", key, value)
	}

	//Append or modify value
	myMap[key] = value
	fmt.Println("\n\tMap after modification")
	for key, value := range myMap {
		fmt.Printf("\t\tKey %v / value %v\n", key, value)
	}
}
