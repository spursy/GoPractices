package main

import "fmt"
type AnimalCategory struct {
	kindom string
	phylum string
	class string
	order string
	family string
	genus string
	species string
}

type Animal struct {
	scientificName string
	AnimalCategory AnimalCategory
}

func (ac AnimalCategory) String()  string {
	return fmt.Sprintf("%s%s%s%s%s%s%s", ac.kindom, ac.phylum, ac.class, ac.order, ac.family, ac.genus, ac.species)
}


func (a Animal) String() string {
	return fmt.Sprintf("%s (category: %s)", a.scientificName, a.AnimalCategory)
}

func main() {
	category := AnimalCategory{species: "cat"} 
	fmt.Printf("The result is %s\n", category)

	animal := Animal {
		scientificName: "American Shorthair",
		AnimalCategory: category,
	}
	fmt.Printf("The animal: %s\n", animal)
}