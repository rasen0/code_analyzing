package main

import (
	"flag"
	"fmt"
)

// E:\CodeSpace\nsq
func main() {
	var dir string
	flag.StringVar(&dir, "d", "", "refer target project path")
	flag.Parse()
	if dir == "" {
		fmt.Println("please entry project path that need to be parsed")
		return
	}
	// todo ...
}
