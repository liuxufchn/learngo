package main

import (
	"fmt"
	"learngo/hello"
)

func main() {
	max := hello.LongestSubStrWithoutRepeatingChars("hello")
	fmt.Println(max)
}
