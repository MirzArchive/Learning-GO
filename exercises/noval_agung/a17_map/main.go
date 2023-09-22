package main

/*
Map is an associative data type in Go using key-value pair.
The concept is simillar to PHP Associative array.
*/

func main() {
	// Ways to declare a map of integer with string as the key-value pair
	var varMap1 map[string]int
	varMap2 := map[string]int{}
	dump(varMap1, varMap2)

	// Map default value is nil, and its element default value is the
	// correlated zero value of that type (ie. integer then its 0)
	if varMap1 == nil {
		println("varMap1 is NIL\n")
	}

	// Because map default value is nil, then it must be initialized
	// first before any element appending.
	// ie.
	//
	// var varMap map[string]int
	// varMap["one"] = 1
	// will output error!

	// Ways to declare and also initialize map (so its not NIL)
	// Note: it doesnt have to use var, operator := is okay
	var varMap3 = map[string]int{}
	var varMap4 = make(map[string]int)
	var varMap5 = *new(map[string]int)
	dump(varMap3, varMap4, varMap5)

	// Iterate map with for-range
	// Instead of index in the left side of for-range, its the map key
	varMap := map[string]int{"one": 1, "two": 2, "three": 3}
	for key, value := range varMap {
		println(key, "\t=>", value)
	}
	println()

	// Delete an item from map with delete() by specifying
	// whats the map and which key
	println("Length before:", len(varMap))
	delete(varMap, "two")
	println("Length after:", len(varMap), "\n")

	// Check item existence in a map by specifying second variable
	// as the boolean variable whether a item with specific key exists or not
	if _, exists := varMap["two"]; exists {
		println("It exists!\n")
	} else {
		println("It dont exists!\n")
	}

	// Combination of Slice & Map
	// The concept is very simillar to array of object in JSON
	sliceOfMap := []map[string]int{
		{"one": 1, "two": 2},
		{"three": 3, "four": 4},
		{"five": 5, "six": 6},
	}

	// Iterate slice of map
	for _, imMap := range sliceOfMap {
		for key, value := range imMap {
			println(key, "\t=>", value)
		}
	}
}

func dump(vars ...interface{}) {}
