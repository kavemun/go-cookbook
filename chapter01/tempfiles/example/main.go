package main

import "github.com/kavemun/go-cookbook/tempfiles"

func main() {
	if err := tempfiles.WorkWithTemp(); err != nil {
		panic(err)
	}
}