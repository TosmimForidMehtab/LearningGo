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
int, float32, float64, bool, complex64, complex128
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
-   ### Make
    -   Allocate memory and initialise
    -   non-zeroed storage
-   Deallocation happens automatically

## Pointers

### Pointres are ez bcoz i'm from C++

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








