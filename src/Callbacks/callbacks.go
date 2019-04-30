package main

import (
	"fmt"
)

func main() {
	//create a new PO
	po := new(PurchaseOrder)
	//Indicate value of po
	po.value = 47.34
	//Create a channel on purchase order pointer
	ch := make(chan *PurchaseOrder)
	//Save PO
	go SavePo(po, ch)
	//Wait for callback, on save in new variable
	newPo := <-ch
	//print new po information
	fmt.Printf("PO : %v\n", newPo)
}

//PurchaseOrder Purchase order structure
type PurchaseOrder struct {
	number int
	value  float64
}

//SavePo Save purchase order in database
func SavePo(po *PurchaseOrder, callback chan *PurchaseOrder) {
	//attribute a number, and can do other thing
	po.number = 1234
	//send po in callback channel
	callback <- po
}
