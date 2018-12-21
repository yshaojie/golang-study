package main

import (
	"math/rand"
	"sync"
	"time"
)

var ch = make(chan string)

func main() {
	var sy sync.WaitGroup
	sy.Add(1)
	product()
	consumer()
	func() {
		sy.Wait()
	}()
}

func product() {
	for i := 0; i < 3; i++ {
		go func() {
			for {
				time.Sleep(3 * time.Second)
				content := randomString(rand.Intn(16))
				println("product content=" + content)
				ch <- content
			}
		}()
	}
}

func consumer() {
	for i := 0; i < 3; i++ {
		go func() {
			for content := range ch {
				println("consumer content=" + content)
			}
		}()
	}
}

func randomString(len int) string {
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		bytes[i] = byte(65 + rand.Intn(25))
	}
	return string(bytes)
}
