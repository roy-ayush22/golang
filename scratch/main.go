package main

import (
	"fmt"
)

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
		fmt.Print("enter command:")
		_, err := fmt.Scan(&input)
		if err != nil {
			fmt.Println("invalid error", err)
			continue
		}
		if input == 1 {
			// code to append the data into the in-memory contacts
			var name string
			var phone string
			var email string
			fmt.Println("enter contact details: [ex:name phone email]")
			_, err := fmt.Scan(&name, &phone, &email)
			if err != nil {
				fmt.Println("invalid input", err)
				continue
			}
			newContact := Contact{
				Name:  name,
				Phone: phone,
				Email: email,
			}
			contacts = append(contacts, newContact)
			fmt.Println("contact added successfully")

		} else if input == 2 {
			// code to display all contacts
			for i, contact := range contacts {
				fmt.Println(i, contact.Name, contact.Phone, contact.Email)
			}
		} else if input == 3 {
			// search contact by name
			var inputName string
			found := false
			fmt.Println("enter name to found: ")
			_, err := fmt.Scan(&inputName)
			if err != nil {
				fmt.Println("invalid input", err)
				continue
			}
			for _, contact := range contacts {
				if contact.Name == inputName {
					fmt.Println("match found:", contact)
					found = true
					break
				}
			}
			if !found {
				fmt.Println("contact not found")
			}

		} else if input == 4 {
			// code to delete contact
			var inputName string
			fmt.Println("enter contact to delete: ")
			found := false

			_, err := fmt.Scan(&inputName)
			if err != nil {
				fmt.Println("invalid input", err)
				continue
			}
			for i, contact := range contacts {
				if contact.Name == inputName {
					contacts = append(contacts[:i], contacts[i+1:]...)
					fmt.Println("contact deleted")
					found = true
					break
				}
			}
			if !found {
				fmt.Println("contact not found")
			}

		} else {
			// handle unvalid input
		}
	}
}
