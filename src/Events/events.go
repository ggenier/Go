package main

import (
	"fmt"
)

func main() {
	//Create a button
	btn := MakeButton()

	//Create handler
	handlerOne := make(chan string, 2)
	handlerTwo := make(chan string, 2)

	//Add two different handler to button listeners
	btn.AddEventListener("click", handlerOne)
	btn.AddEventListener("click", handlerTwo)

	//Open channel for handler one
	go func() {
		for {
			msg := <-handlerOne
			fmt.Println("Hander One : " + msg)
		}
	}()

	//Open channel for handler two
	go func() {
		for {
			msg := <-handlerTwo
			fmt.Println("Hander Two : " + msg)
		}
	}()

	//Trigger event : two messages appears
	btn.TriggerEvent("click", "Button Clicked!")

	//Remove handler Two
	btn.RemoveEventListener("click", handlerTwo)

	//Trigger event : one message appears
	btn.TriggerEvent("click", "Button Clicked again!")

	//Wait for a click
	fmt.Scanln()
}

//Button structure
type Button struct {
	eventListeners map[string][]chan string
}

//MakeButton Create a button, return a reference of button
func MakeButton() *Button {
	result := new(Button)
	result.eventListeners = make(map[string][]chan string)
	return result
}

//AddEventListener Add a listener on button
func (btn *Button) AddEventListener(event string, responseChannel chan string) {
	//If button is not already in map
	if _, present := btn.eventListeners[event]; present {
		//Append listener
		btn.eventListeners[event] = append(btn.eventListeners[event], responseChannel)
	} else {
		//Replace channel
		btn.eventListeners[event] = []chan string{responseChannel}
	}
}

//RemoveEventListener Remove a listener on button
func (btn *Button) RemoveEventListener(event string, listenerChannel chan string) {
	//Check if button is in map
	if _, present := btn.eventListeners[event]; present {
		//Loop on all listeners
		for idx := range btn.eventListeners[event] {
			//If listener is the one to remove
			if btn.eventListeners[event][idx] == listenerChannel {
				//remove listener, there is no function to delete in a map
				btn.eventListeners[event] = append(btn.eventListeners[event][:idx], btn.eventListeners[event][idx+1:]...)
				break
			}
		}
	}
}

//TriggerEvent Trigger an event
func (btn *Button) TriggerEvent(event string, response string) {
	//Check if event is in map
	if _, present := btn.eventListeners[event]; present {
		//Loop on all listener to trigger event
		for _, handler := range btn.eventListeners[event] {
			//Sending message in channel
			go func(handler chan string) {
				handler <- response
			}(handler)
		}
	}
}
