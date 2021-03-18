package hello

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func printTextFile() {
	const filename = "hello/basic.go"
	if content, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", content)
	}
}

func convert2Binary(n int) {
	result := ""
	for ; n > 0; n /= 2 {
		lastBinary := n % 2
		result = strconv.Itoa(lastBinary) + result
	}
	fmt.Println("result = ", result)
}

func printTextFileUsingBufio() {
	const filename = "hello/basic.go"
	file, err := os.Open(filename)
	if err != nil {
		panic(fmt.Sprintf("fail to open file: %s\n", filename))
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func Main02() {
	printTextFile()
	convert2Binary(10)
	convert2Binary(6)
	printTextFileUsingBufio()
}
