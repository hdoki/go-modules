package main

import "fmt"

type customer struct {
	custid   int
	custname string
}

func main() {
	var customers [6]customer
	customers[0] = customer{1, "a"}
	customers[1] = customer{2, "b"}
	customers[2] = customer{3, "c"}
	customers[3] = customer{4, "d"}
	customers[4] = customer{5, "e"}
	customers[5] = customer{6, "f"}

	fmt.Println(customers)
}
