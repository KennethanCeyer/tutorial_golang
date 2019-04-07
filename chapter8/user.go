package main

type Gender int

const (
	DefaultUserName = "새로운 유저"
)

const (
	Male   Gender = 1
	Female Gender = 2
)

type User struct {
	Name   string
	Gender Gender
	Age    int
}
