package main

import "unicode/utf8"

func main() {
	한국어 := "안녕"

	println(한국어)
	println(utf8.RuneCountInString(한국어))
}
