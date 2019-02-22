package helloworld

import "fmt"

const (
	packageVersion = "2.0.0"
)

// HelloWorld - print greeting
func HelloWorld() {
	fmt.Printf("%s (%s)!\n", "Hello world", packageVersion)
}

// Hello - get greeting for name
func Hello(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}
