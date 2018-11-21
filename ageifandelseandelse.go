package main

import "fmt"

func main() {
	age := -1

	if age > 1 && age < 13 {
		fmt.Println("You are young")
	}
	if age > 12 && age < 20 {
		fmt.Println("you're a teenager")
	}
	if age > 19 && age < 30 {
		fmt.Println("you are in your twenties")
	}
	if age > 29 && age < 40 {
		fmt.Println("you are in you thirties")
	}
	if age > 39 && age < 50 {
		fmt.Println("Now you are going older")
	}
	if age > 49 && age < 100 {
		fmt.Println("you are old now")
	} else {
		fmt.Println("You are immortal bro..")
	}
}
