package main

import "fmt"

func main() {
	// 리터럴(literal) 슬라이스
	literalSlice := []int{1, 2, 3, 4, 5, 6}

	fmt.Println("================================================")
	fmt.Println("[literalSlice]")
	fmt.Printf("len: %d, cap: %d\n", len(literalSlice), cap(literalSlice))
	fmt.Printf("value: %v\n", literalSlice)
	fmt.Println("================================================")
}
