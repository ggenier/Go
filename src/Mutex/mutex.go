package main

import (
	"fmt"
	"runtime"

	//"sync"
	"os"
	"time"
)

func main() {
	runtime.GOMAXPROCS(4)
	//mutex := new(sync.Mutex)

	file, _ := os.Create("./logs.txt")
	file.Close()

	logCh := make(chan string, 50)

	go func() {
		for {
			msg, ok := <-logCh
			if ok {
				file, _ := os.OpenFile("./logs.txt", os.O_APPEND, os.ModeAppend)
				logTime := time.Now().Format(time.RFC3339)
				file.WriteString(logTime + " : " + msg)
				file.Close()
			} else {
				break
			}
		}
	}()

	mutex2 := make(chan bool, 1)

	for i := 1; i < 10; i++ {
		for j := 1; j < 10; j++ {
			//mutex.Lock()
			mutex2 <- true
			go func() {
				msg := fmt.Sprintf("%d + %d = %d\n", i, j, i+j)
				logCh <- msg
				fmt.Print(msg)
				//mutex.Unlock()
				<-mutex2
			}()
		}
	}

	fmt.Scanln()
}
