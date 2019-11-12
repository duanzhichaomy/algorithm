package main

import (
	"algorithm/training_7"
	"fmt"
)

func main() {
	w := [][]int{
		{1,3,1},{1,5,1},{4,2,1},
	}
	s := training_7.MinPathSum(w)
	fmt.Println(s)
}