package main

import (
	"fmt"
	"reflect"
)

func main() {
	var array = [3]int{1, 2, 3}
	fmt.Println(reflect.TypeOf(array))
	sli := []int{1, 2, 3}
	fmt.Println(reflect.TypeOf(sli))
	slice := make([]int, 3, 5)
	fmt.Println(reflect.TypeOf(slice))
}
