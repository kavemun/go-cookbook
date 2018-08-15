package main

import "github.com/kavemun/go-cookbook/chapter01/filedirs"

func main() {
	if err := filedirs.Operate(); err != nil {
		panic(err)
	}

	if err := filedirs.CapitalizerExample(); err != nil {
		panic(err)
	}
}
