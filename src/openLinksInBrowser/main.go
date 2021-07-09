package main

import (
	"fmt"
	"os"
)

func main() {
	//separator := "\n"
	//
	//links, err := controls.Parser("./links.txt", separator)
	//if err != nil {
	//	fmt.Println("err", err)
	//}
	//
	//for i := 0; i < len(links); i++ {
	//	fmt.Println("links[i]", links[i])
	//	openbrowser(links[i])
	//}

	argsWithProg := os.Args
	correctArgs := append(argsWithProg[1:])
	command := correctArgs[0]
	fmt.Println("command", command)
	//var data string
	//
	//fmt.Fscan(os.Stdin, &data)
	//fmt.Println("enter", )
	//fmt.Println("data", data)

}
