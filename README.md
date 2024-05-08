# Learning Golang

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
```
