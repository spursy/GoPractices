package main

func main() {
	// demo 1
	// var badMap1 = map[[]int]int{}

	var badMap2 = map[interface{}] int {
		"1": 1,
		[]int{2}: 2,
		// 3: 3
	}
	_ = badMap2;
}