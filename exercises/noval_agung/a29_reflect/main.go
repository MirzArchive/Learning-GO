package main

import "reflect"

/*
Reflection is a technique used to inspect, retrieve
information or even manipulate variable itself
*/

func main() {
	// Basic usage of reflect to check value of itself, and
	// check what its type
	var number = 13
	var reflectVar = reflect.ValueOf(number)
	println("Data type:", reflectVar.Type().String())
	if reflectVar.Kind() == reflect.Int {
		println("Value:", reflectVar.Int())
	}

	// With reflect its possible to check the data type
	// of a value which is assigned to an interface{} variable
	var mysterious interface{}
	mysterious = "123"
	reflectVar = reflect.ValueOf(mysterious)
	println("Data type:", reflectVar.Type().String())
	mysterious = 123.1
	reflectVar = reflect.ValueOf(mysterious)
	println("Data type:", reflectVar.Type().String())
}
