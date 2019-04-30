package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"
	"sync"
)

func main() {
	start := time.Now()

	/*orders := extract()
	orders = transform(orders)
	load(orders)*/

	extractChannel := make(chan *Order)
	transformChannel := make(chan *Order)
	doneChannel := make(chan bool)

	go extract(extractChannel)
	go transform(extractChannel, transformChannel)
	go load(transformChannel, doneChannel)

	<-doneChannel
	fmt.Println(time.Since(start))

}

//Product structure
type Product struct {
	partNumber string
	unitCost   float64
	unitPrice  float64
}

//Order structure
type Order struct {
	customerNumber int
	partNumber     string
	quantity       int
	unitCost       float64
	unitPrice      float64
}

//func extract() []*Order{
func extract(ch chan *Order) {
	f, _ := os.Open("./sources/orders.txt")
	defer f.Close()

	r := csv.NewReader(f)

	for record, err := r.Read(); err == nil; record, err = r.Read() {
		order := new(Order)
		order.customerNumber, _ = strconv.Atoi(record[0])
		order.partNumber = record[1]
		order.quantity, _ = strconv.Atoi(record[2])

		ch <- order
	}

	close(ch)
}

//func transform(orders []*Order) []*Order {
func transform(extractChannel, transformChannel chan *Order) {
	f, _ := os.Open("./sources/productList.txt")
	defer f.Close()

	r := csv.NewReader(f)

	records, _ := r.ReadAll()
	productList := make(map[string]*Product)

	for _, record := range records {
		product := new(Product)

		product.partNumber = record[0]
		product.unitCost, _ = strconv.ParseFloat(record[1], 64)
		product.unitPrice, _ = strconv.ParseFloat(record[2], 64)
		productList[product.partNumber] = product
	}

	mux:=sync.Mutex{}
	numMessages := 0

	for o := range extractChannel {
		mux.Lock()
		numMessages++
		mux.Unlock()
		go func(o *Order) {
			time.Sleep(3 * time.Millisecond)
			o.unitCost = productList[o.partNumber].unitCost
			o.unitPrice = productList[o.partNumber].unitPrice
			transformChannel <- o
			mux.Lock()
			numMessages--
			mux.Unlock()
		}(o)
	}

	//return <- orders
	for numMessages > 0 {
		time.Sleep(1 * time.Millisecond)
	}

	close(transformChannel)
}

//func load(orders []*Order){
func load(transformChannel chan *Order, doneChannel chan bool) {
	f, _ := os.Create("./dest/report.txt")
	defer f.Close()

	fmt.Fprintf(f, "%20s%15s%12s%12s%15s%15s\n",
		"Part Number", "Quantity",
		"Unit Cost", "Unit Price",
		"Total Cost", "Total Price")

	mux:=sync.Mutex{}
	numMessages := 0

	//for _, o := range orders{
	for o := range transformChannel {
		mux.Lock()
		numMessages++
		mux.Unlock()
		go func(o *Order) {
			time.Sleep(1 * time.Millisecond)
			fmt.Fprintf(f, "%20s %15d %12.2f %12.2f %15.2f %15.2f\n",
				o.partNumber, o.quantity, o.unitCost, o.unitPrice,
				o.unitCost*float64(o.quantity), o.unitPrice*float64(o.quantity))
			mux.Lock()
			numMessages--
			mux.Unlock()
		}(o)
	}

	for numMessages > 0 {
		time.Sleep(1 * time.Millisecond)
	}

	doneChannel <- true
}
