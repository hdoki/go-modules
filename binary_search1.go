package main

import "fmt"

func set_list(c []int, i int) {
	fmt.Scan(&c[i])

}

func main() {

	//var list []int
	var l int
	fmt.Println("Enter length of the list")
	fmt.Scan(&l)

	var list = make([]int, l)
	fmt.Println("Enter values")
	for i := 0; i < l; i++ {

		set_list(list, i)

	}
	fmt.Printf("Search List: %d", list)

}
