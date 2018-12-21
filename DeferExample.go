package main

import "os"

func main() {
	println("line 1")
	defer println("line 2, but l am last")
	println("line 3")
	file, err := os.Open("dd")
	println(err)
	file.Close()
}
