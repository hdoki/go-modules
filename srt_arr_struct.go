package main

import (
	"fmt"
	"sort"
)

type Student struct {
	Name string
	Age  int
}

// ByAge implements sort.Interface based on the Age field.
type ByAge []Student

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func (c *Student) set_student() {
	fmt.Scan(&c.Name)
	fmt.Scan(&c.Age)
}

func main() {
	var class [5]Student
	for i := 0; i < 5; i++ {
		fmt.Println("Enter Student Name , Age:")
		class[i].set_student()
	}

	sort.Sort(ByAge(class[:]))
	fmt.Println(class)
}
