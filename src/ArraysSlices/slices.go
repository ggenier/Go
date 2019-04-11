package main

import (
	"fmt"
)

func main() {
	fmt.Println("Empty Slice with make instruction")
	firstSlice := make([]string, 4, 10)
	fmt.Println("   Slice length : ", len(firstSlice), " capacity ", cap(firstSlice))
	for index, value := range firstSlice {
		fmt.Println("   index : ", index, " value : ", value)
	}

	fmt.Println("\nNot empty Slice : 3 books with [] instruction")
	secondSlice := []string{"Booke 1", "Book 2", "Book3"}
	fmt.Println("   Slice length : ", len(secondSlice), " capacity ", cap(secondSlice))
	for index, value := range secondSlice {
		fmt.Println("   index : ", index, " value : ", value)
	}
	fmt.Println("   Append element to secondSlice")
	secondSlice = append(secondSlice, "Book 4")
	fmt.Println("      Slice length : ", len(secondSlice), " capacity ", cap(secondSlice))
	for index, value := range secondSlice {
		fmt.Println("      index : ", index, " value : ", value)
	}

	fmt.Println("\nNot empty Slice : 2 books")
	thirdSlice := []string{"Book 5", "Book 6"}
	fmt.Println("   Slice length : ", len(thirdSlice), " capacity ", cap(thirdSlice))
	for index, value := range thirdSlice {
		fmt.Println("   index : ", index, " value : ", value)
	}
	fmt.Println("\nAppend secondSlice to thirdSlice")
	thirdSlice = append(thirdSlice, secondSlice...)
	fmt.Println("   Slice length : ", len(thirdSlice), " capacity ", cap(thirdSlice))
	for index, value := range thirdSlice {
		fmt.Println("   index : ", index, " value : ", value)
	}

	//Renference test, no change in thirdSlice
	fmt.Println("\nChange value in secondSlice[0] = Book 9")
	secondSlice[0] = "Book 9"
	fmt.Println("\nSecond slice : ")
	fmt.Println("   Slice length : ", len(secondSlice), " capacity ", cap(secondSlice))
	for index, value := range secondSlice {
		fmt.Println("      index : ", index, " value : ", value)
	}
	fmt.Println("\nThird slice : ")
	fmt.Println("   Slice length : ", len(thirdSlice), " capacity ", cap(thirdSlice))
	for index, value := range thirdSlice {
		fmt.Println("   index : ", index, " value : ", value)
	}

	fmt.Println("\nCall function with thirdSlice in parameter")
	modifySlice(thirdSlice)
	fmt.Println("\nThird slice after function call : ")
	fmt.Println("   Slice length : ", len(thirdSlice), " capacity ", cap(thirdSlice))
	for index, value := range thirdSlice {
		fmt.Println("   index : ", index, " value : ", value)
	}
}

func modifySlice(slice []string){
	fmt.Println("\nSlice receive in parameter")
	fmt.Println("   Slice length : ", len(slice), " capacity ", cap(slice))
	for index, value := range slice {
		fmt.Println("   index : ", index, " value : ", value)
	}

	//Modify value
	slice[0] = "Book 10"
	fmt.Println("Slice modified 0=>book 10")
	fmt.Println("   Slice length : ", len(slice), " capacity ", cap(slice))
	for index, value := range slice {
		fmt.Println("   index : ", index, " value : ", value)
	}
}
