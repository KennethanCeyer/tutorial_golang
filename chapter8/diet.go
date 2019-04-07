package main

import (
	"fmt"
	"time"
)

type Diet struct {
	Foods    []Food
	DateTime time.Time
}

func NewDiet(foods []Food, dateTime time.Time) *Diet {
	diet := new(Diet)
	diet.Foods = foods
	diet.DateTime = dateTime
	return diet
}

func (d *Diet) isExists(foodName string) bool {
	for _, food := range d.Foods {
		if food.Name == foodName {
			return true
		}
	}
	return false
}

func (d *Diet) AddFood(foodName string, calory int, hour int) error {
	isExsits := d.isExists(foodName)
	if isExsits {
		return fmt.Errorf("%s 음식은 이미 등록되어있습니다\n", foodName)
	}
	d.Foods = append(
		d.Foods, Food{foodName, calory, hour})
	return nil
}

func (d *Diet) DeleteFood(foodName string) {
	for i, food := range d.Foods {
		if food.Name == foodName {
			d.Foods = append(d.Foods[:i], d.Foods[i+1:]...)
			return
		}
	}
}

func (d *Diet) UpdateFoodName(oldName string, newName string) {
	for _, food := range d.Foods {
		if food.Name == oldName {
			food.Name = newName
		}
	}
}
