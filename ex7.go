package main

import "io/ioutil"
import "fmt"

func main() {
	if txt, err := ioutil.ReadFile("./ex2.go"); err == nil {
		println(txt)
		fmt.Printf("%s\n", txt)
	} else {
		println(err)
	}
}
