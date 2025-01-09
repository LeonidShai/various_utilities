package yamlparser

import (
	"fmt"
	"os"

	yaml "gopkg.in/yaml.v3"
)

var TestFileName = "yaml_parser/test.yaml"

func ParseYamlFile(filename string) {
	type Fru struct {
		Fruits []string `yaml:"fruits"`
	}

	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println(err)

		return
	}

	fru := &Fru{}
	if err := yaml.Unmarshal(data, fru); err != nil {
		fmt.Printf("failed to unmarshal yaml: %v", err)

		return
	}

	for _, elem := range fru.Fruits {
		fmt.Println(elem)
	}
}
