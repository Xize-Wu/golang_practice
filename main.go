package main

import "fmt"

func main() { //must have a main function that doesn't take arguments and don't do anything
	fmt.Println("Hello, world.")

	var whatToSay string 
	var i int 

	whatToSay = "Goodbye, cruel world."
	i = 5

	fmt.Println(whatToSay)
	fmt.Println("i is set to", i)

	whatWasSaid, theOtherThingThatWasSaid := saySomething()
	fmt.Println(whatWasSaid, theOtherThingThatWasSaid)
}

func saySomething() (string, string) {
	return "something", "else"
}
