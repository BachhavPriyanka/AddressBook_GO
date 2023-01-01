package main

import "fmt"

type Contact struct {
	FirstName   string
	LastName    string
	Address     string
	City        string
	State       string
	PhoneNumber string
	Email       string
}

func (a Contact) Display() {
	fmt.Printf("\nFirst Name:%s\nLast Name:%s\nAddress:%s\nCity:%s\nState:%s\nPhone Number:%s\nEmail:%s", a.FirstName, a.LastName, a.Address, a.City, a.State, a.PhoneNumber, a.Email)
}

func main() {
	fmt.Print("********Address book main**********")

	contact1 := Contact{
		FirstName:   "Priyanka",
		LastName:    "Bachhav",
		Address:     "Township",
		City:        "Pune",
		State:       "MH",
		PhoneNumber: "9878564576",
		Email:       "bachhav@gmail.com"}

	contact1.Display()

}
