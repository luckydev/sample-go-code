package main

import (
	"fmt"
	"time"
)

func main() {
	for {
		fmt.Println("testing aws code build")
		fmt.Println("testing aws code pipeline as well")
		d, _ := time.ParseDuration("5s")
		time.Sleep(d)
	}
}
