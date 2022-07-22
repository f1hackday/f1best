package main

import (
	"fmt"

	"github.org/f1hackday/f1lib"
)

func main() {
	name := "frankie"
	got := f1lib.WhosBest(name)
	fmt.Println(got)
}
