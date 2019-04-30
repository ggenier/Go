package main

import (
	"errors"
	"fmt"
	"time"
)

func main() {
	//create a new PO
	po := new(PurchaseOrder)
	//Indicate value of po
	po.value = 47.34

	//Save purchase with promise
	SavePo(po, false).Then(func(obj interface{}) error {
		//Success part
		po := obj.(*PurchaseOrder)
		fmt.Printf("Purchase order with ID %v\n", po.number)

		return nil

		//In this case, second promise diplay error message because first promise failed
		//return errors.New("First promise failed")

	}, func(err error) {
		//Failure part
		fmt.Printf("Failed to save purchase order : " + err.Error() + "\n")
	}).Then(func(obj interface{}) error {
		//Success part
		fmt.Printf("Second promise success\n")
		return nil
	}, func(err error) {
		//Failed part
		fmt.Println("Second promise failed:" + err.Error())
	})

	fmt.Scanln()
}

//Promise structure
type Promise struct {
	successChannel chan interface{}
	failureChannel chan error
}

//PurchaseOrder Purchase order structure
type PurchaseOrder struct {
	number int
	value  float64
}

//SavePo Save purchase order in database
func SavePo(po *PurchaseOrder, shouldFail bool) *Promise {
	result := new(Promise)

	result.successChannel = make(chan interface{}, 1)
	result.failureChannel = make(chan error, 1)

	go func() {
		//To simulate a timed out
		//time.Sleep(2 * time.Second)

		if shouldFail {
			//Simulate a problem during saving
			result.failureChannel <- errors.New("Failed to save purchase order")
		} else {
			po.number = 1234
			result.successChannel <- po
		}
	}()

	return result
}

//Then Execute promise
func (prom *Promise) Then(sucess func(interface{}) error, failure func(error)) *Promise {
	//Create promise
	result := new(Promise)

	//Create two channel for sucess and failure
	result.successChannel = make(chan interface{}, 1)
	result.failureChannel = make(chan error, 1)

	//time.After create a channel for timeout
	timeout := time.After(1 * time.Second)

	go func() {
		select {
		//If success
		case obj := <-prom.successChannel:
			newErr := sucess(obj)
			if newErr == nil {
				result.successChannel <- obj
			} else {
				result.failureChannel <- newErr
			}

		//if error
		case err := <-prom.failureChannel:
			failure(err)
			result.failureChannel <- err

		case <-timeout:
			failure(errors.New("Promise timed out"))
		}

	}()

	return result

}
