package main

import (
	"fmt"
)

func sliiceSum(numbers []int) int{
	sum := 0
	for _, num := range numbers{
		sum += num
	}
	return sum
}

func main(){
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println("sum of", slice, "is" ,sliiceSum(slice))
}