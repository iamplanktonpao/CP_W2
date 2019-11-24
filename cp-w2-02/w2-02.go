package main

import "fmt"

func main() {
	var text string
	fmt.Print(`love :  `)
	fmt.Scan(&text)
	fmt.Println(`Answer  : "`, text /*p*/, `" Go `)
	fmt.Println("-------next------- \r")
	fmt.Println("GO\bOO")
	fmt.Println("GO\tOO")
	fmt.Println("G\rOOO")
	fmt.Println("GO\fOO")
	fmt.Println("GOOO")

	const a = "sa"

	var x string
	x = "KKKK"
	fmt.Println(x)

	var y = "hello world"
	fmt.Println(y)

	z := "hello world"
	fmt.Println(z)
}
