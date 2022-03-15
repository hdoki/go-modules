package main

import (
	"fmt"
	"os"
	"sortdemo/packet"
	"strconv"
)

func main() {

	packets := []packet.Packet{}
	var a string
	for {
		var c, m string
		p := packet.Packet{}
		fmt.Print("Enter packet counter: ")
		fmt.Scanln(&c)
		ci, err := strconv.Atoi(c)
		if err != nil {
			fmt.Print("Packent counter should be a number")
			os.Exit(1)
		}
		p.Count = ci
		fmt.Print("Enter packet message: ")
		fmt.Scanln(&m)
		p.Msg = m
		packets = append(packets, p)
		fmt.Print("Press x and enter if done providing inputs: ")
		fmt.Scanf("%s", &a)
		if a == "x" {
			packet.Sort(packets)
			fmt.Println("Sorted packets: ", packets)
			break
		}

	}
}
