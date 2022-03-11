package main

import "fmt"

type Cust struct {
	custid int
	name   string
}

//type ArrCust []Cust

func (c *Cust) set_cust() {
	fmt.Scan(&c.custid)
	fmt.Scan(&c.name)
}

func sortarr(arr []Cust) []Cust {
	var temp Cust
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
	return arr
}

func main() {
	var customers = make([]Cust, 5)
	var xyz []Cust
	for i := 0; i < 5; i++ {
		fmt.Println("Enter custid , Name:")
		customers[i].set_cust()
	}

	fmt.Println(customers)
	xyz = sortarr(customers)
	fmt.Println(xyz)

}
