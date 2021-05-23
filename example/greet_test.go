//ffs
package greet_test

import (
	"fmt"

	greet "github.com/luxcgo/go-test/example"
)

func Example_demo() {
	greeting, err := greet.Hello("World")
	if err != nil {
		panic(err)
	}
	fmt.Println(greeting)

	// Output:
	// Hello World
}

func ExampleHello_demos() {
	greeting, err := greet.Hello("Adele")
	if err != nil {
		panic(err)
	}
	fmt.Println(greeting)

	// Output:
	// Hello Adele
}

func ExampleDog() {
	fmt.Println("Dog type")

	// Output:
	// Dog type
}

func ExampleDog_demo() {
	fmt.Println("Dog type w/ demo label")

	// Output:
	// Dog type w/ demo label
}

func ExampleDog_Hello_demo() {
	fmt.Println("Dog type w/ demo label")

	// Output:
	// Dog type w/ demo label
}
