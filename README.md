# goLearn
My journey learning Golang 

On this journey we gonna make use of [Learn Go with tests](https://quii.gitbook.io/learn-go-with-tests) as the name says, it teaches Go with practical exercises, testing your actual code and implementations of new features.

## About this repo 

This repo is gonna be divided on folders for each exercise given and letting on this readme a small explanation about what we learned and implemented on that folder.
Let's start 🚀

## [Hello World](https://github.com/benjaAguilar/goLearn/tree/main/helloWorld)

This exercise consist of creating a simple hello world app. Introducing our first steps on the language.

### What we learn:
- Create a go project
- Manage dependencies
With `go mod init example.com/hello` we create a go.mod file that enables dependency tracking allowing us to use go packages like e.g('fmt')
- packages
- imports
- func main()
To run a go file we need to have a main function wich is gonna execute the main behaviour of our app.
- create functions and constants
- capital Letters exports functions and variables
- testing

### A deep dive on testing
Writing a test is just like writing a function, with a few rules:
- It needs to be in a file with a name like xxx_test.go
- The test function must start with the word Test
- The test function takes one argument only t *testing.T
- To use the *testing.T type, you need to import "testing" package

- *t.Run()* allow us to run different subtests describing different scenarios.
- *t.Helper()* allow us to tell the test suite that this method is a helper. By doing this, when it fails, the line number reported will be in our function call rather than inside our test helper.

## Integers

We learn about the use of integers on Go, and something really intresting about Testing wich are the Testable Examples

### What is a testable example?
Testables examples are both test code and code examples on the docs of the package.
this is really usefull cause Example functions are compiled whenever tests are executed. Because such examples are validated by the Go compiler, you can be confident your documentation's examples always reflect current code behavior.

## Iteration
On golang there are no `while`, `do`, `until` keywords, you can only use: `for`.
we make use of `for` in the code exercises and Learn about Benchmarking.

### Benchmarking
Is a test that we can run on our test file to measure the code eficience.
