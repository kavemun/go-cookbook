package confformat

import (
	"fmt"
)

const (
	exampleTOML = `name="Thomas"
age=99
    `

	exampleJSON = `{"name":"Jonathan","age":98}`

	// NOTE: YAML is sensitive to tabs. VSCode autoformat may insert tab character
	//       which will cause an error in processing the YAML
	exampleYAML = `name: Yoshida
age: 97    
    `
)

// UnmarshalAll takes data in various formats and
// converts them into structs
func UnMarshallAll() error {
	t := TOMLData{}
	y := YAMLData{}
	j := JSONData{}

	if _, err := t.Decode([]byte(exampleTOML)); err != nil {
		return err
	}
	fmt.Println("TOML Unmarshal =", t)

	if err := j.Decode([]byte(exampleJSON)); err != nil {
		return err
	}
	fmt.Println("JSON Unmarshal =", j)

	if err := y.Decode([]byte(exampleYAML)); err != nil {
		return err
	}
	fmt.Println("YAML Unmarshal =", y)

	return nil
}
