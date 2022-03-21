package main

import "fmt"

type User struct {
	FirstName        string
	LastName         string
	Age              int
	FullName         string
	AgeAfterTenYears int
	WhenUserGetsixty int
}

func main() {
	u := c()
	user := u(User{
		FirstName: "Mohamad Yehya",
		LastName:  "Shahbour",
		Age:       25,
	}).WhenUserAgeBecomesixty().UserAgeAfterTenYears().GenerateFullName()

	fmt.Println("fullName: ", user.FullName)
	fmt.Println("User age after 10 years: ", user.AgeAfterTenYears)
	fmt.Println("User get 60 years old in: ", user.WhenUserGetsixty)
}

func c() func(user User) User {
	return func(user User) User {
		return user
	}
}

func (user User) GenerateFullName() User {
	user.FullName = user.FirstName + " " + user.LastName
	return user
}

func (user User) UserAgeAfterTenYears() User {
	user.AgeAfterTenYears = user.Age + 10
	return user
}

func (user User) WhenUserAgeBecomesixty() User {
	user.WhenUserGetsixty = 60 - user.Age
	return user
}