# Go Tutorial 

## Introduction
- Developed in 2007 and made public in 2009.
- Focus on simplicity and scalability - inspired from Python
- Statically typed language - Helps to catch type errors.
- Efficient and most suitable for Concurrent programming.

## Basics

- Executable is created based on package main and main function.
- Other functions have to be imported
- It doesnt have OOP concept but we can create like OO using type
- Avoid `this` or `self` which will break Go convention
- Use `_` if we are not using any variable (Eg: index in for loop)
- In a way it resembles C language Structs and Pointers.

## Main Function

- To instruct Go to excute the application, we must have package named main and a function named main.
- You cannot have two main functions in the main package.
- If you are building 3rd party library, you dont need a main function.
```
package main

import "fmt"

func main(){
    fmt.Println("Hello World")
}

```

## Variables
- It is statically typed language.
- A variable holds the value of specific type.
- You cant have unused variables in your code a it is code pollution.


## Error Handling
```
data, err = os.ReadFile("")
if err != nil {
    fmt.Println("")
}

```

- We can also return errors
```
func main(){
    return errors.New("Throw Error")
}
```

- If you want throw an error and stop the execution flow.

```
func main(){

    panic("Throw Error and Stop")
}
```

### Packages

To install the third party packages from go 

``` go get github.com/duktig-solutions/go-random-date-generator ```

this will add the below line in you `go.mod` file 

```
require github.com/duktig-solutions/go-random-date-generator v0.0.0-20220823170630-1c67478b7653 // indirect
```

When you get project, check for the go.mod file and go with `go.get` to install 
all the dependencies mentioned there.


### Pointer operations

**What is Pointer ?**
Any variable that stores the address is a Pointer.

**Why Pointers ?**
1. By default, Go creates a copy when passing values into function. In case of large variable which already occupies large space, till the scope of the function ends there is a copy created. To avoid this copy getting created unnecessarily, we use pointers.
2. Can directly edit the value without the need to return a value. This reduces the unnecessary code everytime to return a value.

&variable - Give me the memory address of the value this variable is pointing at
*pointer - Give me the value this memory address is pointing at.

### Struct
To avoid the confusion in ordering the arguments, we can combine the elements to form a type

```
type user struct {
	firstName string
	lastName  string
	birthDate string
}

```
- Since, struct has become complex variable when passing as an argument to a method, its
  better to use the pointers along with structs when passing as arguments.
- A function attached to a struct is called method.
- `function (d deck) print() {}` This is a **receiver** function and any variable of type deck has got the access to the print method


| Value Types  | Reference Types  |
|--------------|------------------|
|   int        |      slices      |
|   float      |     maps         |
|   string     |   channels       |
|   bool       |     pointers     |
|   structs    |  functions       |


## Testing

- Go has a lightweight test framework composed of the go test command and the testing package.
- You write a test by creating a file with a name ending in _test.go that contains functions named TestXXX with signature func (t *testing.T). 
- The test framework runs each such function; if the function calls a failure function such as t.Error or t.Fail, the test is considered to have failed.

