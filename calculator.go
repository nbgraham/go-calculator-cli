package calculator

import (
	"calculator/internal/add"
	"calculator/internal/multiply"
	"fmt"
	"os"
	"strconv"
)

func Main() int {
	if len(os.Args[1:]) < 3 {
		fmt.Fprintln(os.Stderr, "usage: calculator OPERATION A B")
		return 1
	}

	operation := os.Args[1]

	a, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}

	b, err := strconv.Atoi(os.Args[3])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}

	switch operation {
	case "add":
		fmt.Printf("sum: %d\n", add.Add(a, b))
	case "multiply":
		fmt.Printf("product: %d\n", multiply.Multiply(a, b))
	default:
		fmt.Fprintf(os.Stderr, "unknown operation: %s\n", operation)
		return 1
	}

	return 0
}
