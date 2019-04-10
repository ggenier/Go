package main

import (
	"fmt"
	"time"
)

func main() {

	//Infinite loop (false infinite)
	fmt.Println("Infinite loop (false infinite)")
	i := 10
	for {
		if i == 0 {
			fmt.Println("End of the false ifinite loop")
			break
		}
		i--
	}

	//for loop without break
	fmt.Println("\nfor loop without break")
	for timer := 5; timer >= 0; timer-- {
		if timer == 0 {
			fmt.Println("Boom !")
		}
		fmt.Println(timer)
		time.Sleep(1 * time.Second)
	}

	//for loop with break
	fmt.Println("\nfor loop with break")
	for timer := 5; timer >= 0; timer-- {
		if timer == 0 {
			fmt.Println("Boom !")
			break //end of the loop
		}
		fmt.Println(timer)
		time.Sleep(1 * time.Second)
	}

	//for loop with continue
	fmt.Println("\nfor loop with continue")
	for timer := 5; timer >= 0; timer-- {
		if timer%2 == 0 {
			continue
		}
		fmt.Println(timer)
		time.Sleep(1 * time.Second)
	}

	//For loop with range
	fmt.Println("\nfor loop with range")
	courses := []string{"Book 1", "Book 2", "Book 3"}
	for index, value := range courses{
		fmt.Println(index, " : ", value)
	}
}
