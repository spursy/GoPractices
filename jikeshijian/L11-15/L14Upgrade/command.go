package main

import "fmt"

type Pet interface {
	Name() string
	Category() string
}

type Dog struct{
	name string
}

func (dog *Dog) SetName(name string) {
	dog.name = name
}

func (dog Dog) Name() string {
	return dog.name
}

func (dog Dog) Category() string {
	return "dog"
}

func main() {
	// demo 1
	// dog := Dog{"little pig"}
	// fmt.Printf("The dog's name is %q.\n", dog.name)
	// var pet Pet = dog
	// dog.SetName("monster")
	// fmt.Printf("The dog's name is %q.\n", dog.name)
	// fmt.Printf("This pet is a %s, the name is %q.\n", pet.Category(), pet.Name())
	// fmt.Println()

	// demo 2
	// dog1 := Dog{"little pig"}
	// fmt.Printf("The name of firs dog is %q.\n", dog1.name)
	// dog2 := dog1
	// fmt.Printf("The name od second dog is %q.\n", dog1.name)
	// dog1.name = "monster"
	// fmt .Printf("The name of the first dog id %q.\n", dog1.name)
	// fmt .Printf("The name of the second dog id %q.\n", dog2.name)

	// demo 3
	dog := Dog{"little pig"}
	fmt.Printf("The dog's name is %q\n", dog.Name())
	pet := &dog
	dog.SetName("monster")
	fmt.Printf("The dog's name si %q.\n", dog.Name())
	fmt.Printf("This pet is a %s, the name is %q.\n", pet.Category(), pet.Name())
}