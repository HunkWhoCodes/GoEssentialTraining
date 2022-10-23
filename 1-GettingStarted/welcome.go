// Package
// All GO files belong to a package.
// All files in the same fodler on the disc should be in the same package
// Package main  is a special package that makes the coed compile to an executable and not to a library

// Import Statements:
// fmt is a package which outputs formatted output

// func keyword
// All functions start with func keyword and then parameters if any
// The function body starts with curly braces.
// All the function calls are proceeded with the package name like fmt.println()

// Every function or symbol, being used from another package is going to start with a capital letter like Println()

// Strings in GO are UTF 8

// Most GO code does not have smicolon

package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World!")
}
