package main

import (
	"fmt"

	"github.com/kavemun/go-cookbook/chapter02/ansicolor"
)

func main() {
	r := ansicolor.ColorText{
		TextColor: ansicolor.Red,
		Text:      "I'm red like The Flash!",
	}

	fmt.Println(r.String())

	r.TextColor = ansicolor.Green
	r.Text = "Now I'm green like the Hulk"

	fmt.Println(r.String())

	r.TextColor = ansicolor.ColorNone
	r.Text = " Back to Boring Normal"

	fmt.Println(r.String())
}
