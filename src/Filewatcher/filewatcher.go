package main

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"
)

type (
	invoice struct {
		number              string
		amount              float64
		purchaseOrderNumber int
		invoiceDate         time.Time
	}
)

const sourceDirectory = "./invoices" //Invoice directory
const archiveDirectory = "./archive" //Archive directory for invoices treated

func main() {

	runtime.GOMAXPROCS(4)

	//Never stop, we watch folder
	for {
		//Open directory to watch
		dir, _ := os.Open(sourceDirectory)

		//Read all files in directory
		files, _ := dir.Readdir(-1) //-1 => read all files in directory, if Y > 0, we read only Y files

		//We loop on files
		for _, file := range files {
			//Full path of file
			fullfilePath := sourceDirectory + "/" + file.Name()

			//Open the file, like a directory :-)
			f, _ := os.Open(fullfilePath)

			//Read data from file
			data, _ := ioutil.ReadAll(f)

			//Close the file
			f.Close()

			//Move file to archive directory : old -> new
			os.Rename(fullfilePath, archiveDirectory+"/"+file.Name())

			//go routine
			go func(data string) {

				//read and convert data
				reader := csv.NewReader(strings.NewReader(data))

				//Read all records
				records, _ := reader.ReadAll()
				for _, record := range records {
					invoice := new(invoice)
					//Filles invoice
					invoice.number = record[0]
					invoice.amount, _ = strconv.ParseFloat(record[1], 64)
					invoice.purchaseOrderNumber, _ = strconv.Atoi(record[2])
					unixTime, _ := strconv.ParseInt(record[3], 10, 64)
					invoice.invoiceDate = time.Unix(unixTime, 0)

					//Treatment, here only display
					fmt.Printf("Invoice number : %v, amount : $%.2f, order number : %v, invoice date : %s\n", invoice.number, invoice.amount, invoice.purchaseOrderNumber, invoice.invoiceDate)
				}
			}(string(data))
		}
	}
}
