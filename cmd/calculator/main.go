package main

import (
	"os"

	"calculator"
)

func main() {
	status := calculator.Main()
	os.Exit(status)
}
