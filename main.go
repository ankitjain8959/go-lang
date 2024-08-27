package main

import (
	"go-lang/basics"
	"log"
)

// Global variable in a package available to all functions in package but not outside (as it starts with lowercase letter)
var globalVariable = "global"

// Global variable in a package available to all functions in package and outside (as it starts with uppercase letter)
var GlobalVariable = "Global"

func understandVariables() {
	log.Println(globalVariable, "Variable from main package")
	log.Println(GlobalVariable, "Variable from main package")

	log.Println(basics.GlobalVariable, "Variable from basics package")

	var localVariable = "local" //local variable to function
	log.Println(localVariable, "Variable to function inside main package")

	// variable declaration
	var name string
	//variable assignment/initialization separately
	name = "Ankit"
	log.Println("My name is", name)

	//variable declared and assignment, merged in a single line with data type
	var id int64 = 10
	log.Println("My id is", id)

	//variable assigned/initialized without the data type. The variable will take the data type of the initializer value.
	var isSuperHuman = false // It is same as var isHuman bool = false
	log.Println("Am I a Super Human?:", isSuperHuman)
}

func main() {

	// calling a function
	understandVariables()

	//Inside a function, short assignment statement := can be used in place of var declaration with implicit type.
	//Outside a function, every statement begins with a keyword (var, func, and so on) and so the := construct is not available.
	//when using := it means returnedNameValue variable will be assigned a data type of function understandFunctionReturnSingleValue() returned value
	returnedNameValue := basics.UnderstandFunctionReturnSingleValue()
	log.Println("My name is", returnedNameValue)

	// accepting multiple values returned by a function
	returnedNameValueAgain, returnedId := basics.UnderstandFunctionReturnMultipleValues()
	log.Println("My name is", returnedNameValueAgain, "and my id is", returnedId)

	user := basics.User{
		FirstName: "Ankit",
		LastName:  "Jain",
	}
	log.Println("Type values are:", user.FirstName, user.LastName, user.PhoneNumber, user.Age, user.BirthDate)

	basics.UnderstandMap()

	basics.UnderstandSlices()

	basics.UnderstandDecisionStructures()

	basics.UnderstandLoops()
}
