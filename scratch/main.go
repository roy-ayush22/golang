package main

import "fmt"

type Contact struct {
	Name  string
	Phone string
	Email string
}

var contacts []Contact

func main() {
	var input int
	fmt.Println("1. add contact")
	fmt.Println("2. list all contacts")
	fmt.Println("3. search by name")
	fmt.Println("4. delete contact")

	for {
		fmt.Println("enter command:")
		_, err := fmt.Scan(&input)
		if err != nil {
			fmt.Println("invalid error", err)
			continue
		}
		if input == 1 {
			// var addContact string
			var name string
			var phone string
			var email string
			fmt.Println("enter contact details:")
			_, err := fmt.Scan(&name, &phone, &email)
			if err != nil {
				fmt.Println("invlaid input", err)
				continue
			}
			newContact := Contact{
				Name:  name,
				Phone: phone,
				Email: email,
			}
			contacts = append(contacts, newContact)
			fmt.Println("contact added successfully")
			// code to append the data into the in-memory contacts
		} else if input == 2 {
			// code to display all contacts
		} else if input == 3 {
			// search contact by name
		} else if input == 4 {
			// code to delete contact
		} else {
			// handle unvalid input
		}
	}
}
