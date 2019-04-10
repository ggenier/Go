package main

import (
	"fmt"
)

func main(){
	//Test IF function
	testIf("TOTO")
	testIf("GG")
	testIf("JSE")

	//Test IF function with variable in IF
	testIfWithDefault("GG")

	//Test switch function
	testSwitch("GG")
	testSwitch("JSE")
	testSwitch("TOTO")

	//Test switch function with variable in swicth
	testSwitchWithDefault("GG")

	testSwitchManyValues("MF")
}

func testIf(valueTested string) (err error){

	fmt.Println("testIf function start")
	if valueTested == "GG"{
		fmt.Println("   => It's me")
	}else if (valueTested == "JSE"){
		fmt.Println("   => It's jeremie")
	}else{
		fmt.Println("   => It's someone else")
	}
	fmt.Println("testIf function end")

	return nil
}

func testIfWithDefault(valueTested string) (err error){

	fmt.Println("testIf function start")
	if number := 10;valueTested == "GG"{
		fmt.Println("   => It's me : ",number)
	}else if (valueTested == "JSE"){
		fmt.Println("   => It's jeremie")
	}else{
		fmt.Println("   => It's someone else : ",number)
	}
	fmt.Println("testIf function end")

	//number is not available here, only in if statement

	return nil
}

func testSwitch(valueTested string) (err error){

	fmt.Println("testSwitch function start")
	switch valueTested{
		case "GG" : fmt.Println("   => It's me")
		case "JSE" : fmt.Println("   => It's jeremie")
		default : fmt.Println("   => It's someone else")
	}
	fmt.Println("testSwitch function end")

	return nil
}

func testSwitchWithDefault(valueTested string) (err error){

	fmt.Println("testSwitch function start")
	switch number:=10 ; valueTested{
		case "GG" : fmt.Println("   => It's me : ", number)
		case "JSE" : fmt.Println("   => It's jeremie")
		default : fmt.Println("   => It's someone else")
	}
	fmt.Println("testSwitch function end")

	//number is not available here, only in if statement

	return nil
}

func testSwitchManyValues(valueTested string) (err error){

	fmt.Println("testSwitch function start")
	switch number := 10 ; valueTested{
		case "GG", "MF" : fmt.Println("   => It's GG OR MF : ", valueTested, " ", number)
		case "JSE" : fmt.Println("   => It's jeremie")
		default : fmt.Println("   => It's someone else")
	}
	fmt.Println("testSwitch function end")

	//number is not available here, only in if statement
	
	return nil
}