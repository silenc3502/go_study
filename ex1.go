package main

import "fmt"

// go 는 사용하지 않는 변수를 엄격하게 검사하게 된다.
func main() {
	var 테스트 = "컹컹"
	ident := 10
	fmt.Println(테스트)
	println(ident)
}
