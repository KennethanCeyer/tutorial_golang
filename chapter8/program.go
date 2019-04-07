package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
	"path"
	"sort"
	"time"
)

type Program struct {
	Diets          []*Diet
	User           User
	HealthAnalyzer *HealthAnalyzer
}

const (
	SaveFilePath = "./diet"
	DietFile     = "diet"
	UserFile     = "user"
)

func NewProgram() *Program {
	program := new(Program)
	program.Diets = []*Diet{}
	program.User = User{Name: DefaultUserName}
	program.HealthAnalyzer = NewHealthAnalyzer(program.User)
	program.Init()
	return program
}

func (p *Program) getDietByDateTime(dateTime time.Time) *Diet {
	for _, diet := range p.Diets {
		timeDiff := time.Duration(
			math.Abs(float64(diet.DateTime.Sub(dateTime))))
		if timeDiff < time.Duration(time.Hour*24) {
			return diet
		}
	}
	return nil
}

func (p *Program) DisplayDiet() {
	for _, diet := range p.Diets {
		fmt.Printf("- 날짜: %v\n", diet.DateTime.Format("2006-01-02"))

		allCalories := 0
		for _, food := range diet.Foods {
			allCalories += food.Calories
		}

		recommendedCal, err :=
			p.HealthAnalyzer.GetRecommendedCalories()
		if err != nil {
			fmt.Println(err)
		}

		status, err := p.HealthAnalyzer.GetCaloryStatus(allCalories)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Printf("- 전체 칼로리: %d, 권장 칼로리: %d, 상태: %v\n",
			allCalories, recommendedCal, status)

		p.DisplayFoods(diet.Foods)
		fmt.Print("\n")
	}
}

func (p *Program) DisplayFoods(foods []Food) {
	for _, food := range foods {
		p.DisplayFood(food)
	}
}

func (p *Program) DisplayFood(food Food) {
	fmt.Printf(
		"\t- 시간: %-3v %-3d시, 음식: %-8v | 칼로리: %5d\n",
		GetTimeUnit(food.Time),
		food.Time,
		food.Name,
		food.Calories)
}

func (p *Program) Init() {
	if _, err := os.Stat(SaveFilePath); os.IsNotExist(err) {
		os.Mkdir(SaveFilePath, os.ModePerm)
	}
}

func (p *Program) AddDiet(dateTime time.Time) *Diet {
	diet := p.getDietByDateTime(dateTime)

	if diet == nil {
		diet = NewDiet([]Food{}, dateTime)
		p.Diets = append(p.Diets, diet)
	}
	return diet
}

func (p *Program) DeleteDiet(dateTime time.Time) {
	for i, diet := range p.Diets {
		timeDiff := time.Duration(
			math.Abs(float64(diet.DateTime.Sub(dateTime))))
		if timeDiff < time.Duration(time.Hour*24) {
			p.Diets = append(p.Diets[:i], p.Diets[i+1:]...)
			return
		}
	}
}

func (p *Program) DeleteDietFood(dateTime time.Time, foodName string) {
	diet := p.getDietByDateTime(dateTime)
	if diet == nil {
		return
	}
	for i, food := range diet.Foods {
		if food.Name == foodName {
			diet.Foods = append(diet.Foods[:i], diet.Foods[i+1:]...)
			return
		}
	}
}

func (p *Program) SetUser(user User) {
	p.User = user
	p.HealthAnalyzer = NewHealthAnalyzer(user)
}

func (p *Program) SaveDiet(filePath string) error {
	saveFilePath := path.Join(filePath, DietFile)
	file, err := os.OpenFile(
		saveFilePath, os.O_CREATE|os.O_WRONLY, os.FileMode(0644))
	if err != nil {
		return err
	}

	data, err := json.Marshal(p.Diets)
	if err != nil {
		return err
	}

	_, err = file.Write(data)
	if err != nil {
		return err
	}

	return nil
}

func (p *Program) SaveUser(filePath string) error {
	saveFilePath := path.Join(filePath, UserFile)
	file, err := os.OpenFile(
		saveFilePath, os.O_CREATE|os.O_WRONLY, os.FileMode(0644))
	if err != nil {
		return err
	}

	data, err := json.Marshal(p.User)
	if err != nil {
		return err
	}

	_, err = file.Write(data)
	if err != nil {
		return err
	}

	return nil
}

func (p *Program) Save() error {
	saveFileName := time.Now().Format("200601021504")
	savePath := path.Join(SaveFilePath, saveFileName)

	if _, err := os.Stat(savePath); os.IsNotExist(err) {
		os.Mkdir(savePath, os.ModePerm)
	}

	err := p.SaveDiet(savePath)
	if err != nil {
		return err
	}

	err = p.SaveUser(savePath)
	if err != nil {
		return err
	}

	return nil
}

func (p *Program) LoadList() []string {
	files, err := ioutil.ReadDir(SaveFilePath)
	if err != nil {
		log.Fatal(err)
	}

	var fileList []string
	for _, f := range files {
		fileList = append(fileList, f.Name())
	}
	return fileList
}

func (p *Program) LoadDiet(filePath string) ([]*Diet, error) {
	var diets []*Diet
	loadFilePath := path.Join(SaveFilePath, filePath, DietFile)
	file, err := os.OpenFile(
		loadFilePath, os.O_RDONLY, os.FileMode(0644))
	if err != nil {
		return diets, err
	}

	fileInfo, err := file.Stat()
	if err != nil {
		return diets, err
	}

	data := make([]byte, fileInfo.Size())
	_, err = file.Read(data)
	if err != nil {
		return diets, err
	}

	err = json.Unmarshal(data, &diets)
	if err != nil {
		return diets, err
	}

	return diets, nil
}

func (p *Program) LoadUser(filePath string) (User, error) {
	var user User
	loadFilePath := path.Join(SaveFilePath, filePath, UserFile)
	file, err := os.OpenFile(
		loadFilePath, os.O_RDONLY, os.FileMode(0644))
	if err != nil {
		return user, err
	}

	fileInfo, err := file.Stat()
	if err != nil {
		return user, err
	}

	data := make([]byte, fileInfo.Size())
	_, err = file.Read(data)
	if err != nil {
		return user, err
	}

	err = json.Unmarshal(data, &user)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (p *Program) Load(filePath string) error {
	diets, err := p.LoadDiet(filePath)
	if err != nil {
		return err
	}

	user, err := p.LoadUser(filePath)
	if err != nil {
		return err
	}

	p.Diets = diets
	p.SetUser(user)

	return nil
}

func (p *Program) LoadLast() error {
	loads := p.LoadList()
	sort.Sort(sort.Reverse(sort.StringSlice(loads)))

	if len(loads) == 0 {
		return nil
	}

	p.Load(loads[0])
	return nil
}
