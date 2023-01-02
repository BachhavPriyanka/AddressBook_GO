package main

import (
	"fmt"
)

var contactStorage []Contact

func Operation() {
	var option int
	fmt.Println("\nSelect which operation do you wanna perform \n1.Add Contact\n2.Edit Contact\n3.Delete Contact\n4.Display Contact\n5.Exit")
	fmt.Scanln(&option)
	switch option {
	case 1:
		checkingDuplicateNames := filter(&contactStorage, duplicate)
		if len(checkingDuplicateNames) > 0 {
			fmt.Println("Duplicate entry press")
			Operation()
		} else {
			AddContact()
		}
	case 2:
		EditContacts(&contactStorage)
	case 3:
		DeleteContact(&contactStorage)
	case 4:
		DisplayContacts(&contactStorage)
	case 5:
		break

	}

}

func duplicate(name string) bool {
	var fName string
	fmt.Println("Enter Name : ")
	fmt.Scanln(&fName)
	if name == fName {
		return true
	}
	return false
}

func filter(person *[]Contact, duplicate func(string) bool) []Contact {
	duplicateContact := []Contact{}
	for _, val := range *person {
		singleFirstName := val.FirstName
		if duplicate(singleFirstName) {
			duplicateContact = append(duplicateContact, val)

		}
	}
	return duplicateContact
}

func DeleteContact(persons *[]Contact) {
	//Ability to delete person by using person name

	var nameToDelete string
	fmt.Println("Enter a name which you want to Delete :")
	fmt.Scanln(&nameToDelete)
	for i, person := range *persons {
		if person.FirstName == nameToDelete {
			*persons = append((*persons)[:i], (*persons)[i+1:]...)
		}
	}

	Operation()
}

func EditContacts(persons *[]Contact) {

	fmt.Println("Data in  edit func", *persons)

	var nameToChange, newAddress, newCity, newState, newPhoneNum, newEmail string
	var editField int
	fmt.Println("Enter a name which you want to edit :")
	fmt.Scanln(&nameToChange)
	for i, person := range *persons {
		if person.FirstName == nameToChange {
			fmt.Println("Press : \n1.Address\n2.City \n3.State \n4.Phone Number \n5.Email")
			fmt.Println("Enter what you want to edit : ")
			fmt.Scanln(&editField)
			switch editField {
			case 1:
				fmt.Println("Enter Address :")
				fmt.Scanln(&newAddress)
				(*persons)[i].Address = newAddress
			case 2:
				fmt.Println("Enter City :")
				fmt.Scanln(&newCity)
				(*persons)[i].City = newCity
			case 3:
				fmt.Println("Enter State :")
				fmt.Scanln(&newState)
				(*persons)[i].State = newState
			case 4:
				fmt.Println("Enter Phone Number :")
				fmt.Scanln(&newPhoneNum)
				(*persons)[i].PhoneNumber = newPhoneNum
			case 5:
				fmt.Println("Enter Email :")
				fmt.Scanln(&newEmail)
				(*persons)[i].Email = newEmail
			}

		}
	}
	Operation()
}

func DisplayContacts(persons *[]Contact) {
	fmt.Println("********Displying Contact Details*********")
	for _, p := range *persons {
		fmt.Println("Name :", p.FirstName)
		fmt.Println("Last Name :", p.LastName)
		fmt.Println("  Address :", p.Address)
		fmt.Println("  City :", p.City)
		fmt.Println("  State :", p.State)
		fmt.Println("  Phone Number:", p.PhoneNumber)
		fmt.Println("  Email:", p.Email)
		fmt.Println("-----------------------------------")

	}
	Operation()
}

func AddContact() {

	var p Contact

	fmt.Print("\nEnter Name: ")
	_, err := fmt.Scanf("%s", &p.FirstName)
	Error(err, "Name")

	fmt.Print("Enter Last Name: ")
	_, err = fmt.Scanf("%s", &p.LastName)
	Error(err, "LastName")

	fmt.Print("Enter address: ")
	_, err = fmt.Scanf("%s", &p.Address)
	Error(err, "Address")

	fmt.Print("Enter City: ")
	_, err = fmt.Scanf("%s", &p.City)
	Error(err, "City")

	fmt.Print("Enter State: ")
	_, err = fmt.Scanf("%s", &p.State)
	Error(err, "State")

	fmt.Print("Enter phone number: ")
	_, err = fmt.Scanf("%s", &p.PhoneNumber)
	Error(err, "Phone number")

	fmt.Print("Enter email: ")
	_, err = fmt.Scanf("%s", &p.Email)
	Error(err, "Email")

	contactStorage = append(contactStorage, p)
	Operation()

}

func Error(err error, value string) {
	if err != nil {
		fmt.Println("Error reading", value, ":")
	}
}
