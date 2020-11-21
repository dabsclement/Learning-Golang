package main

import "fmt"

func main() {
	name := "John"

	fmt.Println(name)
	fmt.Println(&name)
	// Declare the namePtr pointer of type string and assign the address of the variable name
	var namePtr *string = &name

	// Print the value of namePtr
	fmt.Println(namePtr)

}
// using pointer in function
package main

import "fmt"

func main() {
	name := "John"

	fmt.Println(name)
	// Call the changeName function and pass the address of the name variable
	changeName(&name)
	
	// Print the name variable
	fmt.Println(name)

}

// Add a parameter pointer of type *string
func changeName(namePtr *string) {
    // Dereference namePtr and update the value of the name variable to "Kate"
	*namePtr = "Kate"
	
}
