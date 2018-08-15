package main

import "github.com/kavemun/go-cookbook/chapter01/tempfiles"

func main() {
	if err := tempfiles.WorkWithTemp(); err != nil {
		panic(err)
	}
}
