package main

type Dog struct {
	name string
}

func New(name string) Dog {
	return Dog{name}
}

func (dog *Dog) SetName(name string) {
	dog.name = name
}

func (dog Dog) Name() string {
	return dog.name
}

func main() {
	// demo 1
	// New("little pig").SetName("monster")

	// demo 2
	map[string]int {"the": 0, "world": 0, "counter": 0}["world"]++
	map1 := map[string]int{"the": 0, "world": 0, "counter": 0}
	map1["world"]++
}

