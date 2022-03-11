package main

import "fmt"

func set_list(c []int, i int) {
	fmt.Scan(&c[i])

}

func Binarysearch(k int, arr []int) bool {
	var b bool
	mid := len(arr) / 2
	if arr[mid] == k {
		return true
	} else {
		if arr[mid] > k {
			arr = arr[:mid]
			b = Binarysearch(k, arr)
		} else {
			if arr[mid] < k {
				arr = arr[mid+1:]
				b = Binarysearch(k, arr)
			} else {
				return false
			}
		}
	}
	return b
}

func main() {

	//var list []int
	var l, k int
	fmt.Println("Enter length of the list")
	fmt.Scan(&l)

	var list = make([]int, l)
	fmt.Println("Enter values")
	for i := 0; i < l; i++ {

		set_list(list, i)

	}
	fmt.Printf("Search List: %d", list)
	fmt.Printf("Enter search value")
	fmt.Scan(&k)
	var b bool
	b = Binarysearch(k, list)
	fmt.Println(b)

} //not working for false
