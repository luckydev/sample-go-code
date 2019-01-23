package main

import (
	"fmt"
	"time"
)

func main() {
	for {
		fmt.Println("vim-go")
		d, _ := time.ParseDuration("5s")
		time.Sleep(d)
	}
}
