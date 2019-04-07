package main

import (
	"fmt"
	"log"
)

func handleSetDiet(program *Program) {
	var date string
	var time int
	var foodName string
	var calory int

	validator := NewValidator()

	fmt.Println(dividerBar)
	for true {
		fmt.Print(
			"어떤 날짜의 칼로리를 작성하고 싶으세요? " +
				"(2020-06-01 형태로 입력): ")
		fmt.Scanf("%s", &date)

		if validator.IsValidDate(date) {
			break
		}

		fmt.Println(dividerBar)
		fmt.Printf("'%s'는 유효하지 않은 날짜 입력 값입니다.\n", date)
		fmt.Println("다시 입력해주세요.")
	}

	for true {
		fmt.Print("시간을 입력해주세요 (0~23): ")
		fmt.Scanf("%d", &time)

		if validator.IsValidTime(time) {
			break
		}

		fmt.Println(dividerBar)
		fmt.Printf("%d시는 유효하지 않은 시간입니다.\n", time)
		fmt.Println("다시 입력해주세요.")
	}

	fmt.Print("어떤 음식을 드셨나요?: ")
	fmt.Scanf("%s", &foodName)

	fmt.Print("몇 칼로리 였나요?: ")
	fmt.Scanf("%d", &calory)

	dateTime, err := GetDateTime(date)
	if err != nil {
		fmt.Println(err)
		return
	}

	diet := program.AddDiet(dateTime)
	err = diet.AddFood(foodName, calory, time)
	if err != nil {
		fmt.Println(err)
	}
}

func handleListDiet(program *Program) {
	program.DisplayDiet()
}

func handleGetDiet(program *Program) {
	var date string

	validator := NewValidator()

	for true {
		fmt.Print("보고싶은 날짜를 입력해주세요: ")
		fmt.Scanf("%s", &date)

		if validator.IsValidDate(date) {
			break
		}

		fmt.Println(dividerBar)
		fmt.Printf("'%s'는 유효하지 않은 날짜 입력 값입니다.\n", date)
		fmt.Println("다시 입력해주세요.")
	}

	dateTime, err := GetDateTime(date)
	if err != nil {
		log.Println(err)
		return
	}
	diet := program.getDietByDateTime(dateTime)
	program.DisplayFoods(diet.Foods)
}

func handleDelDiet(program *Program) {
	var date string

	validator := NewValidator()

	for true {
		fmt.Print(
			"삭제하고 싶은 날짜를 입력해주세요: " +
				"(2020-06-01 형태로 입력): ")
		fmt.Scanf("%s", &date)

		if validator.IsValidDate(date) {
			break
		}

		fmt.Println(dividerBar)
		fmt.Printf("'%s'는 유효하지 않은 날짜 입력 값입니다.\n", date)
		fmt.Println("다시 입력해주세요.")
	}

	dateTime, err := GetDateTime(date)
	if err != nil {
		fmt.Println(err)
	}

	program.DeleteDiet(dateTime)
	fmt.Printf("%s 날짜의 식단 정보가 삭제되었습니다.\n",
		dateTime.Format("2006-01-02"))
}

func handleSave(program *Program) {
	err := program.Save()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("정보가 정상적으로 저장되었습니다!")
}

func handleLoad(program *Program) {
	var choices []int
	var choice int

	loadList := program.LoadList()

	for i, loadName := range loadList {
		choices = append(choices, i)
		fmt.Printf("%2d. %v\n", i, loadName)
	}

	for true {
		fmt.Print("몇 번을 불러올까요?: ")
		fmt.Scanf("%d", &choice)

		if choice >= 0 && choice < len(choices) {
			break
		}

		fmt.Printf("%d는 허용하지 않는 옵션입니다.\n", choice)
		fmt.Println("다시 입력해주세요.")
	}

	err := program.Load(loadList[choice])
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s를 정상적으로 불러왔습니다!\n", loadList[choice])
}
