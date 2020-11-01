package main

import (
	"encoding/json"
	"fmt"

	"github.com/gotidy/copy"
)

type Person struct {
	Name       string
	MiddleName *string
	Surname    string
}

type User struct {
	Person
	Email   string
	Age     int8
	Married bool
}

type Employee struct {
	Name       string
	MiddleName string
	Surname    string
	Email      string
	Age        int
}

func main() {
	src := User{
		Person: Person{
			Name:       "John",
			MiddleName: nil,
			Surname:    "Smith",
		},
		Email:   "john.smith@joy.me",
		Age:     33,
		Married: false,
	}
	dst := Employee{}

	copiers := copy.New() // New("json")
	copier := copiers.Get(&Employee{}, &User{})
	copier.Copy(&dst, &src)
	if data, err := json.MarshalIndent(dst, "", "    "); err == nil {
		fmt.Println(string(data))
	}
}
