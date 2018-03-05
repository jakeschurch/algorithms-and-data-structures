package main

import (
	"fmt"

	"github.com/jakeschurch/algorithms-and-data-structures/algos"
)

// 0 1 1 2 3 5 8 13 21
func main() {
	f := algos.MemoFib(50)
	fmt.Println(f)
}
