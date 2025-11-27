# Go Lang Documentation

This repository contains documentation and resources for learning and using the Go programming language (Golang). Go is an open-source programming language designed for simplicity, efficiency, and reliability.

## Table of Contents

- [Getting Started](#getting-started)
- [Installation](#installation)
- [Basic Concepts](#basic-concepts)
- [Making My first Module](#making-my-first-module)
- [Hello World Example](#hello-world-example)
- [Go fmt](#go-fmt)
- [Go Vet](#go-vet)
- [Makefile](#makefile)
- [The Go Compatibility Promise](#the-go-compatibility-promise)

## Getting Started

To get started with Go, you need to have Go installed on your machine. You can download the latest version of Go from the official website: [https://golang.org/dl/](https://golang.org/dl/).

## Installation

Follow the instructions for your operating system to install Go. After installation, you can verify the installation by running the following command in your terminal:

```bash
go version
```

This should display the installed version of Go. Example output:

```bash
go version go1.25.4 windows/amd64
```

## Basic Concepts

Go is a statically typed, compiled language with a syntax similar to C. It supports [Concurrent](https://www.google.com/search?q=define+concurrent+in+go) programming through [Goroutines](https://go.dev/tour/concurrency/1) and [Channels](https://www.google.com/search?q=define+channels+in+go), making it suitable for building scalable applications.

### Making My first Module

To create a new Go module, follow these steps:

1. Create a new directory for your module and navigate into it:

   ```bash
   mkdir ch1
   cd ch1
   ```

2. Initialize a new Go module using the `go mod init` command followed by your module name (usually a repository path):

   ```bash
    go mod init hello_world
    ```

This command creates a `go.mod` file in your directory, which defines the module's path and its dependencies.

A go module is a collection of related Go packages that are versioned together. It allows you to manage dependencies and versioning for your Go projects effectively.

Inside the go.mod file, you will see something like this:

```Go
module hello_world
go 1.25
```

This indicates the module path and the Go version used.
You can now start adding Go files to your module and manage dependencies using Go's module system.

You should not edit the go.mod file manually; instead, use Go commands to manage dependencies and versions. Example commands include `go get`, `go tidy`, and `go mod download`.

### Hello World Example

Here is a simple "Hello, World!" program in Go:

```go
package main

import "fmt"

func main()  {
    fmt.Println("Hello, World!")
} 
```

Save this code in a file named `hello.go` and run it using the following command:

```bash
go build
```

The output will be an executable file that you can run to see the message "Hello, World!" printed to the console.

```bash
./hello_world
Hello, World!
```

Let's talk about this code. The first line is the package declaration, which defines the package name. In this case, it's `main`. Within a Go module, code is organized into packages. The `main` package is a special package that serves as the entry point for executable programs.

Next they is the import statement, which imports the `fmt` package. The `fmt` package provides functions for formatted I/O (Input & Output) operations, such as printing to the console.

In order to change the name of the output file, you can use the `-o` flag followed by the desired name when running the `go build` command. For example:

```bash
go build -o hello_app
```

### Go fmt

Go provides a built-in tool called `go fmt` that automatically formats your Go code according to standard conventions. To format your code, simply run the following command in your terminal:

```bash
go fmt -w hello.go
```

Using the `./...` tells `go fmt` to format all Go files in the current directory and its subdirectories.

The go fmt tool helps maintain consistent code style across Go projects, making it easier to read and collaborate with other people.

#### Semicolon Insertion

In Go, semicolons are automatically inserted at the end of statements during the compilation process. This means that you do not need to explicitly write semicolons at the end of each line. The Go compiler determines where to insert semicolons based on the structure of your code.

For example, the following code is valid in Go:

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

In this example, there are no semicolons at the end of the lines, but the Go compiler will automatically insert them where necessary during compilation.

The semicolon insertion rules in Go are designed to make the code more readable and reduce clutter. However, there are certain situations where you may need to use semicolons explicitly, such as when writing multiple statements on a single line. In addition, the way in which you format your code (e.g., line breaks, indentation) can affect how semicolons are inserted. For example, if you place multiple statements on the same line, you will need to separate them with semicolons.

For example:

```go
package main; import "fmt"; func main() { fmt.Println("Hello, World!"); }
```

This code is valid because the semicolons explicitly separate the statements on the same line.

### Go Vet

Go provides a built-in tool called `go vet` that analyzes your Go code for potential errors and issues. It performs static analysis to identify common mistakes, such as incorrect format strings, unreachable code, and other potential problems. To run `go vet`, use the following command in your terminal:

```bash
go vet ./...
```

The `./...` tells `go vet` to analyze all Go files in the current directory and its subdirectories.
The `go vet` tool helps catch potential issues early in the development process, improving code quality and reliability. It is recommended to run `go vet` regularly as part of your development workflow.

## Makefile

A Makefile is a special file used to automate the build process of a project. It contains a set of rules and instructions that define how to compile and link the code, as well as other tasks such as cleaning up generated files or running tests.

Here is an example of a simple Makefile for a Go project:

```Makefile
.PHONY: build clean run
build:
    go build -o hello_app
clean:
    rm -f hello_app
run: build
    ./hello_app
```

In this Makefile, we have defined three targets: `build`, `clean`, and `run`. The `build` target compiles the Go code and creates an executable file named `hello_app`. The `clean` target removes the generated executable file. The `run` target first builds the project and then runs the resulting executable.

## The Go Compatibility Promise

The Go Compatibility Promise is a commitment made by the Go development team to ensure that programs written in Go will continue to work with future versions of the language. This promise guarantees that code written in a specific version of Go will remain compatible with all subsequent versions, as long as it adheres to the language's specifications and does not rely on undefined behavior.
