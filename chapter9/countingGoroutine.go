package main

import (
	"fmt"
)

func couting1(i int) {
	fmt.Printf("counting1: %d번째 인덱스\n", i)
}

func couting2(i int) {
	fmt.Printf("counting2: %d번째 인덱스\n", i)
}

func main() {
	for i := 0; i < 5; i++ {
		go couting1(i)
	}

	for i := 0; i < 5; i++ {
		couting2(i)
	}
}
