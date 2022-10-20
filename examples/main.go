package main

import (
	"fmt"
	"github.com/tinysystemsio/errorpanic"
	"log"
)

func someErrorPanic() error {
	panic("oops")
	return fmt.Errorf("handled error")
}

func someErrorArbitraryArguments(i int) error {
	return fmt.Errorf("error%d", 100/i)
}

func main() {
	// safely run function which may panic
	if err := errorpanic.Wrap(someErrorPanic); err != nil {
		log.Printf("result: %v", err)
	}

	// example of run a function with arbitrary set of arguments
	if err := errorpanic.Wrap(func() error {
		return someErrorArbitraryArguments(13)
	}); err != nil {
		log.Printf("result: %v", err)
	}

	// example of run a function with arbitrary set of arguments; panic
	if err := errorpanic.Wrap(func() error {
		return someErrorArbitraryArguments(0)
	}); err != nil {
		// runtime error: integer divide by zero
		log.Printf("result: %v", err)
	}

}
