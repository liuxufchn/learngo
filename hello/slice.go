package hello

import "fmt"

func updateSlice(s []int) {
	s[0] = 100
}

func Main06() {
	arr := [9]int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	s1 := arr[2:]
	s2 := arr[:5]
	updateSlice(s1)
	fmt.Println(s1, s2)

	s3 := arr[2:6] // 2, 3, 4, 5
	s4 := s3[3:5]  // 5, 6
	fmt.Println(s4)

	fmt.Println("before extend.")
	fmt.Println("s1 = ", s1, "len(s1) = ", len(s1))
	fmt.Println("s2 = ", s2, "len(s2) = ", len(s2))
	fmt.Println("s3 = ", s3, "len(s3) = ", len(s3))
	fmt.Println("s4 = ", s4, "len(s4) = ", len(s4))
	s4 = append(s4, 10)
	updateSlice(s4)
	fmt.Println("after extend.")
	fmt.Println("s1 = ", s1, "len(s1) = ", len(s1))
	fmt.Println("s2 = ", s2, "len(s2) = ", len(s2))
	fmt.Println("s3 = ", s3, "len(s3) = ", len(s3))
	fmt.Println("s4 = ", s4, "len(s4) = ", len(s4))

}
