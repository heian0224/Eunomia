package main

import (
	"fmt"
	"time"
)

func main() {
	for true {
		fmt.Println("Infinite loop")
		time.Sleep(1000)
	}
}
