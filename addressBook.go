package main

import "fmt"

var contactStorage []Contact

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
	fmt.Println(persons)
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
	fmt.Println(persons)
	DeleteContact(persons)

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
	// fmt.Println("What is getting printed", *persons)
	EditContacts(persons)
}

func AddContact() {
	for {
		var p Contact

		fmt.Print("\nEnter Name: ")
		_, err := fmt.Scanf("%s", &p.FirstName)
		if err != nil {
			fmt.Println("Error reading name:", err)
			continue
		}

		fmt.Print("Enter Last Name: ")
		_, err = fmt.Scanf("%s", &p.LastName)
		if err != nil {
			fmt.Println("Error reading last name:", err)
			continue
		}

		fmt.Print("Enter address: ")
		_, err = fmt.Scanf("%s", &p.Address)
		if err != nil {
			fmt.Println("Error reading address:", err)
			continue
		}
		fmt.Print("Enter City: ")
		_, err = fmt.Scanf("%s", &p.City)
		if err != nil {
			fmt.Println("Error reading email:", err)
			continue
		}
		fmt.Print("Enter State: ")
		_, err = fmt.Scanf("%s", &p.State)
		if err != nil {
			fmt.Println("Error reading email:", err)
			continue
		}
		fmt.Print("Enter phone number: ")
		_, err = fmt.Scanf("%s", &p.PhoneNumber)
		if err != nil {
			fmt.Println("Error reading phone number:", err)
			continue
		}
		fmt.Print("Enter email: ")
		_, err = fmt.Scanf("%s", &p.Email)
		if err != nil {
			fmt.Println("Error reading email:", err)
			continue
		}

		contactStorage = append(contactStorage, p)

		fmt.Print("Add another person? (y/n) ")
		var answer string
		_, err = fmt.Scanf("%s", &answer)
		if err != nil {
			fmt.Println("Error reading answer:", err)
			continue
		}
		if answer != "y" {
			break
		}
	}
	DisplayContacts(&contactStorage)
}
