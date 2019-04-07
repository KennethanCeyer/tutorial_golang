package main

import (
	"fmt"
	"time"
)

func runHeavyJob() {
	fmt.Println("무거운 작업 실행 중")
	time.Sleep(time.Second * 5)
	fmt.Println("무거운 작업 실행 완료")
}

func main() {
	fmt.Println("main 함수 실행")
	runHeavyJob()
	fmt.Println("main 함수 종료")
}
