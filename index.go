// => Basic Structure

// package main

// import "fmt"

// func main(){
// 	fmt.Println("Hello, from Go");
// }

// -------------------------------------
// => Variables : In Go, there are two ways to declare a variable. It has case sensitive variable name (age & AGE has different values)

// -> Using "var" keyword

// package main

// import (
// 	"fmt"
// )

// var name ="John Doe"

// func main(){

// var fruit = "Apple";
// fmt.Println(name)
// fmt.Println(fruit)
// }

// --------------------------------------

// -> Using ":=" sign followed by the variable value

// package main

// import (
// 	"fmt"
// )

// func main(){
// variableOne := 100;
// variableTwo := 8;
// fmt.Println(variableOne)
// fmt.Println(variableTwo)
// }

// --------------------------------------

// -> Assigning Types + Types Inferred (It means we dont have to explicitly define the type of the variable)

// package main

// import (
// 	"fmt"
// )

// func main(){

// 	var fruit string = "Mango" // Type is string
// 	var user = "Sdey" // Type is inferred
// 	number := 20 // Type is inferred

// 	fmt.Println(fruit)
// 	fmt.Println(user)
// 	fmt.Println(number)
// }

// --------------------------------------

// -> Go Multiple Variable Declaration

// package main

// import (
// 	"fmt"
// )

// func main(){

// 	var one, two, three int = 1,2,3

// 	fmt.Println(one)
// 	fmt.Println(two)
// 	fmt.Println(three)
// }

// --------------------------------------

// -> Go Variable Declaration In A Block (Rarely used)

// package main

// import (
// 	"fmt"
// )

// func main(){

// 	var (
// 		num int // default value set to 0
// 		number int = 1
// 		greet string = "hello"
// 	)

// 	fmt.Println(num)
// 	fmt.Println(number)
// 	fmt.Println(greet)
// }

// ______________________________________

// => "const" Keyword : The const keyword declares the variable as "constant", which means that it it unchangeable and read-only. Constant names follow the same naming rules as variables, constant names are usually written in uppercase letters, constants can be declared both inside and outside of a function.

// package main

// import (
// 	"fmt"
// )

// const user = "admin" // cannot be changed

// func main(){

// 	fmt.Println(user)

// Boolean data type
// 	isGOlangPL := true
// 	isHTMLPL := false

// 	fmt.Println(isGOlangPL)
// 	fmt.Println(isHTMLPL)

// }

// ______________________________________

// => Strings : Must be quoted with double quotes or back ticks

// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func main(){

// ->> Multi-line string using back ticks ``
// name:= `Hello
// hey
// hi`
// fmt.Println(name)

// ->> String Concatenation
// greetings := "Hello " + "Good Morning"
// fmt.Println(greetings)

// ->> Access a specific character from the string

// message := "Hello"
// fmt.Println(message[0]) // prints the ASCII/Unicode of the first character of the string

/// To print the character, we need to convert it to a string by encoding it to uint8 to UTF-8, using printf because Println will print the ASCII/Unicode value, but Printf will print the character associated with the ASCII/Unicode value
// fmt.Printf("%c", message[0]) // Prints the first character of the string

// ->> String Length
// fruit := "Mango"
// fmt.Println(len(fruit)) // Prints the length of the string

// ->> String Standard Library (Read docs at https://pkg.go.dev/std)

// strings.Compare(str1, str2) Method : It returns an integer comparing two strings lexicographically. The result will be 0 if a == b, -1 if a < b, and +1 if a > b. (a,b = Sum of the each character's ASCII code of string)

// var1 := "one"
// var2 := "One"

// println(strings.Compare(var1,var2))

// strings.Contains(str, substr) Method : Checks whether str contains the substr.

// findChar := "Golang Programmin

// strings.ToLower(str1) & strings.ToUpper(str2) Methods : Converts the string str1 & str2 to lowercase & uppercase respectively.

// strOne := "HELLO"
// fmt.Println(strings.ToLower(strOne))

// strTwo := "welcome"
// fmt.Println(strings.ToUpper(strTwo))
// }

// ______________________________________

// => Go Condiitional Statements

// Syntax:
// if condition { <code> }
// else if condition { <code> }
// else { <code> }

// package main

// import (
// 	"fmt"
// )

// func main(){
// 	password := "Hello, Buddies"

// 	if len(password)>=8 {
// 		fmt.Println("Password setup complete :)")
// 	} else if len(password)<8 {
// 		fmt.Println("Password is short :(")
// 	} else {
// 		fmt.Println("Invalid Password!")
// 	}
// }

// ______________________________________

// => Loops

// Syntax :
// for initialExpression; condition; increment { <code> }

// package main

// import "fmt"

// func main(){

// ->> For Loop :

// >-> Single For Loop

// for i:=1;i<=10;i++ {

/// Continue Statement is used to skip one or more iterations in the loop

// 	if i==3 {
// 		continue
// 	}

// 	fmt.Println(i)

/// Break Statement is used to terminate the loop execution

// 	if i==8 {
// 		break
// 	}
// }

// >-> Nested For Loop

// for i:=0; i<5;i++ {
// 	fmt.Println("Outer Loop",i)
// 	for j:=0; j<3;j++ {
// 		fmt.Println("Inner Loop",j)
// 	}
// }

// ->> While Loop :

// 	num := 1

// 	for num <= 10000000 {
// 		fmt.Println(num)
// 		num++
// 	}
// }

// ______________________________________

// => Switch Statement

// package main

// import (
// 	"fmt"
// )

// func main(){
// 	// e.g - 1

// 	day := 4
// 	switch day {
// 	case 1 :
// 		fmt.Println("Mon")
// 	case 2 :
// 		fmt.Println("Tue")
// 	case 3 :
// 		fmt.Println("Wed")
// 	case 4 :
// 		fmt.Println("Thur")
// 	case 5 :
// 		fmt.Println("Fri")
// 	case 6 :
// 		fmt.Println("Sat")
// 	case 7 :
// 		fmt.Println("Sun")
// 	default:
// 		fmt.Println("Invalid day :(")
// 	}

// 	// e.g - 2

// 	num := 10
// 	switch {
// 	case num > 0:
// 		fmt.Println("Number is positive")
// 	case num < 0:
// 		fmt.Println("Number is negative")
// 	default:
// 		fmt.Println("INumber is 0")
// 	}
// }

// ______________________________________

// => Arrays : Store same type of multiple values in a single variable. Zero (0) index based.

// Syntax :
// var arr = [size]data_type{elements_of_array}

// package main

// import (
// 	"fmt"
// )

// func main(){

// var arr1 = [5]int{}
// var arr2=[5]int{1,2,3,4,5}
// arr3 := [3]int{5,9,8}

// fmt.Println(arr1) // It returns [0 0 0 0 0] as default value
// fmt.Println(arr2)
// fmt.Println(arr3)

/// Get & set the value in array
// arr1[0] = 10
// arr1[1] = 30
// arr1[2] = 20

// fmt.Println("After Updation In arr1", arr1) // Other two indexes remain empty with the value 0

// str_arr := [5]string{"Aizen","Zoro","Goku"}

// fmt.Println(str_arr)
// fmt.Println(len(str_arr)) // Length of the array
// fmt.Println(cap(str_arr)) // Capacity of the array, same as of the length

/// Iterating over arrays
// 	for i:=0; i<len(str_arr);i++ {
// 		fmt.Println(i,str_arr[i])
// 	}
// }

// ______________________________________

// => Slices : Slices are also used to store multiple values of the same type in a single variable, however unlike arrays, the length of a slice can grow and shrink as we see fit.

// -> There are several ways to create a slice :
// 	>-> Using the []data_type{values} format
// 	>-> Create a slice from an array
// 	>-> Using the make() function : The make function will create a zeroed array and return a slice referencing an array. This is a great way to create a dynamically sized array. To create a slice using the make function, we need to specify three arguments: type, length and the capacity.

// package main

// import (
// 	"fmt"
// )

// func main() {
// num := []int{10,20,30,40,50}

// fmt.Println(num[:]) // Return all elements
// fmt.Println(num[0:3]) // Returns a slice from index 0 to index 2 (which is 3)
// fmt.Println(num[2:]) // Returns a slice from index 2 to end of the array

/// Append values to a slice

// num = append(num, 60,70,80)
// fmt.Println(num) // Returns [10 20 30 40 50 60 70 80]

// make() function

// slice := make([]int, 3, 5) // 3 is the length and 5 is the capacity

// slice[0] = 10
// slice[1] = 20
// slice[2] = 30
// // slice[3] = 40 // This will throw an error as the length of the slice is 3, but we are trying to assign value to index 3

// fmt.Println("Length",len(slice))
// fmt.Println("Capacity",cap(slice))
// fmt.Println(slice)

/// Iterating over slices

// ppls := []string{"Aizen", "Zoro", "Goku", "Naruto", "Luffy"}

// // for i:=0; i<len(ppls);i++ {
// // 	fmt.Println(i, ppls[i])
// // }

/// ->> for range loop
// for index, value := range ppls {
// 	fmt.Println(index, value)
// }
// }

// ______________________________________

// => Maps : Maps are a data structure which allow us to store data values in key-value pairs. A map is an unordered and changeable collection that does not allow duplicates. The default value of a map is nil.

// package main

// import (
// 	"fmt"
// )

// func main(){

/// Create a map with string:int key-value pair
// userInfo := map[string]int{
// 	"John": 25,
// 	"Bhim": 30,
// 	"Bob": 35,
// 	"Mike": 40,
// }

/// Getting & Setting a map value using key
// userInfo["John"] = 50 // Get & set the value of John

// fmt.Println(userInfo) // Returns the map

/// Iterating over a map
// for prop, value := range userInfo {
// 	fmt.Println(prop, value) // Prints the key-value pair
// }
// }

// ______________________________________

// => Structs (Structures) : A struct is used to create a collection of members of different data types, into a single variable. (Same as interface/type in TS)

// package main

// import (
// 	"fmt"
// )

// type Person struct {
// 	name string
// 	age int
// 	job string
// 	salary int
// }

/// Nested Structs
// type Programmer struct {
// 	Person
// 	isProgrammer bool
// }

// func main(){

// var user Person // Creating a variable of type Person...under the hood user := Person{}

// user.name = "John"
// user.age = 25
// user.job = "Developer"
// user.salary = 50000

// fmt.Println(user) // OUTPUT - {John 25 Developer 50000}

/// Nested Structs

// h:= Programmer{
// 	Person : Person{
// 		name: "John",
// 		age: 25,
// 		job: "Developer",
// 		salary: 50000,
// 	},
// 	isProgrammer: true,
// }

// fmt.Println(h) // OUTPUT - {{John 25 Developer 50000} true}

/// Anonymous Structs : Anonymous structs are structs that do not have a name. They are useful when we need to create a struct for a one-time use, and we do not need to reuse it later.

// 	u := struct {
// 		fname string
// 		lname string
// 	}{
// 		fname: "John",
// 		lname: "Doe",
// 	}

// 	fmt.Println(u) // OUTPUT - {John Doe}
// }

// ______________________________________

// => Functions

// package main

// import (
// 	"fmt"
// )

// func hello() {
// 	fmt.Println("Hello from function")
// }

/// In Go, pass the function parameter as (variable_name variable_type)
// func greet(name string) {
// 	fmt.Println("Hello", name)
// }

/// ->> Variadic Function : A variadic function is a function that can take a variable number of arguments. The variadic function must be the last parameter in the function signature. The variadic parameter is represented by an ellipsis (...) followed by the type of the parameter. The variadic parameter is treated as a slice of the specified type. (Its like the rest operator in JS)
// func showUsers(s ...string){
// 	fmt.Println("Users are", s)
// }

/// a, b int means that the function takes two parameters of type int, it has a return type of int. We can also manually write the parameter as (a int, b int)
// func add(a,b int)int{
// 	return a+b
// }

// func main(){

// hello() // Calling the void function
// greet("Aizen") // Calling the function with parameter

// showUsers("Aizen", "Zoro", "Goku", "Naruto", "Luffy") // Calling the function with multiple parameters, it can take any number of parameters & it'll be treated as a slice/array

// fmt.Println(add(10,20)) // Calling the function with return type

/// ->> Functional Expression : When we store the function in a variable, we can call the function using the variable name.

// add_num := func(a,b int)int{ return a + b }

// fmt.Println(add_num(10, 20)) // Calling the function using the variable name

/// ->> Anonymous Function : A function that does not have a name. It is useful when we need to create a function for a one-time use, and we do not need to reuse it later.

// func() {
// 	fmt.Println("Hello from anonymous function")
// }() // Calling the anonymous function

/// ->> Closure Function : A closure function is a function that has access to the variables in its lexical scope, even when the function is executed outside that lexical scope. It is useful when we need to create a function that has access to the variables in its parent function, even after the parent function has returned.

// closure := func() func() {
// 	x := 10
// 	return func() {
// 		fmt.Println(x)
// 	}
// }
// closure()() // Calling the closure function
// }

// ______________________________________

// => Methods : A method is a function associated with a particular type. It is a way to define behavior for a specific type. In Go, methods are declared with a receiver, which is a special type of parameter that appears in the method signature. The receiver indicates on which type the method operates.(Its like class in OOP)

// package main

// import "fmt"

// type Person struct {
// 	gender string
// 	fname string
// 	lname string
// 	age int
// 	profession string
// 	isMarried bool
// }

/// Method with receiver (first parameter), it means that the method is associated with the Person struct.

// func (p Person) printInfo(){
// 	fmt.Println(p)
// }

// func main(){

// 	p1 := Person{
// 		gender: "male",
// 		fname: "John",
// 		lname: "Doe",
// 		age: 25,
// 		profession: "Developer",
// 		isMarried: false,
// 	}

// 	p1.printInfo() // Calling the method using the object of the struct
// }

// ______________________________________

// => Callback Functions : If a function is passed as an argument to another function, then such types of functions arer known as Higher-Order function. This passing function as an argument is also known as a callback function or first-class function.

// package main

// import (
// 	"fmt"
// )

// func addName (name string, callback func(string)){
// 	callback(name)
// }

// func main(){

// 	addName("Aizen", func(name string){
// 		fmt.Printf("Hello %v", name)
// 	}) // Calling the function with callback function as an argument
// }

// ____________________________________

// => Defer Keyword : The defer keyword is used to delay the execution of a function or a statement until the nearby function returns. In simple words, defer will move the execution of the statement to the very end inside a function.

// package main

// import (
// 	"fmt"
// )

// func main(){

// 	defer fmt.Println("Buddies")
// 	fmt.Println("Heyy")

// }

// ____________________________________

// => Scope : Think of scope as the visibility or accessibility of things (like variables or functions) in our code. It's like where things are known or can be used.

// -> Block Scope : Variables declared inside a block (like inside a function or a loop) are only accessible within that block. Once you leave that block, those variables are no longer available.

// -> Function Scope : Variables declared inside a function are only accessible within that function. They are not visible outside of it.

// -> Package Scope : Variables declared outside of any function but within a package are accessible to all functions within that package. They are not accessible from other packages unless they are exported (start with an uppercase letter).

// package main

// import (
// 	"fmt"
// )

// var p = 24 // Package scope

// func main(){

// 	m := 20 // Function scope
// 	fmt.Println(m)

// 	for i:=0; i<10; i++ {
// 		fmt.Println(i) // Block scope
// 	}

// 	fmt.Println(p) // Package scope
// }

// __________________________________

// => Pointers : A pointer in Go is a variable that stores the memory address of another variable. By using pointers, we can indirectly access and modify the value of the variable whose address is stored in the pointer.
// We can create a pointer using the * (asterisk) symbol. The & (ampersand) symbol is used to obtain the memory address of a variable.

// package main

// import (
// 	"fmt"
// )

// func main(){

// 	x := 20

// 	var pointerToX *int = &x

// 	fmt.Println("Memory address of x: ",pointerToX) // Printing the memory address of x

// 	fmt.Println("Value pointed to by pointerToX", *pointerToX) // Dereferencing the pointer to get the value of x
// }

// _________________________________

// => panic() : Similar to exceptions raised at runtime when an error is encountered, panic() is either raised by the program itself when an unexpected error occurs or the programmer throws the exception on purpose for handling particular errors. (Its like try-catch block's catch part)

// package main

// import "fmt"

// func employee(name *string, age int){ // Pointing the address of the variable to the function
// 	if age > 60 {
// 		panic("Age cannot be greater than retirement age !")
// 	} else {
// 		fmt.Println("Employee Name:", *name)
// 	}
// }

// func main(){

// 	empName := "John"
// 	age := 59

// 	employee(&empName, age) // Passing the address of the variable to the function
// }

// _________________________________

// => Interface : An interface is a type that specifies a set of method signatures. It defines a contract for what methods a type must have, but it does not provide the implementation is provided by concrete types that satisfy the interface. (Its like abstract class in OOP)

// e.g - 1

// package main

// import (
// 	"fmt"
// )

// // Define an interface called Animal. Any type that has a GetInfo() string method will satisfy this interface
// type Animal interface {
// 	GetInfo() string
// }

// // Define a struct for Cat
// type Cat struct {
// 	Name  string
// 	Color string
// }

// // Define a struct for Dog
// type Dog struct {
// 	Name  string
// 	Breed string
// }

// // Implement the GetInfo() method for Cat. This satisfies the Animal interface
// func (c Cat) GetInfo() string {
// 	return fmt.Sprintf("Cat: %s, Color: %s", c.Name, c.Color)
// }

// // Implement the GetInfo() method for Dog. This also satisfies the Animal interface
// func (d Dog) GetInfo() string {
// 	return fmt.Sprintf("Dog: %s, Breed: %s", d.Name, d.Breed)
// }

// // This function takes any Animal type and prints its info. It uses interface-based polymorphism
// func printAnimalInfo(animal Animal) {
// 	fmt.Println(animal.GetInfo())
// }

// func main() {
// 	// Create a Cat instance
// 	cat := Cat{Name: "Ken", Color: "Orange"}
// 	// Pass the Cat to the printAnimalInfo function
// 	printAnimalInfo(cat)

// 	// Create a Dog instance
// 	dog := Dog{Name: "Max", Breed: "Husky"}
// 	// Pass the Dog to the printAnimalInfo function
// 	printAnimalInfo(dog)
// }

// e.g - 2

// package main

// import (
// 	"fmt"
// )

// // Step 1: Define the interface with the method signature
// type Describer interface {
// 	Describe() string
// }

// // Step 2: Define different struct types
// type Car struct {
// 	brand string
// 	model string
// }

// type Person struct {
// 	name      string
// 	profession string
// 	age        int
// }

// // Step 3: Implement the Describe method for each struct type. These methods satisfy the Describer interface
// func (c Car) Describe() string {
// 	return fmt.Sprintf("Car is a %v %v model", c.brand, c.model)
// }

// func (p Person) Describe() string {
// 	return fmt.Sprintf("My name is %v, %v years old, working as a %v", p.name, p.age, p.profession)
// }

// // Step 4: Global function that takes a Describer and calls its Describe method
// func printDescribeInfo(d Describer) {
// 	fmt.Println(d.Describe())
// }

// func main() {
// 	// Create a Person and describe it
// 	person := Person{name: "John", age: 28, profession: "Backend Developer"}
// 	printDescribeInfo(person)

// 	// Create a Car and describe it
// 	car := Car{brand: "Toyota", model: "Hyryder"}
// 	printDescribeInfo(car)
// }

// e.g - 3

// package main

// import (
// 	"fmt"
// )

// // Define interface with multiple methods signatures
// type Human interface {
// 	UserName() string
// 	Profession() string
// }

// // Create different users structs
// type User1 struct{
// 	Age int
// 	Gender string
// }

// type User2 struct{
// 	Gender string
// 	Age int
// 	Location string
// }

// // Methods for each users
// func (j User1) Profession() string {
// 	return "Frontend Dev"
// }

// func (j User1) UserName() string {
// 	return "John"
// }

// func (a User2) Profession() string {
// 	return "Backend Dev"
// }

// func (a User2) UserName() string {
// 	return "Alisha"
// }

// // Info function for calling the methods
// func printInfo(h Human) {
// 	fmt.Println("My name is ",h.UserName()," and I'm a professional ",h.Profession())
// }

// func main(){

// 	person1 := User1 {Gender: "Male", Age: 25}
// 	printInfo(person1)

// 	person2 := User2 {Gender: "Female",Age: 22,Location: "Finland"}
// 	printInfo(person2)

// }

// _________________________________

// => Generics : Generics in Go provide a way to write functions and data structures that can work with any data type while maintaining type safety. Generics in Go involve defining functions or data structures with type parameters. Type parameters are specified inside square brackets ([]) before the function or type signature. It basically prevent us creating multiple function which does the same work but with different types.

// e.g - 1

// package main

// import (
// 	"fmt"
// )

/// -> Without Generic Function

// func printInteger(item, defaultValue int) (int,int) {
// 	return item, defaultValue
// }

// func printString(item, defaultValue string) (string,string) {
// 	return item, defaultValue
// }

// func printBoolean(item, defaultValue bool) (bool,bool) {
// 	return item, defaultValue
// }

/// -> With Generic Function

// func printItem[T any] (item, defaultValue T) (T, T) {
// 	return item, defaultValue
// }

// func main(){

/// Without Generic Function
// // num1, num2 := printInteger(10, 20)
// // str1, str2 := printString("One", "Two")
// // bool1, bool2 := printBoolean(true, false)

// // fmt.Println(num1,num2)
// // fmt.Println(str1,str2)
// // fmt.Println(bool1,bool2)

/// With Generic Function
// num1, num2 := printItem(10, 20)
// str1, str2 := printItem("One", "Two")
// bool1, bool2 := printItem(true, false)

// 	fmt.Println(num1,num2)
// 	fmt.Println(str1,str2)
// 	fmt.Println(bool1,bool2)

// }

// e.g - 2

// package main

// import (
// 	"fmt"
// )

// func PrintSlice[T any](s [] T) {
// 	for _,v := range s {
// 		fmt.Println(v)
// 	}
// }

// func main(){

// 	// Integers Slice
// 	intSlice := []int{1, 2, 3, 4, 5}
// 	PrintSlice(intSlice)

// 	// Strings Slice
// 	strSlice := []string{"Aizen", "Zoro", "Goku", "Ichigo", "Vegeta"}
// 	PrintSlice(strSlice)

// 	// Boolean Slice
// 	boolSlice := []bool{true, false, true, true}
// 	PrintSlice(boolSlice)

// }

// _________________________________

// => Goroutines : Goroutine is a lightweight, independently running function that allows concurrent execution in a program. It's like a mini-thread managed by the Go runtime, making it easy to run multiple tasks simultaneously. Goroutines are a key feature for concurrent programming in Go language.

// -> Characteristics :

// 1. Concurrency : Goroutines let us do multiple things at the same time in a straight forward way. It's like juggling several tasks without droppin any tasks.

// 2. Low Overhead : Goroutines are like very lightweight helpers that don't take up much space. We can have lots of them in our program without getting it too heavy.

// 3. Concurrency With Share Memory : Goroutines communnicate and synchronize using shared memory (via channels), making it easy to coordinate the execution of concurrent tasks.

// 4. Language-Level Support : Goroutines are like little workers provided by the Go language itself. We don't have to worry about complicated details because the Go language helps us manage them easily. It's like having a team of worker's that already know how to cooperate.

// --------------------------------------------
// ## Concurrent Programming :

// Concurrent programming is a style of programming where multiple tasks can be in progress at the same time.    Tasks start, run, and progress independently — but not necessarily at the exact same moment.
// Lets imagine we're cooking a meal - we boil water, while it’s boiling, you start chopping vegetables, then you fry onions while the rice is cooking. These tasks overlap — we're not finishing one before starting the next. That’s concurrency!
// Concurrency : Multiple tasks in progress — might switch back and forth (interleaving). Parallelism : Multiple tasks executing at the exact same time — usually on multiple CPUs/cores. So, all parallel programs are concurrent, but not all concurrent programs are parallel.
// --------------------------------------------

// -> Syntax :
// go functionName()

// package main

// import (
// 	"fmt"
// 	"time"
// )

// func printNumbers() {
// 	for i := 1; i <= 5; i++ {
// 		time.Sleep(100 * time.Millisecond) // Pause for 100 milliseconds between each number
// 		fmt.Printf("%d ", i)
// 	}
// }

// func main() {

// Launch printNumbers() as a goroutine. This means it will run in the background (concurrently with the main function). Acts like a thread.

// go printNumbers()

// This will print immediately because the goroutine runs asynchronously.

// fmt.Println("This is the main function call.")

/// Prevent the main function from exiting before the goroutine finishes. printNumbers() takes approximately 500 ms total (100ms * 5). So we wait 1 second to ensure it completes. If we don't add this delay, the program may finish before the goroutine runs.

// time.Sleep(1 * time.Second)
// }

// _________________________________

// => Channels : Channels in Go are like communication pipes between different parts of our program. They let different pieces of our code talk to each other by sending and receiving messages. It's like passing notes or messages between friends to coordinate what needs to be done.

// -> Imagine we have two friends working on a project. One friend can put notes in the channel, and the other friend can pick up those notes. This helps them share information and work together without getting in each other's way. That's what channels do in Go - they help different parts of our program  communicate smoothly.

// -> Creating a channel is like making a special mailbox for messages. We can use the make() function to create this mailbox. It's like setting up a mailbox with a specific type of messages it can hold.

// package main

// import (
// 	"fmt"
// )

// s3 - sendData() is a function that will run in a separate goroutine. It sends the integer value 20 into a channel.
// func sendData(ch chan<- int) {
// 	fmt.Println("Sending the value...")
// 	ch <- 20 				// Sends the value 20 into the channel. Blocks until the value is received on the other end.
// }

// func main(){

// s1 - Create a channel that can  carry integers. Channels are like pipes through which goroutines can send/receive values.
// 	ch := make(chan int)

// s2 - Start goroutine which will run concurrently with main(). This means sendData() will start running concurrently, without blocking the rest of main(). But ch <- 20 will block inside sendData() until someone is ready to receive that value.
// 	go sendData(ch)

// s4 - Receive the value from the channel. The program will wait (block) here until it receives a value from the channel.
// 	value := <-ch

// 	fmt.Println("Received : ",value)
// }

// -> Unbuffered Channels :

// 1. An unbuffered channel has a capacity of 0.
// 2. Each send operation on an unbuffered channel blocks until there is a corresponding receive operation, and vicce versa.
// 3. This creates a direct, synchronous communication between the sender and the receiver.

// -> Buffered Channels :

// 1. A buffered channel has a specified capacity greater than 0.
// 2. It allows multiple values to be sent into the channel without an immediate corresponding receiver.
// 3. Send operations on a buffered channel block only when the buffer is full, and receive operations block only when the buffer is empty.

// package main

// func main(){

// Unbuffered Channel: This channel has no capacity, so both send and receive operations will block until the other side is ready.
// unbufferedChannel := make(chan int)

// Start a goroutine with anonymous func to send a value into the unbuffered channel
// go func(){
// 	unbufferedChannel <- 20 // This will block until the main goroutine is ready to receive
// }()

// Main goroutine receives the value from the channel. The receive will unblock the send above
// value := <- unbufferedChannel

// fmt.Println("Received from unbuffered channel :", value)

// Buffered Channel with a capacity of 2.  Unlike unbuffered, this allows up to 2 values to be sent without an immediate receiver
// bufferedChannel := make(chan string, 2)

// Start a goroutine with anonymous function that sends 2 messages into the buffered channel
// go func(){
// 	bufferedChannel <- "Hello" // First value goes into the buffer
// 	bufferedChannel <- "World" // Second value fills the buffer
// 	close(bufferedChannel) // Close the channel to signal no more data will be sent
// }()

// Receiving from buffered channel normally. This also works but only receives fixed number of values. So I used two variables for receiving two values.
// value1 := <- bufferedChannel
// value2 := <- bufferedChannel

// fmt.Println(value1,value2)

// Receiving from buffered channel using a for-range loop. This is preferred when you want to receive all values until the channel is closed.
// for msg := range bufferedChannel {
// 	fmt.Print(msg," ")
// }

// }

// _________________________________

// => Wait Groups : In Go, WaitGroups (from the sync package) are used to wait for a collection of goroutines to finish executing. It’s a way to coordinate concurrent tasks.

// Lets, imagine we have a group of friends working on different tasks, and we want to make sure everyone finishes before moving on. A wait group is like a coordinator or a counter that helps us wait for all our friends to complete their tasks.

// package main

// import (
// 	"fmt"
// 	"sync"
// )

// func printMessage(wg *sync.WaitGroup, msg string) {
// 	defer wg.Done() // Signal that this goroutine is done
// 	fmt.Println(msg)
// }

// func main() {

// -> Create a WaitGroup to synchronize goroutines (Like creating the friends group)
// var wg sync.WaitGroup

// -> Add goroutines to the WaitGroup before starting them
// wg.Add(3) // We are waiting for 3 goroutines (like adding 3 friends in their group)

// -> Start goroutines (our "friends" doing tasks)
// Goroutines are not synchronous. Even if you started them in a certain order, they run in parallel and may finish in any order, depending on: CPU scheduling, Execution time, OS thread handling, Other factors like system load.
// go printMessage(&wg, "Hello")
// go printMessage(&wg, "Go")
// go printMessage(&wg, "WaitGroup")
// The output order is non-deterministic because goroutines run concurrently and are scheduled independently by the Go runtime.

// -> Wait for all goroutines to complete (Waiting for completion of friend's tasks)
// 	wg.Wait()

// 	fmt.Println("All messages printed!")
// }

// _________________________________

// Select Statement : The select statement in Go is used for handling multiple channels in a non-blocking way. It's a powerful feature that allows us to wait on multiple channels and execute the first one that is ready, much like a switch statement but for channels.

package main

import (
	"fmt"
	"time"
)

func main(){
	// Create two channels to communicate
	ch1 := make(chan string)
	ch2 := make(chan string)

	// Goroutine 1: sends a message after 2 seconds to channel ch1
	go func(){
		time.Sleep(2 * time.Second)
		ch1 <- "Message from channel 1"
	}()
	
	// Goroutine 2: sends a message after 4 seconds to channel ch2
	go func(){
		time.Sleep(4 * time.Second)
		ch2 <- "Message from channel 2"
	}()

	// Select statement begins here to listen to both channels
	select {
	case msg1 := <- ch1: // If ch1 is ready (after 2 seconds), print its message
		fmt.Println(msg1)

	case msg2 := <- ch2: // If ch2 is ready (after 4 seconds), print its message
		fmt.Println(msg2)

	case <- time.After(5 * time.Second): // Timeout after 5 seconds. If no channel is ready within 5 seconds, print this message
		fmt.Println("No communication is ready to receive")
	}
}
