package main

import (
	"fmt"
	"sync"
	"time"
)

func runHeavyJob(wg *sync.WaitGroup) {
	// runHeavyJob 함수가 종료될 때 대기 그룹의 작업을 하나 끝냈다고 알림
	defer wg.Done()

	fmt.Println("무거운 작업 실행 중")
	time.Sleep(time.Second * 5)
	fmt.Println("무거운 작업 실행 완료")
}

func main() {
	// 대기 그룹 생성
	wg := sync.WaitGroup{}
	// 대기 그룹에 작업을 하나 추가
	wg.Add(1)

	fmt.Println("main 함수 실행")
	go runHeavyJob(&wg)
	fmt.Println("main 함수 종료")

	// 대기 그룹의 모든 작업이 끝나기를 기다림
	wg.Wait()
}
