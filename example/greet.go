// example/greet.go

// Show how to write testable examples.
package greet

import "fmt"

type Dog struct{}

func (d Dog) Hello() string {
	return fmt.Sprint("woof woof!")
}

func Hello(name string) (string, error) {
	return fmt.Sprintf("Hello %s", name), nil
}
