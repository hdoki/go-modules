package packet

import "fmt"

type Packet struct {
	Count int
	Msg   string
}

func Sort(p []Packet) error {
	n := len(p)
	swapped := true
	for swapped {
		// set swapped to false
		swapped = false
		// iterate through all of the elements in our list
		for i := 0; i < n-1; i++ {
			// if the current element is greater than the next
			// element, swap them
			if p[i].Count > p[i+1].Count {
				// log that we are swapping values for posterity
				fmt.Println("Swapping")
				// swap values using Go's tuple assignment
				p[i], p[i+1] = p[i+1], p[i]
				// set swapped to true - this is important
				// if the loop ends and swapped is still equal
				// to false, our algorithm will assume the list is
				// fully sorted.
				swapped = true
			}
		}
	}
	return nil
}
