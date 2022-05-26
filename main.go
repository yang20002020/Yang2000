package main

import (
	SUB "00/add"
	"fmt"
	"os"
)

type mes struct {
	a int
}

func main() {
	r := SUB.Add(10, 20)
	fmt.Println("A=", r)
	cmds := os.Args
	for key, value := range cmds {
		fmt.Println("key=", key, "value=", value)
		fmt.Println("*********")
	}
}
