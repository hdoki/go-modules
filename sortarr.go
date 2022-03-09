package main

import "fmt"

type cust struct {
	custid int
	name   string
}
type arrCust []cust

func (c *cust) set_cust() {
	fmt.Scan(&c.custid)
	fmt.Scan(&c.name)
}

/* func sortarr(arr []cust) {
	var temp cust
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i].custid > arr[j].custid {
				temp = arr[i]
				arr[i] = arr[j]
				arr[j] = temp
			}

		}
	}
	fmt.Println("sorted array is", arr)
}
*/
func main() {
	var customers [5]cust
	for i := 0; i < 5; i++ {
		fmt.Println("Enter custid , Name:")
		customers[i].set_cust()
	}

	fmt.Println(customers)
	//fmt.Println(sortarr(customers))
}
