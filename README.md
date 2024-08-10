# Learning Golang

## Reasons to choose Golang
-	Faster Build time
-	Fast Sartup time
-	Performance and Efficiency
-	Concurrency Model
-	Static Typing
-	Less Build bundle size

## Variables

### Declaration

```go
    var name string = "Kekekeke"
```

**OR**

```go
    var name = "Kekekeke"
```

**OR**

```go
    name := "Kekekeke"
```

_NOTE: The ':=' operator does not work in the global scope_

## Data Types

```go
int, string, float32, float64, bool, complex64, complex128
```

**Check type of variables/values:**
```go
	fmt.Println(reflect.TypeOf(name)) // string
	fmt.Println(reflect.TypeOf(25)) // int
```
**Casting one type to another:**
```go
	a := 1.3
	fmt.Println(reflect.TypeOf(a)) // float64
	b := int(a)
	fmt.Println(reflect.TypeOf(b)) // int
	fmt.Println(b) // 1

	// String to int
	c := "100"
	d, err := strconv.Atoi(c)
	if err == nil{
		fmt.Println(reflect.TypeOf(d)) // int
	}
	// int to string
	e := 100
	f := strconv.Itoa(e)
	fmt.Println(reflect.TypeOf(f)) // string

	// string to float
	g := "100.5"
	h, err := strconv.ParseFloat(g, 64)
	if err == nil {
		fmt.Println(reflect.TypeOf(h)) // float64
	}

	// float to string
	i := 100.5
	j := strconv.FormatFloat(i, 'f', 1, 64)
	fmt.Println(reflect.TypeOf(j)) // string
	// OR
	k := fmt.Sprintf("%f", i)
	fmt.Println(reflect.TypeOf(k)) // string
```

## Taking input

```go
    reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your name: ")

	input, _ := reader.ReadString('\n')
	fmt.Println("Hello", input)

    var a, b int
	fmt.Println("Enter 2 numbers: ")
    num, err := strconv.ParseInt(strings.TrimSpace(input), 10, 64)
```

**OR**

```go
    var a, b int
	fmt.Println("Enter 2 numbers: ")
	fmt.Scan(&a, &b)
	fmt.Println(a, b)
```

_NOTE: Properties that have the starting letter capitalised are Public_

## Running Go Program

-   ### Compile and then run
    -   `go build <file_name>`
    -   `./a.exe`
        **OR** `./a.out`
-   ### Compile and run together
    -   `go run <file_name>`

## Initialising go module

#### **A module is a collection of packages**

`go mod init <module_name>`

## Building for different Platforms

We can build our go code for a different platform using this:
`GOOS=<platform_name> go build`
Platform name can be:

-   `darwin` for macOS
-   `windows` for windows
-   `linux` for linux

## Exploring Go env variables

-   GOPATH is where your Go projects reside.

-   GOBIN is where executable binaries will be placed after compilation.

-   GOOS specifies the operating system for which you're compiling.

    -   linux
    -   windows
    -   darwin(macOS)

-   GOARCH specifies the architecture for which you're compiling.
    -   amd64
    -   arm
-   GOROOT points to the directory where Go itself is installed.

## Memory Management

-   `new` and `make` are used for allocation
-   ### New
    -   Allocate but no INIT
    -   zeroed storage
	-	Returns a pointer
-   ### Make
    -   Allocate memory and initialise
    -   non-zeroed storage
	-	Returns a value
-   Deallocation happens automatically

## Pointers

### Pointres are ez bcoz i'm from C++

## Strings
-   Strings are immutable in Go
-   Strings are slices of bytes
```go
	// Replace a string with another
	s := "Hello World"
	replacer := strings.NewReplacer("World", "Tos")
	s = replacer.Replace(s)
	fmt.Println(s)      // Hello Tos
	fmt.Println(len(s)) // 9

	// Check if a string contains another string
	if strings.Contains(s, "Hello") {
		fmt.Println("Hello is in the string") // This is executed and printed
	} else {
		fmt.Println("Hello is not in the string")
	}

	// Find the first occuring index of a string
	fmt.Println(strings.Index(s, "o")) // 4

	// Search and replace every o with 0
	fmt.Println(strings.Replace(s, "o", "0", -1)) // Hell0 T0s

	// Split string
	fmt.Println(strings.Split(s, " ")) // [Hello Tos]

	// Uppercase and lowercase
	fmt.Println(strings.ToUpper(s)) // HELLO TOS
	fmt.Println(strings.ToLower(s)) // hello tos

	// Check if a string starts with another string
	fmt.Println(strings.HasPrefix(s, "He")) // true

	// Check if a string ends with another string
	fmt.Println(strings.HasSuffix(s, "ld")) // false
```
## Runes
-	Runes are used to handle individual characters in strings, especially when dealing with multi-byte characters or special characters
-	It's an alias for the int32 type
```go
	str := "abcdefg"
	fmt.Println(utf8.RuneCountInString(str)) // 7

	// Printing the index, utf8 code and character
	for idx, val := range str {
		fmt.Printf("%d : %#U : %c\n", idx, val, val)
	}
	// OUTPUT:
	/*
		0 : U+0061 'a' : a
		1 : U+0062 'b' : b
		2 : U+0063 'c' : c
		3 : U+0064 'd' : d
		4 : U+0065 'e' : e
		5 : U+0066 'f' : f
		6 : U+0067 'g' : g
	*/

	// Converting string to runes
	r := []rune(str)
	for _, val := range r {
		fmt.Println(val)
	}

	// Converting byte array to string
	b := []byte{'a', 'b', 'c'}
	bStr := string(b[:])
	fmt.Println(bStr)
```
## Arrays

-   Arrays are strange and weird
-   Only `len()` is available

```go
    var a = [5]int{1, 2, 3, 4}

	fmt.Println(a)      // [1 2 3 4 0] The last item is 0 as it is not initialized
	fmt.Println(len(a)) // 5

	var b [5]string
	b[0] = "Hello"
	b[1] = "World"
	b[2] = "!"
	b[3] = "!"

	fmt.Println(b)      // [Hello World ! ! ] The last item is an empty space as it is not initialized
	fmt.Println(len(b)) // 5
```

## Slices

-   Slices are awesome

```go
    // Not providing the size makes it a slice
	var a = []string{}
	fmt.Printf("%T", a) // []string
	fmt.Println()
	a = append(a, "Hello")
	a = append(a, "World")
	a = append(a, "!")
	fmt.Println(a) // [Hello World !]

	b := append(a[1:2]) // slicing similar to python
	fmt.Println(b)

	// Way 2 of creating a slice
	c := make([]int, 3) // c initialized with 3 elements with all values 0
	c[0] = 1
	c[1] = 3
	c[2] = 2
	// c[3] = 4 // out of range error
	c = append(c, 4, 5, 6) // No error
	fmt.Println(c)

	if !sort.IntsAreSorted(c) {
		sort.Ints(c)
		fmt.Println("Sorting...")
	}
	fmt.Println(c) // Sorted slice

	// Removing an element
	indexToRemove := 2
	c = append(c[:indexToRemove], c[indexToRemove+1:]...)
	fmt.Println(c) // [1 2 4 5 6]

	// Capacity
	fmt.Println(cap(c)) // 6

	// Copying a slice
	arr := []int{}
	arr = append(arr, 1, 2, 3, 4, 5)
	arr2 := make([]int, len(arr))

	copy(arr2, arr)
	fmt.Println(arr, arr2) // [1 2 3 4 5] [1 2 3 4 5]

	// Check if 2 slices are same
	// fmt.Println(arr == arr2) // This won't work
	fmt.Println(slices.Equal(arr, arr2)) // true

	// 2D slice
	arr3 := [][]int{{1, 2}, {3, 4}}
	fmt.Println(arr3) // [[1 2] [3 4]]
```

## Constants
- Declare variables that cannot be reassigned during the program's life.
```go
	const name string = "John"
	// name = "Joe" // Not allowed
	fmt.Println(name) // John

	// Grouping multiple constants
	const (
		pi       = 3.14
		language = "Go"
	)
	fmt.Println(pi, language) // 3.14 Go

```

## Loops
**For is the only looping construct in go. There is no while loop**
```go
	// For loop
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// There is no while loop construct in go
	// Instead we can use for loop
	i := 5
	for i < 10 {
		fmt.Println(i)
		i++
	}

	// For loop with break and continue. Once we reach 5, the looping will stop and we will skip 3
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		if i == 3 {
			continue
		}
		fmt.Println("Num -> ", i)
	}

	// Range bases loop
	for ele := range 5 {
		fmt.Println(ele) // 0 1 2 3 4
	}

	arr := []int{1, 2, 3, 4, 5}
	for index, value := range arr {
		fmt.Println(index, value)
	}


	// Infinite loop
	for {
		fmt.Println("Infinite loop")
	}
```

## If-Else
```go
	if condition {
		statements
	}

	if condition {
		statements
	} else {
		statements
	}

	if condition-1 {
		statements
	} else if condition-2 {
		statements
	}

	age := 18
	if age >= 18 {
		fmt.Println("You are an adult")
	}
	else if age >= 13 {
		fmt.Println("You are a teenager")
	}
	else {
		fmt.Println("You are a Kodomo")
	}	
	// The same thing can be written as:
	if age := 18; age >= 18 {
		fmt.Println("You are an adult")
	}
	else if age >= 13 {
		fmt.Println("You are a teenager")
	}
	else {
		fmt.Println("You are a Kodomo")
	}
```
**NOTE:** *There is no ternary operator in go*

## Switch-case
```go
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		println("It's the weekend")

	default:
		println("It's a weekday")
	}

	// Type switch
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			println("I'm a bool")
		case int:
			println("I'm an int")
		case string:
			println("I'm a string")
		default:
			println("Other", t)
		}
	}
	whatAmI(true)
	whatAmI("hello")
	whatAmI(23.5)
```
## Maps
- Maps are unordered collections of key-value pairs.

```go
	// Declaring a map
	m := make(map[string]int)
	// OR
	// m := map[string]int{"k0": 0}

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m)       // map: map[k1:7 k2:13]
	fmt.Println("len:", len(m))  // len: 2
	fmt.Println("k1: ", m["k1"]) // 7
	// Deleting map entry
	delete(m, "k2")
	fmt.Println(m)

	// Check if key exists
	val, ok := m["k1"]
	if ok {
		fmt.Println("Key exists", val)
	} else {
		fmt.Println("Key does not exist")
	}

	// Check if 2 maps are equal
	m1 := map[string]int{"foo": 1, "bar": 2}
	m2 := map[string]int{"foo": 1, "bar": 2}
	// fmt.Println(m1 == m2) // This won't work
	fmt.Println(maps.Equal(m1, m2)) // true

	// Iterating over map
	for k, v := range m1 {
		fmt.Println(k, "->", v)
	}

	// CLearing the map
	clear(m)
	fmt.Println(m) // map[]
```
**NOTE:** *If key does not exist in the map, it returns zeroed value*

## Functions
-	Functions are first class citizens in go.
-	Functions can be passed as arguments to other functions.
-	Functions can be returned from other functions.
-	Functions can be treated as variables.

```go
	package main

	import "fmt"

	func add(a int, b int) int {
		return a + b
	}

	func sub(a, b int) int {
		return a - b
	}

	func getNames() (string, string) {
		return "foo", "bar"
	}

	func checkEven(a int) bool {
		return a%2 == 0
	}

	func filter(a []int, f func(int) bool) []int {
		var result []int
		for _, v := range a {
			if f(v) {
				result = append(result, v)
			}
		}
		return result
	}

	func main() {
		println(add(1, 2))
		println(sub(1, 2))
		l1, l2 := getNames()
		println(l1, l2)

		arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		fmt.Println(filter(arr, checkEven))

	}

```

## Variadic
-	Concept of passing multiple number of arguments and accepting them as a slice.
```go
	func sum(nums ...int) int {
		total := 0
		for _, num := range nums {
			total += num
		}
		return total
	}

	func main() {
		total := sum(1, 2, 3, 4, 5)
		fmt.Println(total)
		intArr := []int{6, 7, 8, 9, 10}
		total = sum(intArr...)
		fmt.Println(total)
	}
```

## Closures
-	Binds the parent environment variables with the embedded environment
```go
	func activateGiftCard() func(int) int {
		balance := 100
		return func(debitAmount int) int {
			balance -= debitAmount
			return balance
		}
	}

	func main() {
		useGiftCard1 := activateGiftCard()
		fmt.Println(useGiftCard1(25)) // 75

		useGiftCard2 := activateGiftCard()
		fmt.Println(useGiftCard2(50)) // 50
	}
```

## Templates
-	Template is a way to generate code for different types.
```go
	type Number interface {
		int | int32 | int64 | float32 | float64
	}

	func addNums[T Number](numbers []T) T {
		var sum T
		for i := range numbers {
			sum += numbers[i]
		}
		return sum
	}

	func main() {
		nums := []int{1, 2, 3, 4, 5}
		sum := addNums(nums)
		nums2 := []float32{1.1, 2.2, 3.3, 4.4, 5.5}
		sum2 := addNums(nums2)
		fmt.Println("Int sum -> ", sum)
		fmt.Println("float32 sum -> ", sum2)

	}
```
## Math
-	Math library is used for mathematical operations. Some of functions are listed here
```go
	package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	// Random number
	seedSecs := time.Now().Unix()
	rand.NewSource(seedSecs)
	randNum := rand.Intn(10) + 1
	fmt.Println("Random :", randNum)

	// Absolute of a number
	fmt.Println("Absolute of -5 is:", math.Abs(-5)) // 5

	// Square root of a number
	fmt.Println("Square root of 25 is:", math.Sqrt(25)) // 5

	// Power of a number
	fmt.Println("2 to the power of 3 is:", math.Pow(2, 3)) // 8

	// Floor of a number
	fmt.Println("Floor of 3.7 is:", math.Floor(3.7)) // 4

	// Ceil of a number
	fmt.Println("Ceil of 3.7 is:", math.Ceil(3.7)) // 3

	// Min and max of numbers
	fmt.Println("Min of 10 and 20 is:", math.Min(10, 20)) // 10
	fmt.Println("Max of 10 and 20 is:", math.Max(10, 20)) // 20

	// Logarithm of a number base 10
	fmt.Println("Logarithm of 10 is:", math.Log10(10)) // 1

	// Logarithm of a number base 2
	fmt.Println("Logarithm of 8 is:", math.Log2(8)) // 3
}
```
## Format specifiers
- Format specifiers are used to format the output of the program.

-	`%d` for integers
-	`%f` for floating point numbers
-	`%s` for strings
-	`%x` for hexadecimal numbers
-	`%c` for characters
-	`%o` for octal numbers
-	`%t` for booleans
-	`%v` guesses based on datatype
-	`%T` for types

## Structs (Structure)
- A struct is a collection of fields of different data types.
```go
	type Engine struct {
		id     int16
		milage float32
	}
	type Car struct {
		brand      string
		reg_num    int
		is_working bool
		price      float64
		engine     Engine
	}

	func printCar(car Car) {
		fmt.Println("Brand ", car.brand)
		fmt.Println("Reg. ", car.reg_num)
		fmt.Println("Status ", car.is_working)
		fmt.Println("Price ", car.price)
		fmt.Println("Engine_id ", car.engine.id)
		fmt.Println("Milage ", car.engine.milage)
		fmt.Println("=========================================")
	}

	func main() {
		var car1 Car
		car1.brand = "Yotato"
		car1.reg_num = 12345
		car1.is_working = true
		car1.price = 220000
		car1.engine.id = 11
		car1.engine.milage = 22
		printCar(car1)

		car2 := Car{
			brand:      "ATAT",
			reg_num:    67890,
			is_working: true,
			price:      300000,
			engine: Engine{
				id:     21,
				milage: 30,
			},
		}
		printCar(car2)

		// Anonymous structs
		u := struct {
			firstName string
			lastName  string
		}{
			firstName: "John",
			lastName:  "Doe",
		}
		fmt.Println(u.firstName, " ", u.lastName)
	}
```

## Methods
-	A method is a function associated with a particular type. It is a way to define behavior for specific type
-	Methods are declared inside the type definition
-	Methods can be called on instances of the type
-	`func (receiver) func_name() {func_body}`
```go
	type Person struct {
		first_name string
		last_name  string
	}

	func (p Person) getFullName() string {
		return p.first_name + " " + p.last_name
	}
	func main() {
		person := Person{first_name: "John", last_name: "Doe"}
		fmt.Println(person.getFullName())
	}
```

## Defer keyword
-	Defer keyword is used to schedule a function to be called when the surrounding function returns
-	Defer is used to clean up resources like closing files, releasing locks, etc.
```go
	func greet() {
		defer fmt.Println("Hello, World!") // This will be executed at last
		fmt.Println("This is a greeting")
	}

	func main() {
		greet()
	}
```

## Interface
-	An interface is a set of methods that a type must implement
-	Interfaces are used to define a contract that a type must follow
-	Interfaces are defined using the `interface` keyword
```go
	type Speaker interface {
		Speak() string
	}

	type Cat struct{}

	func (c Cat) Speak() string {
		return "Meoww!"
	}

	func speak(speaker Speaker) {
		fmt.Println("Speaker says ", speaker.Speak())
	}

	func main() {
		speak(Cat{}) // Speaker says  Meoww!
	}
```

## Goroutines and Waitgroups
-	Goroutines are lightweight threads that can run concurrently with the main program
-	Goroutines are scheduled by the Go runtime
-	Goroutines are created using the `go` keyword
-	Wait group is like a coordinator that helps to Wait for other code to finish
-	Wait group is created using the `sync` module
```go
	func printNums(wg *sync.WaitGroup) {
		defer wg.Done()
		for i := 1; i <= 10; i++ {
			time.Sleep(100 * time.Millisecond)
			fmt.Printf("%d ", i)
		}
	}

	func main() {
		fmt.Println("Welcome")
		var wg sync.WaitGroup
		wg.Add(1)
		go printNums(&wg)
		wg.Wait()
	}
```

## Channels
-	Channels are a way to communicate between goroutines
-	Channels are typed, meaning they can only send and receive values of a specific type
-	Channels are created using the `chan` keyword
-	Channels can be buffered or unbuffered
-	Unbuffered channels have capacity of 0. Each send operation on an unbuffered channel blocks until there is a corresponding receive operation and vice versa.
-	A buffered channel has capacity of greater than 0. It allows multiple send operations into the channel without an immediate corresponding reeponse. Operations are blocked only when the buffer is full 
```go
	func sendInt(ch chan<- int) {
		ch <- 1
	}

	func main() {
		// Create a channel
		ch := make(chan int)
		go sendInt(ch) // If `go` not used, it will give error
		val := <-ch
		fmt.Println(val)

		// Unbuffered channel
		unbuff := make(chan int)
		// Buffered channel
		buff := make(chan string, 2)

		go sendInt(unbuff)
		unbuffVal := <-unbuff
		fmt.Println("Value from Unbuffered channel", unbuffVal)

		go func() {
			buff <- "Hello"
			buff <- "World"
			close(buff)
		}()
		// If we don't close the channel, it will block indefinitely
		for msg := range buff {
			fmt.Println("Data from Buffered channel", msg)
		}
	}
```
## Select
-	Select is a statement that waits on multiple channels and executes the first one that is ready
-	Select can be used with channels, timers, and even the default case
```go
	func main() {
		ch1 := make(chan string)
		ch2 := make(chan string)
		// We use the select statement to wait on both channels. If both channels are ready, the
		// select statement will execute the first case that is ready. If only one channel is ready,
		// the select statement will execute the case for that channel. If neither channel is ready,
		// the select statement will block until one of the channels is ready.

		go func() {
			time.Sleep(1 * time.Second)
			ch1 <- "Channel1 message"
		}()

		go func() {
			time.Sleep(2 * time.Second)
			ch2 <- "Channel2 message"
		}()

		select {
		case msg1 := <-ch1:
			fmt.Println("Received message from channel 1:", msg1)
		case msg2 := <-ch2:
			fmt.Println("Received message from channel 2:", msg2)
		case <-time.After(5 * time.Second):
			fmt.Println("Timeout occurred")
		}
	}
```

## Mutex
- Mutex is a lock that can be used to protect shared resources from concurrent access
- Mutex can be used to prevent data races in concurrent programs
```go
	package main
	import (
		"fmt"
		"sync"
		"time"
		)
		var mu sync.Mutex
		var data int
		func main() {
			go func() {
				mu.Lock()
				data = 10
				mu.Unlock()
			}()
			go func() {
				mu.Lock()
				fmt.Println("Data:", data)
				mu.Unlock()
			}()
			time.Sleep(1 * time.Second)
		}

```