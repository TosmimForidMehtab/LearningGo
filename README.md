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
