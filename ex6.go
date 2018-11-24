package main

import "fmt"

func main() {
	str := "문자열"
	fmt.Printf("%s\n", *(&str))
}
