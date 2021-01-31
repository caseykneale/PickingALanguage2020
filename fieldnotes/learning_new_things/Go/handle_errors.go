package main

import (
	"fmt"
	"os"
)

func handle_errors(err error, txt string) {
	if err != nil {
		fmt.Printf("Failed to %s: %s\n", txt, err)
		os.Exit(1)
	}
}

func handle_panic(err error, int_err error) {
	if err != nil {
		panic(int_err)
	}
}
