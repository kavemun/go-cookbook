package main

import "github.com/kavemun/go-cookbook/bytestrings"

func main() {

	err := bytestrings.WorkWithBuffer()
	if err != nil {
		panic(err)
	}

	//each of these print to stdout
	bytestrings.SearchString()
	bytestrings.ModifyString()
	bytestrings.StringReader()
}
