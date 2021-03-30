package hello

import "fmt"

func swap(a, b *int) {
	*a, *b = *b, *a
}

func changeSlice(s []int) {
	s = make([]int, 10)
	s[1] = 10
	fmt.Println("s in changeSlice: ", s)
}

func Main04() {
	a, b := 3, 4
	swap(&a, &b)
	fmt.Printf("a = %d, b = %d\n", a, b)
	s := []int{1, 2, 3, 4, 5}
	changeSlice(s)
	fmt.Println("s = ", s)
}
