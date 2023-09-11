package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello, playground")

	go fmt.Println("hello from goroutine")

	fmt.Println("end")
	time.Sleep(1 * time.Second)

}
