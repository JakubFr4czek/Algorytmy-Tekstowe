package main

import (
	"fmt"
)

func main() {

	slice := []int{1, 2, 3, 4}

	slice = append([]int{3}, slice...)

	fmt.Printf("%d\n", slice)

}
