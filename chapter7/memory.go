package main

import "fmt"

func main() {
	// 메모리 공간에 name, phone 변수 할당
	var name string
	var phone string

	fmt.Print("혹시 성함이 어떻게 되시나요?: ")
	fmt.Scanf("%s", &name)

	fmt.Printf("감사합니다! %s님, 그러면 연락처는 어떻게 되세요?: ", name)
	fmt.Scanf("%s", &phone)

	// 메모리에 저장된 값 출력
	fmt.Print("\n")
	fmt.Println("================================================")
	fmt.Printf("Name:  %s\n", name)
	fmt.Printf("Phone: %s\n", phone)
	fmt.Println("================================================")

	// 함수가 종료되면서 메모리에 저장된 name, phone 변수 반환
}
