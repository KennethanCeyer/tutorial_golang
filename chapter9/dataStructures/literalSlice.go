package main

import "fmt"

func main() {
	// 리터럴(literal) 슬라이스
	literalSlice := []int{1, 2, 3}

	// make를 이용한 슬라이스
	makeSlice := make([]int, 3)
	for i := 0; i < len(makeSlice); i++ {
		makeSlice[i] = i + 1
	}

	fmt.Println("================================================")
	fmt.Println("[literalSlice]")
	fmt.Printf("len: %d, cap: %d\n", len(literalSlice), cap(literalSlice))
	fmt.Printf("value: %v\n", literalSlice)
	fmt.Println("================================================")
	fmt.Print("\n")

	fmt.Println("================================================")
	fmt.Println("[makeSlice]")
	fmt.Printf("len: %d, cap: %d\n", len(makeSlice), cap(makeSlice))
	fmt.Printf("value: %v\n", makeSlice)
	fmt.Println("================================================")
}
