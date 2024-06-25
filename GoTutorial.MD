# Go Tutorial 

## Basics

- Executable is created based on package main and main function
- Other functions have to be imported
- It doesnt have OOP concept but we can create like OO using type
- Avoid `this` or `self` which will break Go convention
- Use `_` if we are not using any variable (Eg: index in for loop)
- In a way it resembles C language Structs and Pointers

### Receiver

`function (d deck) print() {}` This is a receiver function and any variable of type deck has got the access to the print method

### Pointer operations

&variable - Give me the memory address of the value this variable is pointing at
*pointer - Give me the value this memory address is pointing at

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

