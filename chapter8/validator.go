package main

import "regexp"

type Validator struct {
}

var validator *Validator

func NewValidator() *Validator {
	if validator == nil {
		validator = new(Validator)
	}

	return validator
}

func (v Validator) IsValidDate(date string) bool {
	match, _ := regexp.MatchString(`\d{4}\-\d{2}\-\d{2}`, date)
	return match
}

func (v Validator) IsValidTime(time int) bool {
	if time < 0 || time > 23 {
		return false
	}
	return true
}

func (v Validator) IsValidGender(gender int) bool {
	enumGender := Gender(gender)
	isValidGender := enumGender == Male || enumGender == Female
	return isValidGender
}

func (v Validator) IsValidAge(age int) bool {
	if age < 1 || age > 130 {
		return false
	}
	return true
}
