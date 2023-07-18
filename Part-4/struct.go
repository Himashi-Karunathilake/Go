package main

import "fmt"

type Person struct { // Define a struct
	Name string
	Age  int
}

type Employee struct { // Define another struct embedding the Person struct
	Person
	EmployeeID int
}

func (person Person) SayHello() { // Define a method associated with the Person struct
	fmt.Printf("Hi, %s!\n", person.Name)
}

func Struct() {
	fmt.Println("==================== PERSON STRUCT ====================")
	fmt.Println()

	// Create a new instance of the Person struct
	person := Person{
		Name: "Rachael",
		Age:  23,
	}

	// Access and modify the fields of the Person struct
	fmt.Println("Name of the individual: ", person.Name)
	fmt.Println("Age of the individual: ", person.Age)

	person.Age = 29
	fmt.Println("Updated age of the individual: ", person.Age)

	// Call the SayHello() method on the Person struct
	person.SayHello()

	fmt.Println()
	fmt.Println("==================== EMPLOYEE STRUCT ====================")
	fmt.Println()

	employee := Employee{
		Person: Person{
			Name: "Ross",
			Age:  35,
		},
		EmployeeID: 003452,
	}

	employee.SayHello()
	fmt.Println("Here are your details.")
	fmt.Println("Name: ", employee.Name)
	fmt.Println("Age: ", employee.Age)
	fmt.Println("Employee ID: ", employee.EmployeeID)

	fmt.Println()
	fmt.Println("==================== PETDOG STRUCT ====================")
	fmt.Println()

	petDog := struct { // Create an anonymous struct
		Name  string
		Breed string
		Age   int
	}{
		Name:  "Allie",
		Breed: "Labrador",
		Age:   2,
	}

	fmt.Println("Here are the details of your pet dog.")
	fmt.Println("Name of pet dog: ", petDog.Name)
	fmt.Println("Breed of pet dog: ", petDog.Breed)
	fmt.Println("Age of pet dog: ", petDog.Age)
}
