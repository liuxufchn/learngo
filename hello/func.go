package hello

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

func eval(a, b int, op string) int {
	switch op {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		return a / b
	default:
		panic("unsupported operation:" + op)
	}
}

func div(a, b int) (int, int) {
	return a / b, a % b
}

func apply(op func(int, int) int, a, b int) (result int) {
	result = op(a, b)
	pointer := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(pointer).Name()
	fmt.Printf("calling function: %s(%d, %d);\n", opName, a, b)
	return
}

func Main03() {
	result := apply(func(a int, b int) int {
		return int(math.Pow(float64(a), float64(b)))
	}, 10, 2)
	fmt.Println(result)
}
