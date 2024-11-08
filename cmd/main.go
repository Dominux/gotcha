package main

import (
	"github.com/Dominux/gotcha/internal"
)

func main() {
	r := internal.NewRouter()
	r.RunRouter("8000")
}
