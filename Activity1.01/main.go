package main

import (
	"fmt"
)

/*
type FormMedical struct {
	FirstName     string
	FamilyName    string
	Age           int
	PeanutAllergy bool
}*/

func main() {
	/*var form FormMedical
	fmt.Print("First Name: ")
	fmt.Scanln(&form.FirstName)
	fmt.Print("Family Name: ")
	fmt.Scanln(&form.FamilyName)
	fmt.Print("Age: ")
	fmt.Scanln(&form.Age)
	fmt.Print("Peanut Allergy (true/false): ")
	fmt.Scanln(&form.PeanutAllergy)


	fmt.Printf("\nForm Medical:\nFirst Name: %s\nFamily Name: %s\nAge: %d\nPeanut Allergy: %t\n",
		form.FirstName, form.FamilyName, form.Age, form.PeanutAllergy)*/

	var name string = "Cristian"
	var familyName string = "Quezada"
	var age int = 28
	var peanutAllergy bool = false

	fmt.Printf("Form Medical:\nFirst Name: %s\nFamily Name: %s\nAge: %d\nPeanut Allergy: %t\n",
		name, familyName, age, peanutAllergy)

}
