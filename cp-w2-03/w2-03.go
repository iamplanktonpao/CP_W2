package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	const name, age = "Kim", 22
	s := fmt.Sprint(name, " is ", age, " years old.\n")

	io.WriteString(os.Stdout, s) // Ignoring error for simplicity.

	const aa, tt = "sd", 35
	fmt.Println(len(aa))
}
