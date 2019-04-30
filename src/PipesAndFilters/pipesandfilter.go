package main

import(
	"fmt"
	"runtime"
)

func main(){
	runtime.GOMAXPROCS(4)
	
	//Create channel
	ch := make(chan int)
	
	//go routine, to generate number
	go generate(ch)

	//Loop to filter all number
	for{
		//Read a number
		prime := <- ch
		fmt.Println(prime)
		//Create a new channel to filter
		ch1 := make(chan int)
		//Filter number
		go filter(ch, ch1, prime)
		//Switch channel, to plug new channel
		ch = ch1
	}
}

//generate Generate next number
func generate(ch chan int){
	for i:=2;;i++{
		//Send a number
		ch <-i
	}
}

//filter Filter number if is prime or not
func filter(in, out chan int, prime int){
	for{
		//Read a number
		i := <-in
		if i%prime != 0 {
			//is not prime
			out <- i
		}
	}
}