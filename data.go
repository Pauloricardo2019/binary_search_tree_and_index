package main

import (
	"github.com/jaswdr/faker"
)

type Person struct {
	ID    int
	Name  string
	Age   int
	Email string
	left  *node
	right *node
}

func GenerateObject(quantity int) ([]Person, error) {
	fakerGen := faker.New()
	people := make([]Person, 0)
	for i := 0; i < quantity; i++ {
		c := &Person{}
		c.ID = fakerGen.IntBetween(1, quantity)
		c.Name = fakerGen.Person().Name()
		c.Age = fakerGen.IntBetween(15, 78)
		c.Email = fakerGen.Internet().Email()
		people = append(people, *c)
	}

	return people, nil
}
