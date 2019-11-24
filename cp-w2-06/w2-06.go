package main

import "fmt"

func main() {
	var text string
	var t2 string
	fmt.Printf("My name is %s , i am %d year old \n", "Peter", 20)
	fmt.Printf("Computer Science %s", "#?")
	fmt.Printf("`I LOVE ` %s \n", "CODING")
	fmt.Scanf("What do you want : %s   \n"+"%s", &text, &t2)
	fmt.Println(len(t2))

}
