package main

import (
	"net/http"
	"io/ioutil"
	"encoding/xml"
	"fmt"
	"time"
	"sync"
	"runtime"
)

type (
	 xmlResponse struct {
		Status string
		Name string
		Symbol string
		LastPrice float32
		Change float32
		ChangePercent float32
		Timestamp string
		MSDate int
		MarketCap int
		Volume int
		ChangeYTD float32
		ChangePercentYTD float32
		High float32
		Low float32
		Open int
	}
)

var (
	//Global I want to use everywhere
	waitGrp sync.WaitGroup
)

func main() {

	//List of entreprise we want price
	entreprisesList := []string{
		"googl",
		"msft",
		"aapl",
		"bbry",
		"hpq",
		"vz",
		"t",
		"tmus",
		"s",
	}

	//Begin of the request
	start := time.Now()
	
	//Modify number of CPU
	runtime.GOMAXPROCS(4)

	//Add number of goroutines
	waitGrp.Add(9)

	for _, entreprise := range(entreprisesList){
		go func(entreprise string) {
			defer waitGrp.Done()//My go routine is done
			//Connection to the WS
			queryURL := fmt.Sprintf("http://dev.markitondemand.com/Api/v2/Quote?symbol=%s", entreprise)
			//resp, _ := http.Get("http://dev.markitondemand.com/Api/v2/Quote?symbol=" + entreprise)
			resp, _ := http.Get(queryURL)

			//Close response after execution
			defer resp.Body.Close()

			//Read answer from response
			body, _ := ioutil.ReadAll(resp.Body)

			//CReation of a variable of go structure
			xmlResponse := new(xmlResponse)

			//Conversion response to XML
			xml.Unmarshal(body, &xmlResponse)

			//Print information
			fmt.Printf(entreprise + " -> Name : %s -> Last price : %.2f\n", xmlResponse.Name, xmlResponse.LastPrice)
		}(entreprise)
	}

	waitGrp.Wait() //wait end of execution of my two goroutines

	//Time elapsed
	elapsed := time.Since(start)

	fmt.Printf("Execution time %s\n", elapsed)

}