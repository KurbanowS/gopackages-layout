package main

import (
	"fmt"
	"gopackages/internal/wordz"
	col "gopackages/pkg/color"

	"github.com/fatih/color"
)

func main() {
	col.Greet()
	color.Red("Hello world")
	for i := 0; i < 5; i++ {
		fmt.Println(wordz.Hello)
		fmt.Println(wordz.Random())

	}
}
