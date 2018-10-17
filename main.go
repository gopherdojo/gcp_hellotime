package main

import (
	"fmt"
	"time"
)

func main() {
	for {
		fmt.Printf("Hello %s\n", time.Now())
		time.Sleep(5 * time.Second)
	}
}
