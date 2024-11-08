package main

import (
	"fmt"

	"github.com/Dominux/gotcha/internal"
)

func main() {
	fmt.Println("Hello, lol")

	r := internal.NewRouter()
	r.RunRouter("8000")
}
