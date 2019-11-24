package main

import "fmt"

func main() {
	var text string
	fmt.Println("-------SHOP------")
	//fmt.Scan(`Welcome : `, text)
	var B string
	var A = "Welcome : "
	fmt.Scanln(&text, `Welcome : `)
	fmt.Println(A+text, &B)
	B = A + text
	fmt.Println(B)
}
