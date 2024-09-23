package user

import (
	"errors"

	"fmt"
)

type User struct {
	firstName string
	lastName  string
	birthDate string
}

type Admin struct {
	email    string
	password string
	User
}

func NewAdmin(email, password string) *Admin {
	return &Admin{
		email:    email,
		password: password,
		User: User{
			firstName: "ADMIN",
			lastName:  "NAME",
			birthDate: "1/1/1997",
		},
	}
}

// Constructor function
func NewUser(firstName, lastName, birthdate string) (*User, error) {
	if firstName == "" || lastName == "" {
		return nil, errors.New("Empty firstname and lastname")
	}
	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthdate,
	}, nil

}

func (u *User) ClearUserName() {
	u.firstName = ""
}

func (p User) ReceiverPrintUserData() {
	fmt.Println(p.firstName, p.lastName, p.birthDate)
}

func PrintUserData(p *User) {
	fmt.Println(p.firstName, p.lastName, p.birthDate)
}
