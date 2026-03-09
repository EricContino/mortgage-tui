package main

import (
	"fmt"
	"os"
)

func commandExit(args ...string) error {
	fmt.Println("Closing the Mortgage TUI... Goodbye!")
	os.Exit(0)
	return nil
}
