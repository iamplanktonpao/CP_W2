package main

import "fmt"

func main() {
	var text string
	var num int
	var flot float32
	fmt.Print("WHAT DO U WANT", "DDDDDD \n")
	fmt.Println("HELLO", "I AM", "PLANKTON")
	fmt.Printf("Hello : %s I am \t %d", "Peter", 05)
	n, err := fmt.Scan(&text, &num, &flot)
	fmt.Println(text, num, flot)
	fmt.Println(n)
	fmt.Println(err)
}
