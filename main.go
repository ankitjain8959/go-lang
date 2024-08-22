package main

import (
	"log"
	"sort"
	"time"
)

var globalVariable = "Global" //Global variable in a package available to all functions in package but not outside (as it starts with lowercase letter)
var GlobalVariable = "Global" //Global variable in a package available to all functions in package and outside (as it starts with uppercase letter)

func main() {
	log.Println("Global variable", globalVariable)
	var localVariable = "Local" //local variable to function
	log.Println("Local variable", localVariable)

	// calling a function
	understandVariables()

	//Inside a function, short assignment statement := can be used in place of var declaration with implicit type.
	//Outside a function, every statement begins with a keyword (var, func, and so on) and so the := construct is not available.
	//when using := it means returnedNameValue variable will be assigned a data type of function understandFunctionReturnSingleValue() returned value
	returnedNameValue := understandFunctionReturnSingleValue()
	log.Println("My name is", returnedNameValue)

	// accepting multiple values returned by a function
	returnedNameValueAgain, returnedId := understandFunctionReturnMultipleValues()
	log.Println("My name is", returnedNameValueAgain, "and my id is", returnedId)

	user := User{
		FirstName: "Ankit",
		LastName:  "Jain",
	}
	log.Println("Type values are:", user.FirstName, user.LastName, user.PhoneNumber, user.Age, user.BirthDate)

	understandMap()

	understandSlices()

	understandDecisionStructures()

	understandLoops()

}

func understandVariables() {
	log.Println("global Variable:", globalVariable)
	log.Println("Global Variable:", GlobalVariable)

	// variable declaration
	var name string
	//variable initialization separately
	name = "Ankit"
	log.Println("My name is", name)

	//variable declared and initialized in a single line with data type
	var id int64 = 10
	log.Println("My id is", id)

	//variable defined/initialized without the data type. The variable will take the data type of the initializer value.
	var isSuperHuman = false // It is same as var isHuman bool = false
	log.Println("Am I a Super Human?:", isSuperHuman)
}

// function returning a single value - string
func understandFunctionReturnSingleValue() string {
	return "Ankit"
}

// function returning multiple values - string and int
func understandFunctionReturnMultipleValues() (string, int) {
	return "Ankit", 10
}

// Function or Global Variable or type variable if defined with a Capital letter: They will be available/accessible from outside the package else they are available only inside package
func visibleInsidePackage() {
	//This function is available only within this package since it starts with lowercase v
	// It can be also called an "unexported" method since it's name does not begins with a capital letter
}

func VisibleOutsidePackage() {
	//This function is available outside this package since it starts with uppercase V
	// It can be also called an "exported" method since it's name begins with a capital letter
}

// struct is a collection of fields and these fields are accessed using a dot
// Structure for User type
type User struct {
	FirstName   string
	LastName    string
	PhoneNumber string
	Age         int
	BirthDate   time.Time
	//Variables inside the type if starts with uppercase letter, are visible and accessible outside the package
}

func understandMap() {
	// Empty map or nil map.
	// Zero value of map is nil.
	// A nil map doesn't contain any key. If you try to add a key-value pair in the nil map, then compiler will throw runtime error
	var nilMap map[int]string
	if nilMap == nil {
		log.Println("Empty Map")
	}

	// Creating and initializing a map using a short hand declaration
	firstNamesMap := map[int]string{
		1: "Ankit",
		2: "Dimple",
	}
	log.Println("FirstNames Map values are:", firstNamesMap[1], "-", firstNamesMap[2])

	// Creating a map using make() function. make() function always returns a map which is initialized
	lastNamesMap := make(map[int]string)
	lastNamesMap[1] = "Jain"
	lastNamesMap[2] = "Mehta"

	log.Println("LastNames Map values are:", lastNamesMap[1], "-", lastNamesMap[2])

	customMap := make(map[int]User)
	myUser := User{
		FirstName: "Ankit",
		LastName:  "Jain",
	}
	customMap[1] = myUser
	log.Println("My Custom Map values are", customMap[1])
	log.Println("My Custom Map FirstName value is", customMap[1].FirstName)
}

func understandSlices() {
	//slice (like, arrays) of integers
	var numbers []int
	numbers = append(numbers, 3)
	numbers = append(numbers, 1)
	numbers = append(numbers, 2)
	log.Println("Slice values are", numbers)

	sort.Slice(numbers, func(i, j int) bool {
		return numbers[i] < numbers[j]
	})

	log.Println("Slice values after sorting in ascending order are", numbers)

	// Zero value of slice is nil.
	// a nil slice has a length and capacity of 0
	var nilSlice []int
	log.Println("Nil Slice is ", nilSlice) //Output: "Nil Slice is []"

	//slice of integers using short hand expression
	newNumbers := []int{6, 4, 7, 5}
	log.Println(newNumbers)

	//Range of values from the slice - newNumber[x:y] where x is inclusive and y is exclusive
	log.Println(newNumbers[0:2])
	log.Println(newNumbers[1:3])
}

func understandDecisionStructures() {
	// if else statement
	isHuman := false
	if isHuman {
		log.Println("I am a Human")
	} else {
		log.Println("I am not a Human")
	}

	//switch statement
	name := "Ankit"
	switch name {
	case "Ankit":
		log.Println("My Name is Ankit")
	case "Dimple":
		log.Println("My Name is Dimple")
	default:
		log.Println("My Name is something else")
	}
}

// Go language contains only one loop i.e. for-loop. However, the for loop can be used in different form
func understandLoops() {
	// 1. simple for loop with it's three component: the init statement, the condition expression and the post statement
	for index := 0; index < 10; index++ {
		log.Println("simple for loop with it's three component", index)
	}

	animals := []string{"Dog", "Cat"}
	for index := 0; index < len(animals); index++ {
		log.Println("simple for loop with it's three component", animals[index])
	}

	// init statement and post statement are optional
	// 2. In Go, for loop can work as a while loop
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	log.Println("for loop works as while loop", sum)

	// 3.1. using range in for loop
	for index, animal := range animals {
		log.Println("using range in for loops", index, animal)
	}

	// 3.2. using range in for loops and using underscore when you don't care about the index/current iteration
	for _, animal := range animals {
		log.Println("using range in for loops with underscore in place of index", animal)
	}

	//Note: we can range over custom objects, slices, string, maps

	// 4. for loop can iterate over the key and value pairs of the map
	myMap := map[int]string{
		1: "Ankit",
		2: "Dimple",
	}
	for key, value := range myMap {
		log.Println("Iterate map using for loop:", key, value)
	}
	// 5. infinite for loop when all three component are missing: the init statement, the condition expression and the post statement
	/*
		for {}
	*/
}
