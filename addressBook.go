package main

import "fmt"

var contactStorage []Contact

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
