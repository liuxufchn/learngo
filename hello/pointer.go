package hello

import "fmt"

func swap(a, b *int) {
	*a, *b = *b, *a
}

func Main04() {
	a, b := 3, 4
	swap(&a, &b)
	fmt.Printf("a = %d, b = %d", a, b)
}
