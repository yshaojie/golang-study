package main

import (
	"fmt"
	"sync"
)

func init() {
	var waitGroup sync.WaitGroup
	waitGroup.Add(4)
	func() {
		fmt.Println("xxxx")
	}()

	fmt.Println("exec init function")
}

func main() {
	fmt.Println("Hello World")

}
