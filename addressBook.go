package main

import (
	"fmt"
	"strings"
)

var contactStorage []Contact

func Operation() {
	var option int
	fmt.Println("\nSelect which operation do you wanna perform \n1.Add Contact\n2.Edit Contact\n3.Delete Contact\n4.Display Contact\n5.Search By city and state \n6.Search By City and State in Map Format \n7.Exit")
	fmt.Scanln(&option)
	switch option {
	case 1:
		var fName string
		fmt.Println("Enter Name : ")
		fmt.Scanln(&fName)
		AddContact(fName)
	case 2:
		EditContacts(&contactStorage)
	case 3:
		DeleteContact(&contactStorage)
	case 4:
		DisplayContacts(&contactStorage)
	case 5:
		CityStatePersonSlice := Search(&contactStorage)
		fmt.Println(CityStatePersonSlice)
		Operation()
	case 6:
		SearchStateOrCity(&contactStorage)
	case 7:
		CityStatePersonCount := SearchCount(&contactStorage)
		fmt.Println("Person from same City/State : ", CityStatePersonCount)
		Operation()
	case 8:
		break
	}

}

func SearchCount(people *[]Contact) int {
	var cityState string
	fmt.Println("Enter City or State")
	fmt.Scanln(&cityState)
	count := 0

	CityStateString := strings.TrimSpace(cityState)
	for _, person := range *people {
		if strings.Contains(person.City, CityStateString) || strings.Contains(person.State, CityStateString) {
			count++
		}
	}
	return count
}
func CityFilter(ContactList *[]Contact, findingCity func(string) bool) []Contact {
	CitySlice := []Contact{}
	for _, val := range *ContactList {
		if findingCity(val.City) {
			CitySlice = append(CitySlice, val)
		}
	}
	return CitySlice
}
func StateFilter(ContactList *[]Contact, findingState func(string) bool) []Contact {
	StateSlice := []Contact{}
	for _, val := range *ContactList {
		if findingState(val.State) {
			StateSlice = append(StateSlice, val)
		}
	}
	return StateSlice
}

func SearchStateOrCity(people *[]Contact) {
	var cityState string
	CityMap := make(map[string]string, 0)
	stateMap := make(map[string]string, 0)

	fmt.Println("Enter City or State to find How many Persons belongs to it")
	fmt.Scanln(&cityState)

	ResultOfSameCity := CityFilter(people, func(cityName string) bool {
		if cityState == cityName {
			return true
		} else {
			return false
		}
	})

	ResultOfSameState := StateFilter(people, func(stateName string) bool {
		if cityState == stateName {
			return true
		} else {
			return false
		}
	})

	//Printing person names with same state in Map
	for _, val := range ResultOfSameState {
		stateMap[val.State] = val.FirstName
		fmt.Println(stateMap)
	}

	//Printing person names with same city in Map
	for _, val := range ResultOfSameCity {
		CityMap[val.City] = val.FirstName
		fmt.Println(CityMap)
	}
	Operation()
}

func Search(people *[]Contact) []Contact {
	var cityState string
	fmt.Println("Enter City or State")
	fmt.Scanln(&cityState)

	storingCityState := []Contact{}
	CityStateString := strings.TrimSpace(cityState)
	for _, person := range *people {
		if strings.Contains(person.City, CityStateString) || strings.Contains(person.State, CityStateString) {

			storingCityState = append(storingCityState, person)
		}
	}
	return storingCityState
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

func AddContact(fName string) {

	var p Contact

	checkingDuplicateNames := filter(&contactStorage, func(name string) bool {
		if name == fName {
			return true
		}
		return false
	})

	if len(checkingDuplicateNames) > 0 {
		fmt.Println("Duplicate entry press")
		Operation()
	} else {
		p.FirstName = fName

		fmt.Print("Enter Last Name: ")
		_, err := fmt.Scanf("%s", &p.LastName)
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

}

func Error(err error, value string) {
	if err != nil {
		fmt.Println("Error reading", value, ":")
	}
}
