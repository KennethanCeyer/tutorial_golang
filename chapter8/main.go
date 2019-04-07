package main

import (
	"fmt"
	"sort"
)

const dividerBar = "================================================"
const (
	MenuSetDietNum  = 0
	MenuListDietNum = 1
	MenuGetDietNum  = 2
	MenuDelDietNum  = 3
	MenuSaveNum     = 4
	MenuLoadNum     = 5
	MenuExitNum     = 6

	MenuSetDiet  = "식단 입력"
	MenuListDiet = "모든 일자 식단 보기"
	MenuGetDiet  = "특정 일자 식단 보기"
	MenuDelDiet  = "특정 일자 식단 삭제"
	MenuSave     = "저장"
	MenuLoad     = "불러오기"
	MenuExit     = "끝내기"
)

var (
	Menus = map[int]string{
		MenuSetDietNum:  MenuSetDiet,
		MenuListDietNum: MenuListDiet,
		MenuGetDietNum:  MenuGetDiet,
		MenuDelDietNum:  MenuDelDiet,
		MenuSaveNum:     MenuSave,
		MenuLoadNum:     MenuLoad,
		MenuExitNum:     MenuExit,
	}
)

func getUser() User {
	var name string
	var gender int
	var age int

	validator := NewValidator()

	fmt.Print("먼저 이름을 알려주시겠어요?: ")
	fmt.Scanf("%s", &name)

	for true {
		fmt.Printf("성별을 입력해주세요 (%d: 남자, %d: 여자): ",
			Male, Female)
		fmt.Scanf("%d", &gender)

		if validator.IsValidGender(gender) {
			break
		}

		fmt.Printf("%d는 올바르지 않은 옵션입니다.\n", gender)
		fmt.Println("정보를 다시 입력해주세요")
		fmt.Println(dividerBar)
	}

	for true {
		fmt.Print("연령이 어떻게 되시나요?: ")
		fmt.Scanf("%d", &age)

		if validator.IsValidAge(age) {
			break
		}

		fmt.Printf("%d는 올바르지 않은 연령입니다.\n", gender)
		fmt.Println("정보를 다시 입력해주세요")
		fmt.Println(dividerBar)
	}

	fmt.Println(dividerBar)
	fmt.Println("정보 입력이 정상적으로 완료되었습니다. 감사합니다!")

	return User{Name: name, Gender: Gender(gender), Age: age}
}

func handleInterface(program *Program) {
	var choice int
	var menuKeys []int
	for key := range Menus {
		menuKeys = append(menuKeys, key)
	}
	sort.Ints(menuKeys)

	for true {
		fmt.Println(dividerBar)
		fmt.Println("[선택하기]")
		for _, key := range menuKeys {
			fmt.Printf("%d. %s\n", key, Menus[key])
		}
		fmt.Println(dividerBar)
		fmt.Printf(
			"위 메뉴 중 원하는 기능의 번호를 입력해주세요: (0~%d): ",
			len(Menus)-1)
		fmt.Scanf("%d", &choice)
		isExit := handleMenu(choice, program)
		if isExit {
			return
		}
	}
}

func handleMenu(choice int, program *Program) (isExit bool) {
	switch choice {
	case MenuSetDietNum:
		handleSetDiet(program)
	case MenuListDietNum:
		handleListDiet(program)
	case MenuGetDietNum:
		handleGetDiet(program)
	case MenuDelDietNum:
		handleDelDiet(program)
	case MenuSaveNum:
		handleSave(program)
	case MenuLoadNum:
		handleLoad(program)
	case MenuExitNum:
		fmt.Println("프로그램을 종료합니다.")
		return true
	default:
		fmt.Printf("%d는 지원하지 않는 메뉴입니다.\n", choice)
		fmt.Println("다시 선택해주세요.")
	}
	fmt.Println(dividerBar)
	return false
}

func main() {
	program := NewProgram()
	program.LoadLast()

	fmt.Println(dividerBar)
	fmt.Printf("안녕하세요 %s님!\n", program.User.Name)

	if len(program.Diets) > 0 {
		fmt.Printf("%d건의 칼로리 입력 데이터가 있네요.\n",
			len(program.Diets))
	}

	fmt.Println(dividerBar)

	if program.User.Name == DefaultUserName {
		user := getUser()
		program.SetUser(user)
	}
	handleInterface(program)
}
