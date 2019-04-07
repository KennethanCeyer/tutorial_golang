package main

import (
	"fmt"
	"os"
	"strings"
)

// 파일에 저장된 데이터를 불러오는 함수
// 파일이 없으면 빈 문자열을 반환
func loadData() (name string, phone string) {
	// output.txt 이름의 파일을 열어봄
	// os.O_CREATE: 파일이 없다면 생성함
	// os.O_RDWR:   읽기 전용으로 파일을 염
	// FileMode 0644: USER: 읽기/쓰기, GROUP: 읽기, ANONYMOUS: 읽기
	file, err := os.OpenFile("output.txt",
		os.O_CREATE|os.O_RDONLY, os.FileMode(0644))

	// 파일을 여는데 실패했다면 오류 메시지 출력
	if err != nil {
		fmt.Println(err)
		return
	}

	// 파일의 정보를 가지고 옴
	f, err := file.Stat()

	// 파일의 정보를 가지고 오는데 실패했다면 오류 메시지 출력
	if err != nil {
		fmt.Println(err)
		return
	}

	// 파일의 내용 크기에 맞게 byte 슬라이스 생성
	data := make([]byte, f.Size())

	// 방금 생성한 byte 슬라이스에 파일 내용을 옮김
	_, err = file.Read(data)

	// []byte 자료형을 string 타입으로 변환(캐스팅)
	value := string(data)

	// 저장된 값이 없다면 name, phone을 빈 문자열로 반환
	if value == "" {
		return "", ""
	}

	// 내용이 있다면 콤마(,)를 기준으로 문자열을 배열로 나눔
	// 최대 2개의 요소로 분리
	splitValues := strings.SplitN(value, ",", 2)

	// 분리된 배열의 요소를 각각 변수에 할당
	// 0번 인덱스 요소를 name, 1번 인덱스 요소를 phone으로 지정
	nameValue, phoneValue := splitValues[0], splitValues[1]

	// 읽어온 name, phone 값을 반환
	return nameValue, phoneValue
}

// 파일에 데이터를 저장하는 함수
func saveData(name string, phone string) {
	// output.txt 이름의 파일을 열어봄
	// os.O_CREATE: 파일이 없다면 생성함
	// os.RDWR:     읽기/쓰기 전용으로 파일을 염
	// os.O_TRUNC   파일을 덮어쓰기 형태로 저장
	// FileMode 0644: USER: 읽기/쓰기, GROUP: 읽기, ANONYMOUS: 읽기
	file, err := os.OpenFile("output.txt",
		os.O_CREATE|os.O_RDWR|os.O_TRUNC, os.FileMode(0644))

	// 파일을 여는데 실패했다면 오류 메시지 출력
	if err != nil {
		fmt.Println(err)
		return
	}

	// 이후에 오류가 발생할 경우 file 디스크립터를 닫음
	defer file.Close()

	// name과 value를 사이에 콤마(,)를 넣은 형태로 value 변수에 할당
	value := fmt.Sprintf("%s,%s", name, phone)

	// value 변수를 []byte로 변환(캐스팅)하여 파일에 쓰기
	_, err = file.Write([]byte(value))

	// 쓰기에 실패했다면 오류 메시지를 출력
	if err != nil {
		fmt.Println(err)
		return
	}
}

func main() {
	// 메모리 공간에 name, phone 변수 할당
	name, phone := loadData()

	// 저장된 값이 없을 때 사용자로부터 입력을 받음.
	if name == "" {
		fmt.Print("혹시 성함이 어떻게 되시나요?: ")
		fmt.Scanf("%s", &name)
	}

	if phone == "" {
		fmt.Printf("연락처는 어떻게 되세요?: ")
		fmt.Scanf("%s", &phone)
	}

	// 저장된 값 출력
	fmt.Println("================================================")
	fmt.Printf("Name:  %s\n", name)
	fmt.Printf("Phone: %s\n", phone)
	fmt.Println("================================================")

	saveData(name, phone)
}
