package main

import "fmt"

func main() {
	/*	data := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

		s1 := data[8:]
		s2 := data[:5]
	*/
	data2 := []int{5, 6, 7, 8, 9}

	d1 := data2[:4:4]
	d2 := data2[3:5]
	fmt.Println("d1", d1[2:], cap(d1))
	fmt.Println("d2", d2, cap(d2))
	copy(d1[2:], d2)

	//	copy(s2[4:], s1)
	//	fmt.Println("s1= ", s1, "\n", "s2= ", s2)

	fmt.Println(d1)

}
