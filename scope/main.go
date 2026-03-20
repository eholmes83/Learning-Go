package main

// variables challenge
	// declare package level variable for the main package
	// package named myVar

	// delcare a block variable for the main function
	// called blockVar

	// declare a package level variable in the packageone
	// package named PackageVar

	// create an exported function in the packageone called PrintMe

	// in the main function, print out the values of myVar, blockVar, and PackageVar on one line
	// using PrintMe function in packageone

import (
	"myapp/packageOne"
)

var myVar = "This is a package level variable!"

func main() {
	blockVar := "This is a block variable."
	packageone.PrintMe("myVar:" + myVar, " blockVar:" + blockVar)

}
