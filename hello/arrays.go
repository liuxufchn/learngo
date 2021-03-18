package hello

import "fmt"

func printArray(arr *[5]int) {
	for i, v := range arr {
		fmt.Printf("%d, %d\t", i, v)
	}
	fmt.Println()
}

func Main05() {
	var arr01 [5]int
	arr02 := [5]int{1, 2, 3, 4, 5}
	arr03 := [...]int{2, 4, 6, 8, 10}

	printArray(&arr01)
	printArray(&arr02)
	printArray(&arr03)

}
