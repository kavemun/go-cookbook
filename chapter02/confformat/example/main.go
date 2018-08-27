package main

import "github.com/kavemun/go-cookbook/chapter02/confformat"

func main() {
	if err := confformat.MarshalAll(); err != nil {
		panic(err)
	}

	if err := confformat.UnMarshallAll(); err != nil {
		panic(err)
	}

	if err := confformat.OtherJSONExamples(); err != nil {
		panic(err)
	}

}
