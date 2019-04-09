package main

import (
	"fmt"
	"strings"
)

func main(){
	title := "DEVELOP IN GO"
	author := "gregoire genier"

	fmt.Println(converter(title, author))
	fmt.Println("Multiply 3*3 : ", multiply(3, 3))

	fmt.Println("High level : ", calculateHighLevel(8, 10, 5, 45, 10, 60))

}

func converter(title, author string) (s1, s2 string) {
	title = strings.ToLower(title)
	author = strings.ToTitle(author)

	return title, author
}

func multiply(number1, number2 int) (result int){
	return number1 * number2
}

func calculateHighLevel(levels ...int) (result int){
	result = levels[0]
	for _, f_i := range(levels){
		if f_i > result{
			result = f_i
		}
	}

	return result
}
