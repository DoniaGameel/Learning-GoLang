package main

import (
	"fmt"
	"time"
)

// defining the struct
type user struct{
	firstName string
	lastName string
	birthDate string
	createdAt time.Time
}

// method of the struct
func (u user) outputUserDetails(){
	fmt.Println(u.firstName, u.lastName, u.birthDate)
}

// method of the struct
func (u *user) clearUserName(){
	u.firstName = ""
	u.lastName = ""
}

// creation/constructor function of the struct
// this name is convention. not a must but it's a good practice
// constructor isn't a built-in feature in Go. It is just a convention
// normal function but we deal it as the constructor
func newUser(userFirstName,userLastName,userBirthdate string) user{
	return user{
		firstName: userFirstName,
		lastName: userLastName,
		birthDate: userBirthdate,
		createdAt: time.Now(),
	}
}

// struct embeddings
type Admin struct {
	email string
	password string
	// it is caled anynomous embeddings
	// using it we canm directly call the fields of User struct using an instance of Admin struct
	user // it can also be u user (can be given a name)
}

// constructor of Admin
func newAdmin(email , password string){
	email: email,
	password: password,
	User: User{
		firstName: "ADMIN",
		lastName: "ADMIN",
		birthDate: "--------",
		createdAt: time.Now(),
	},
}

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	// instantiate an instance of the user struct type
	var appUser user 
	appUser = newUser(userFirstName,userLastName,userBirthdate)

	/*
	can also be defined as so:
	appUser = user{
		userFirstName,
		userLastName,
		userBirthdate,
		time.Now(),
	}
	*/
	// ... do something awesome with that gathered data!

	fmt.Println(userFirstName, userLastName, userBirthdate)
}

func outputUserDetails(u user){
	fmt.Println(u.firstName, u.lastName, u.birthDate)
}

func outputUserDetailsPointer(u *user){
	fmt.Println(u.firstName, u.lastName, u.birthDate)
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
