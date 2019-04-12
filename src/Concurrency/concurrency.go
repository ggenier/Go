package main

import (
	"fmt"
	"time"
	"sync"
)

var (
	//Global I want to use everywhere
	waitGrp sync.WaitGroup
)

func main(){

	waitGrp.Add(2) //Add two goroutines

	go showMessageWithSleep("Message with sleep")
	go showMessageWithoutSleep("Message without sleep")

	waitGrp.Wait() //wait end of execution of my two goroutines
}

func showMessageWithSleep(message string){
	defer waitGrp.Done()//My go routine is done

	time.Sleep(5 * time.Second) //5 seconds sleep
	fmt.Println(message)
}

func showMessageWithoutSleep(message string){
	defer waitGrp.Done()//My go routine is done

	fmt.Println(message)
}