package main

import (
	"fmt"
	"time"
)

func main() {
	for {
		fmt.Println("testing aws code build")
		d, _ := time.ParseDuration("5s")
		time.Sleep(d)
	}
}
