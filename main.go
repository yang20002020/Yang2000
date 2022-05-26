package main

import (
	SUB "00/add"
	"fmt"
	"os"
)

func main() {
	r := SUB.Add(10, 20)
	fmt.Println("A=", r)
	cmds := os.Args
	for key, value := range cmds {
		fmt.Println("key=", key, "value=", value)
	}
}
