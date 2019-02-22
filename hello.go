package helloworld

import "fmt"

const (
	packageVersion = "2.0.1"
)

// HelloWorld - print greeting
func HelloWorld() {
	fmt.Println(Hello("World"))
	fmt.Printf("[%s]\n", packageVersion)
}

// Hello - get greeting for name
func Hello(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}
