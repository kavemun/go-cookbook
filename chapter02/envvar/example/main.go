package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/kavemun/go-cookbook/chapter02/envvar"
)

// Config will hold the config we capture from json file and envs vars
type Config struct {
	Version string `json:"version" required:"true"`
	IsSafe  bool   `json:"is_safe" default:"true"`
	Secret  string `json:"secret"`
}

func main() {
	var err error

	//create a temp file to hold an example json file
	tempfile, err := ioutil.TempFile("", "tmp")
	if err != nil {
		panic(err)
	}
	defer tempfile.Close()
	defer os.Remove(tempfile.Name())

	// create json file to hold our secrets
	secrets := `{
		"secret": "oh so secret"
		}`

	if _, err = tempfile.Write(bytes.NewBufferString(secrets).Bytes()); err != nil {
		panic(err)
	}

	// we can easily set environment variable as needed
	if err = os.Setenv("EXAMPLE_VERSION", "1.0.0.0"); err != nil {
		panic(err)
	}
	if err = os.Setenv("EX_ISSAFE", "false"); err != nil {
		panic(err)
	}

	c := Config{}
	if err = envvar.LoadConfig(tempfile.Name(), "EXAMPLE", &c); err != nil {
		panic(err)
	}

	fmt.Println("Secrets and lies contains =", secrets)

	// we can also read them back
	fmt.Println("EX_VERSION =", os.Getenv("EXAMPLE_VERSION"))
	fmt.Println("EX_ISSAFE =", os.Getenv("EX_ISSAFE"))

	//final config is a mix of json and environemnt variables
	fmt.Printf("Final COnfig: %#v\n", c)

}
