package main

import ("fmt")

type Animal interface {
	ScientificName() string
	Category() string
}

type Named interface {
	Name() string
}

type Pet interface {
	Animal
	Named
}

type PetTag struct {
	name string
	owner string
}

func (pt PetTag) Name() string{
	return pt.name
}

func (pt PetTag) Owner() string {
	return pt.owner
}

type Dog struct {
	PetTag
	scirntificName string
}

func (dog Dog) Category() string {
	return "dog"
}

func main() {
	petTag := PetTag{name: "little pig"}
	_, ok := interface{}(petTag).(Named)
	fmt.Printf("PetTag implements interface named: %v\n", ok)
	
	dog := Dog {
		PetTag: petTag,
		scirntificName: "Labrador Retriever",
	}
	_, ok = interface{}(dog).(Animal)
	fmt.Printf("Dog implements interface Animal: %v\n", ok)

	_, ok = interface{}(dog).(Named)
	fmt.Printf("Dog implements interface Named: %v\n", ok)
	_, ok = interface{}(dog).(Pet)
	fmt.Printf("Dog implements interface Pet: %v\n", ok)
}