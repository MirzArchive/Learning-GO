package main

/*
Pointer is a reference or memory address. A Pointer Variable
is a variable which contain a memory address.
*/

func main() {
	// Ways to make a pointer variable
	var number1 *int
	number2 := new(int)
	dump(number1, number2)

	// Example of pointer reference meaning
	numberA := 1
	numberB := &numberA
	println("BEFORE:")
	println("number A:", numberA)
	println("number B:", *numberB)
	numberA = 2
	println("AFTER:")
	println("number A:", numberA)
	println("number B:", *numberB)
	println()

	// A non-pointer variable pointer could be taken with
	// ampersand (&) symbol
	var number int
	numberPointer := &number
	dump(numberPointer)

	// There's also the opposite, that is dereferencing
	// which take the actual value from that pointer
	// which in this case are an memory addresss
	number = 3
	println("DEREFERENCE:", *numberPointer)

	// It also possible to have nested pointer. Its
	// weird but interesting
	var quirky **************************int
	dump(quirky)
}

func dump(vars ...interface{}) {}
