package main

import (
	"fmt"
	"strings"
)

type (
	message struct {
		To      []string
		From    string
		Content string
	}

	failedMessage struct {
		ErrorMessage    string
		originalMessage message
	}
)

func main() {

	fmt.Println("Test with a simple channel")
	//Test with a simple channel
	func() {
		//Ceeate a channel
		ch := make(chan string, 1) //Without 1, program deadlock

		//Add information into channel
		ch <- "Hello"

		//Display channel content
		fmt.Println(<-ch)
	}()

	fmt.Println("")
	fmt.Println("Send phrase word by word on a channel, and read then word by word from the channel")
	//Send phrase word by word on a channel, and read then word by word from the channel
	func() {
		phrase := "Once upon a time a princess encounter a Prince, in real life it does not append :)\n"

		//Split the phrase
		splitPhrase := strings.Split(phrase, " ")

		//Create a buffered channel
		//Second argument is number of information we send
		ch := make(chan string, len(splitPhrase)) //Buffered channel

		//Add split phrase in channel
		for _, word := range splitPhrase {
			ch <- word
		}

		//Read from channel and display complet phrase
		for i := 0; i < len(splitPhrase); i++ {
			fmt.Print(<-ch + " ")
		}
	}()

	fmt.Println("")
	fmt.Println("Same test with close channel")
	//Send phrase word by word on a channel, and read then word by word from the channel
	func() {
		phrase := "Once upon a time a princess encounter a Prince, in real life it does not append :)\n"

		//Split the phrase
		splitPhrase := strings.Split(phrase, " ")

		//Create a buffered channel
		//Second argument is number of information we send
		ch := make(chan string, len(splitPhrase)) //Buffered channel

		//Add split phrase in channel
		for _, word := range splitPhrase {
			ch <- word
		}

		close(ch)

		//Read from channel and display complet phrase
		/*
			for i := 0; i < len(splitPhrase); i++ {
				if msg, ok := <- ch; ok {// if there is more message
					fmt.Print(msg + " ")
			}else{
				break
			}
		*/

		//Better syntax
		for msg := range ch {
			fmt.Print(msg + " ")
		}

	}()

	//With two channels- version 1
	fmt.Println("")
	fmt.Println("Example with two channels - version 1")
	func() {

		//Create channels
		msgCh := make(chan message, 1)
		failCh := make(chan failedMessage, 1)

		//Create messages
		msg := message{
			To:      []string{"foo@goo.com"},
			From:    "from@goo.com",
			Content: "Keep it secret, keep it safe",
		}

		failMsg := failedMessage{
			ErrorMessage:    "Message intercepted by spider man",
			originalMessage: message{},
		}

		//Send messages on respective channel
		msgCh <- msg
		failCh <- failMsg

		//print both channel
		fmt.Println(<-msgCh)
		fmt.Println(<-failCh)

	}()

	//With two channels- version 2
	fmt.Println("")
	fmt.Println("Example with two channels - version 2")
	func() {

		//Create channels
		msgCh := make(chan message, 1)
		failCh := make(chan failedMessage, 1)

		//Create messages
		msg := message{
			To:      []string{"foo@goo.com"},
			From:    "from@goo.com",
			Content: "Keep it secret, keep it safe",
		}

		failMsg := failedMessage{
			ErrorMessage:    "Message intercepted by spider man",
			originalMessage: message{},
		}

		msgCh <- msg
		failCh <- failMsg

		//Only first case is executed
		select {
		case receiveMsg := <-msgCh:
			fmt.Println(receiveMsg)

		case failedMsg := <-failCh:
			fmt.Println(failedMsg)

		default:
			fmt.Println("No message on channel")
		}

		//Three cases are executed
		fmt.Println("")
		fmt.Println("Test with a for loop, to treat all messages")
		msgCh <- msg
		//failCh <- failMsg : Message is still in channel
		exitFor := false
		for {
			select {
			case receiveMsg := <-msgCh:
				fmt.Println(receiveMsg)

			case failedMsg := <-failCh:
				fmt.Println(failedMsg)

			default:
				fmt.Println("No message on channel")
				exitFor = true
			}
			if exitFor {
				break
			}
		}
	}()

}
